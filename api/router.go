package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/theplant/htmlgo"
	"log"
	"os"
	"strconv"
)

func NewRouter() (engine *gin.Engine) {
	engine = gin.Default()
	engine.Use(CORSMiddleware())

	engine.GET("/movies", listMovies)
	engine.POST("/movie", addMovie)
	engine.DELETE("/movie", deleteMovie)
	engine.PUT("/movie", updateMovie)
	return engine
}

func updateMovie(c *gin.Context) {

}

func deleteMovie(c *gin.Context) {
	movieId := c.Query("id")

	sqliteDB.Where("id = ?", movieId).Delete(&Movie{})
	listMovies(c)
}

func addMovie(c *gin.Context) {
	// 获取formdata的参数
	// 保存到数据库
	// 返回
	title := c.PostForm("title")
	director := c.PostForm("director")
	newMovie := Movie{
		Title:           title,
		Director:        director,
		Year:            2019,
		Rating:          "9.9",
		Genres:          "Drama",
		Runtime:         120,
		Country:         "China",
		Language:        "English",
		ImdbScore:       9.9,
		ImdbVotes:       10000,
		MetacriticScore: 9.9,
		ImdbId:          "123",
	}
	sqliteDB.Create(&newMovie)
	listMovies(c)
}

func listMovies(c *gin.Context) {
	movies, page := listMoviesFromDB(c)
	prevUrl := fmt.Sprintf("http://127.0.0.1:%s/movies?page=%d", os.Getenv("API_PORT"), page-1)
	afterUrl := fmt.Sprintf("http://127.0.0.1:%s/movies?page=%d", os.Getenv("API_PORT"), page+1)
	table := MovieTable(prevUrl, page, afterUrl)

	Fprint(c.Writer, table(movies), c)
}

func listMoviesFromDB(c *gin.Context) ([]Movie, int) {
	var movies []Movie
	param := c.Query("page")
	page := 1
	if param != "" {
		pageParam, err := strconv.Atoi(param)
		if err != nil {
			log.Fatal(err)
		}
		if pageParam < 1 {
			pageParam = 1
		}
		page = pageParam
	}

	pageSize := 10
	offset := (page - 1) * pageSize
	// 计算页数

	result := sqliteDB.Limit(pageSize).Offset(offset).Find(&movies)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return movies, page
}
