package controller

import (
	"agmc/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	return e
}

func TestGetBooksSuccess(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	req := httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := GetBooks(c)
	if !assert.NoError(t, err) {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    []Book `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	if !assert.NotNil(t, rBody.Data) {
		t.Fatal("data nil")
	}

	assert.IsType(t, []Book{}, rBody.Data)
}

func TestGetBookByIDSuccess(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	req := httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetBookByID(c)
	if !assert.NoError(t, err) {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    Book   `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	if !assert.NotNil(t, rBody.Data) {
		t.Fatal("data nil")
	}

	assert.IsType(t, Book{}, rBody.Data)
}

func TestGetBookByIDFail_NotFound(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	req := httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("10")

	err := GetBookByID(c)
	if !assert.NoError(t, err) {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusNotFound, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}
	assert.Nil(t, rBody.Data)
}

func TestCreateBookFail_NoPayload(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	// payload := CreateBookRequest{
	// 	Name: "",
	// }

	req := httptest.NewRequest(http.MethodPost, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	CreateBook(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	assert.Nil(t, rBody.Data)
}

func TestCreateBookFail_NameEmpty(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"
	var payload bytes.Buffer

	if err := json.NewEncoder(&payload).Encode(&CreateBookRequest{
		Name: "",
	}); err != nil {
		t.Fatal(err.Error())
	}

	req := httptest.NewRequest(http.MethodPost, path, &payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	CreateBook(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	assert.Nil(t, rBody.Data)
}

func TestCreateBookSuccess(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"
	var payload bytes.Buffer

	if err := json.NewEncoder(&payload).Encode(&CreateBookRequest{
		Name: "Buku",
	}); err != nil {
		t.Fatal(err.Error())
	}

	req := httptest.NewRequest(http.MethodPost, path, &payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	CreateBook(c)

	if !assert.Equal(t, http.StatusCreated, rec.Code) {
		t.Fatal("wrong response status code")
	}
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}

	assert.NotNil(t, rBody.Data)
	assert.Positive(t, rBody.Data.ID)
}

func TestDeleteByIDFail_NotFound(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	req := httptest.NewRequest(http.MethodDelete, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("10")

	DeleteBookByID(c)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}
	assert.Nil(t, rBody.Data)
}

func TestDeleteByIDSuccess(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	req := httptest.NewRequest(http.MethodDelete, path, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteBookByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}
	assert.Nil(t, rBody.Data)
}

func TestUpdateByIDFail(t *testing.T) {
	e := initEcho()
	path := "http://localhost:8080/v1/books"

	var testCases = []struct {
		name         string
		id           string
		codeExpected int
		payload      *UpdateBookRequest
	}{
		{
			name:         "ID not found",
			id:           "10",
			codeExpected: http.StatusNotFound,
			payload: &UpdateBookRequest{
				Name: "Buku saya",
			},
		},
		{
			name:         "Book Name Empty",
			id:           "1",
			codeExpected: http.StatusBadRequest,
			payload: &UpdateBookRequest{
				Name: "",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name,
			func(t *testing.T) {
				var payload = new(bytes.Buffer)
				if testCase.payload != nil {
					if err := json.NewEncoder(payload).Encode(testCase.payload); err != nil {
						t.Fatal(err.Error())
					}
				}

				req := httptest.NewRequest(http.MethodPut, path, payload)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()

				c := e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues(testCase.id)

				UpdateBookByID(c)

				assert.Equal(t, testCase.codeExpected, rec.Code)
				type response struct {
					Message string `json:"message"`
					Data    *Book  `json:"data"`
				}

				var rBody response
				if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
					t.Fatalf("error decode: %s", err.Error())
				}
				assert.Nil(t, rBody.Data)
			},
		)

	}

}

func TestUpdateByIDSuccess(t *testing.T) {
	e := initEcho()

	path := "http://localhost:8080/v1/books"

	var payload = new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(UpdateBookRequest{
		Name: "Buku saya",
	}); err != nil {
		t.Fatal(err.Error())
	}

	req := httptest.NewRequest(http.MethodDelete, path, payload)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	UpdateBookByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	type response struct {
		Message string `json:"message"`
		Data    *Book  `json:"data"`
	}

	var rBody response
	if err := json.NewDecoder(rec.Body).Decode(&rBody); err != nil {
		t.Fatalf("error decode: %s", err.Error())
	}
	assert.NotNil(t, rBody.Data)
}
