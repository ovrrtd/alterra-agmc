package migration

import (
	"agmc/database"
	"agmc/internal/model"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {

	conn := database.GetConnection()

	conn.AutoMigrate(tables...)
}

func Rollback() {

	conn := database.GetConnection()

	for i := len(tables) - 1; i >= 0; i-- {
		conn.Migrator().DropTable(tables[i])
	}
}
