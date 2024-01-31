package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"htmxgo/api"
	"htmxgo/render"
	"os"
	"os/signal"
	"strconv"
)

//go:embed static/*
var content embed.FS

func main() {
	render.Generate()

	// 启动 API 服务器
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		panic(err)
	}
	go startAPIServer(apiPort)

	// 启动 HTML 文件服务器
	htmlPort, err := strconv.Atoi(os.Getenv("HTML_PORT"))
	if err != nil {
		panic(err)
	}
	go startHTMLServer(htmlPort)

	// 留一个缓冲可以防止阻塞并且防止丢失信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func startAPIServer(port int) {
	api.NewSqliteDB()
	engine := api.NewRouter()
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("HTML Server listening on", addr)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}

func startHTMLServer(port int) {
	engine := gin.Default()
	engine.Any("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("HTML Server listening on", addr)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}
