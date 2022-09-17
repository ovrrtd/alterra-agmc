package controller

import (
	"agmc/config"
	"agmc/model"
	"agmc/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupDB() {

	cfg := config.ConfigDB{
		User: os.Getenv("APP_DB_TEST_USER"),
		Pass: os.Getenv("APP_DB_TEST_PASS"),
		Port: os.Getenv("APP_DB_TEST_PORT"),
		Host: os.Getenv("APP_DB_TEST_HOST"),
		Name: os.Getenv("APP_DB_TEST_NAME"),
	}

	config.InitDB(cfg)
}

func initEchoWithDB() *echo.Echo {
	err := godotenv.Load("../.env")

	if err != nil {
		panic(err)
	}

	setupDB()
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	return e
}

func truncateDB() {
	if config.DB == nil {
		panic("db not initialized")
	}

	config.DB.Unscoped().Where("id IS NOT NULL").Delete(&model.User{})
}

func initDummyUser() model.User {
	if config.DB == nil {
		panic("db not initialized")
	}
	user := model.User{
		Name:     "coba",
		Password: "1234",
		Email:    "coba@gmail.com"}

	if err := config.DB.Create(&user).Error; err != nil {
		panic(err)
	}

	return user
}

func TestCreateUserSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()

	path := "http://localhost:8080/v1/users"
	var payload bytes.Buffer

	if err := json.NewEncoder(&payload).Encode(&CreateUserRequest{
		Name:     "coba",
		Password: "1234",
		Email:    "coba@gmail.com",
	}); err != nil {
		t.Fatal(err.Error())
	}
	req := httptest.NewRequest(http.MethodPost, path, &payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := CreateUser(c)
	if !assert.NoError(t, err) {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusCreated, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    User   `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	if !assert.NotNil(t, rBody.Data) {
		t.Fatal("data nil")
	}

	assert.IsType(t, User{}, rBody.Data)
}

func TestCreateUserFail(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()

	var testCases = []struct {
		name         string
		codeExpected int
		payload      *CreateUserRequest
	}{
		{
			name:         "Payload Name Empty",
			codeExpected: http.StatusBadRequest,
		},
		{
			name:         "Payload Email invalid",
			codeExpected: http.StatusBadRequest,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var payload bytes.Buffer
			if testCase.payload != nil {
				if err := json.NewEncoder(&payload).Encode(testCase.payload); err != nil {
					t.Fatal(err.Error())
				}
			}

			req := httptest.NewRequest(http.MethodPost, path, &payload)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			err := CreateUser(c)
			if !assert.NoError(t, err) {
				t.Fatal(err.Error())
			}

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestGetUserByIDFail(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
	}{
		{
			name:         "User Not Found",
			id:           strconv.Itoa(int(userDummy.ID) + 1),
			codeExpected: http.StatusNotFound,
		},
		{
			name: "User Not Found",
			id:   strconv.Itoa(int(userDummy.ID) + 2),

			codeExpected: http.StatusNotFound,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			GetUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestGetUserByIDSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
	}{
		{
			name:         "User Found",
			id:           strconv.Itoa(int(userDummy.ID)),
			codeExpected: http.StatusOK,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			GetUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.NotNil(t, rBody.Data)
			assert.Equal(t, userDummy.ID, rBody.Data.ID)
			assert.Equal(t, userDummy.Name, rBody.Data.Name)
			assert.Equal(t, userDummy.CreatedAt, rBody.Data.CreatedAt)

		})
	}

}

func TestGetUsersSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	initDummyUser()

	var testCases = []struct {
		name         string
		codeExpected int
	}{
		{
			name:         "Users Found",
			codeExpected: http.StatusOK,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			GetUsers(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    []User `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.NotNil(t, rBody.Data)
		})
	}

}

func TestDeleteUserByIDSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
	}{
		{
			name:         "Delete User Success",
			id:           strconv.Itoa(int(userDummy.ID)),
			codeExpected: http.StatusOK,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			DeleteUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestDeleteUserByIDFail(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
	}{
		{
			name:         "User ID Not Found",
			id:           strconv.Itoa(int(userDummy.ID) + 2),
			codeExpected: http.StatusNotFound,
		},
		{
			name:         "User ID Not Found",
			id:           strconv.Itoa(int(userDummy.ID) + 3),
			codeExpected: http.StatusNotFound,
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			DeleteUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestUpdateUserByIDFail(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
		payload      *UpdateUserRequest
	}{
		{
			name:         "User ID not found",
			id:           strconv.Itoa(int(userDummy.ID) + 1),
			codeExpected: http.StatusNotFound,
			payload: &UpdateUserRequest{
				Name:     "alterra",
				Password: "1234"},
		},
		{
			name:         "Payload name empty",
			id:           strconv.Itoa(int(userDummy.ID)),
			codeExpected: http.StatusBadRequest,
			payload: &UpdateUserRequest{
				Name:     "",
				Password: "1234",
			},
		},
		{
			name:         "Payload Password empty",
			id:           strconv.Itoa(int(userDummy.ID)),
			codeExpected: http.StatusBadRequest,
			payload: &UpdateUserRequest{
				Name:     "alterra",
				Password: "",
			},
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var payload bytes.Buffer

			if err := json.NewEncoder(&payload).Encode(testCase.payload); err != nil {
				t.Fatal(err.Error())
			}
			req := httptest.NewRequest(http.MethodPut, path, &payload)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			UpdateUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestUpdateUserByIDSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
		payload      *UpdateUserRequest
	}{
		{
			name:         "Update success",
			id:           strconv.Itoa(int(userDummy.ID)),
			codeExpected: http.StatusOK,
			payload: &UpdateUserRequest{
				Name:     "alterra",
				Password: "1234",
			},
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var payload bytes.Buffer

			if err := json.NewEncoder(&payload).Encode(testCase.payload); err != nil {
				t.Fatal(err.Error())
			}
			req := httptest.NewRequest(http.MethodPut, path, &payload)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			UpdateUserByID(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string `json:"message"`
				Data    *User  `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}

func TestLoginSuccess(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		codeExpected int
		payload      *LoginRequest
	}{
		{
			name:         "Login success",
			codeExpected: http.StatusOK,
			payload: &LoginRequest{
				Email:    userDummy.Email,
				Password: userDummy.Password,
			},
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var payload bytes.Buffer

			if err := json.NewEncoder(&payload).Encode(testCase.payload); err != nil {
				t.Fatal(err.Error())
			}
			req := httptest.NewRequest(http.MethodPost, path, &payload)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			Login(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string         `json:"message"`
				Data    *LoginResponse `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.NotNil(t, rBody.Data)
			assert.NotNil(t, rBody.Data.Token)
		})
	}

}

func TestLoginFail(t *testing.T) {
	e := initEchoWithDB()
	truncateDB()
	userDummy := initDummyUser()

	var testCases = []struct {
		name         string
		codeExpected int
		payload      *LoginRequest
	}{
		{
			name:         "Login fail empty email",
			codeExpected: http.StatusBadRequest,
			payload: &LoginRequest{
				Email:    "",
				Password: userDummy.Password,
			},
		},
		{
			name:         "Login fail wrong email format",
			codeExpected: http.StatusBadRequest,
			payload: &LoginRequest{
				Email:    "asd",
				Password: userDummy.Password,
			},
		},
		{
			name:         "Login fail user not found",
			codeExpected: http.StatusNotFound,
			payload: &LoginRequest{
				Email:    "asd@gmail.com",
				Password: userDummy.Password,
			},
		},
		{
			name:         "Login fail empty password",
			codeExpected: http.StatusBadRequest,
			payload: &LoginRequest{
				Email:    userDummy.Email,
				Password: "",
			},
		},
		{
			name:         "Login fail wrong password",
			codeExpected: http.StatusBadRequest,
			payload: &LoginRequest{
				Email:    userDummy.Email,
				Password: "1",
			},
		},
	}

	path := "http://localhost:8080/v1/users"
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var payload bytes.Buffer

			if err := json.NewEncoder(&payload).Encode(testCase.payload); err != nil {
				t.Fatal(err.Error())
			}
			req := httptest.NewRequest(http.MethodPost, path, &payload)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			Login(c)

			assert.Equal(t, testCase.codeExpected, rec.Code)
			type response struct {
				Message string         `json:"message"`
				Data    *LoginResponse `json:"data"`
			}

			var rBody response
			if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
				t.Fatalf("error decode: %s", err.Error())
			}
			assert.Nil(t, rBody.Data)
		})
	}

}
