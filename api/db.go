package api

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqliteDB *gorm.DB

type Movie struct {
	Id       int
	Title    string
	Director string
	Year     int
}

func NewSqliteDB() {

	db, err := gorm.Open(sqlite.Open("./api/movies.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqliteDB = db
	sqliteDB.AutoMigrate(&Movie{})

}
