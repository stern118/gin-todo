package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ルーティングの設定
func setupRouter() *gin.Engine {
	r := gin.Default()
	// HTMLファイル読み込み
	message := "Todo List"
	r.LoadHTMLGlob("app/templates/*")

	// 初期表示
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": message,
		})
	})

	// タスク作成
	r.POST("/create", func(c *gin.Context) {
		task := c.PostForm("task")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"task": task,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
