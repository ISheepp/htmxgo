package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/theplant/htmlgo"
	"htmxgo/render"
	"log"
	"os"
	"strconv"
)

func NewRouter() (engine *gin.Engine) {
	engine = gin.Default()
	engine.Use(CORSMiddleware())

	engine.GET("/movies", listMovies)
	engine.GET("/movie", getMovieById)
	engine.POST("/movie", addMovie)
	engine.DELETE("/movie", deleteMovie)
	engine.PUT("/movie", updateMovie)
	return engine
}

func getMovieById(c *gin.Context) {
	movieId := c.Query("id")

	movie := Movie{}
	sqliteDB.Where("id = ?", movieId).First(&movie)
	// dialog := render.MovieDialog("update", "update_model", render.HxPut, movie.Title, movie.Director, strconv.Itoa(movie.Id))
	dialog := Div(
		Input("movieId").Type("hidden").Value(movieId).Class("input input-bordered w-full max-w-xs"),
		Input("update_title").Type("text").Value(movie.Title).Placeholder("Title").Class("input input-bordered w-full max-w-xs"),
		Input("update_director").Type("text").Value(movie.Director).Placeholder("Director").Class("input input-bordered w-full max-w-xs mt-4"),
	).Class("mt-4").Id("UpdateDialog")
	Fprint(c.Writer, dialog, c)
}

func updateMovie(c *gin.Context) {
	movieId := c.PostForm("movieId")
	if movieId == "" {
		log.Println("movieId is empty")
		return
	}
	movie := Movie{}
	sqliteDB.Where("id = ?", movieId).First(&movie)
	movie.Title = c.PostForm("update_title")
	movie.Director = c.PostForm("update_director")
	sqliteDB.Save(&movie)
	dialog := render.UpdateDialog("Update", "update_model", render.HxPut, "", "", "")
	movies, table := listMoviesAction(c)
	components := HTMLComponents{
		table(movies),
		dialog,
	}
	Fprint(c.Writer, components, c)
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
	dialog := render.MovieDialog("Add", "add_model", render.HxPost, "", "", "")
	Fprint(c.Writer, dialog, c)
}

func listMovies(c *gin.Context) {
	movies, table := listMoviesAction(c)

	Fprint(c.Writer, table(movies), c)
}

func listMoviesAction(c *gin.Context) ([]Movie, func(movies []Movie) HTMLComponent) {
	movies, page := listMoviesFromDB(c)
	prevUrl := fmt.Sprintf("http://127.0.0.1:%s/movies?page=%d", os.Getenv("API_PORT"), page-1)
	afterUrl := fmt.Sprintf("http://127.0.0.1:%s/movies?page=%d", os.Getenv("API_PORT"), page+1)
	table := MovieTable(prevUrl, page, afterUrl)
	return movies, table
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
