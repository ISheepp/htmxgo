package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"htmxgo/api"
	"htmxgo/render"
	"os"
	"os/signal"
)

func main() {
	render.Generate()

	// 启动 API 服务器
	apiPort := 8080
	go startAPIServer(apiPort)

	// 启动 HTML 文件服务器
	htmlPort := 8081
	go startHTMLServer(htmlPort)

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
		c.File("./render/index.html")
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("HTML Server listening on", addr)
	err := engine.Run(addr)
	if err != nil {
		panic(err)
	}
}
