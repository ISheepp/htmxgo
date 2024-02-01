package api

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var sqliteDB *gorm.DB

type Movie struct {
	Id              int     `gorm:"id" json:"id"`
	ImdbId          string  `gorm:"imdb_id" json:"imdbId"`
	Title           string  `gorm:"title" json:"title"`
	Director        string  `gorm:"director" json:"director"`
	Year            int     `gorm:"year" json:"year"`
	Rating          string  `gorm:"rating" json:"rating"`
	Genres          string  `gorm:"genres" json:"genres"`
	Runtime         int     `gorm:"runtime" json:"runtime"`
	Country         string  `gorm:"country" json:"country"`
	Language        string  `gorm:"language" json:"language"`
	ImdbScore       float64 `gorm:"imdb_score" json:"imdbScore"`
	ImdbVotes       int     `gorm:"imdb_votes" json:"imdbVotes"`
	MetacriticScore float64 `gorm:"metacritic_score" json:"metacriticScore"`
}

func NewSqliteDB() {

	db, err := gorm.Open(sqlite.Open("./api/movies.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqliteDB = db
	sqliteDB.AutoMigrate(&Movie{})

}
