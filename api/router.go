package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/theplant/htmlgo"
	"htmxgo/render"
	"log"
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

}

func addMovie(c *gin.Context) {

}

func listMovies(c *gin.Context) {
	movies, page := listMoviesFromDB(c)
	prevUrl := fmt.Sprintf("http://127.0.0.1:8080/movies?page=%d", page-1)
	afterUrl := fmt.Sprintf("http://127.0.0.1:8080/movies?page=%d", page+1)
	table := func(movies []Movie) HTMLComponent {
		movieTrs := make([]HTMLComponent, 0)
		for _, movie := range movies {
			movieTrs = append(movieTrs, MovieTableBody(movie))
		}

		return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
			table :=
				Table(
					MovieTableHead(),
					Tbody(movieTrs...),
				).Class("table min-w-full divide-y divide-gray-300")
			pagination :=
				Div(
					Button("«").Attr(render.HxGet, prevUrl, render.HxSwap, "innerHTML", render.HxTarget, "#movieTable").Class("join-item btn btn-sm"),
					Button(strconv.Itoa(page)).Class("join-item btn btn-sm"),
					Button("»").Attr(render.HxGet, afterUrl, render.HxSwap, "innerHTML", render.HxTarget, "#movieTable").Class("join-item btn btn-sm"),
				).Class("join flex justify-center mt-4")
			hs := HTMLComponents{
				table,
				pagination,
			}
			return hs.MarshalHTML(ctx)
		})
	}
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
