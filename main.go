package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ルーティングの設定
func setupRouter(tasksPtr *[]string) *gin.Engine {
	r := gin.Default()
	// HTMLファイル読み込み
	r.LoadHTMLGlob("app/templates/*")

	// 初期表示
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": *tasksPtr,
		})
	})

	// タスク作成
	r.POST("/create", func(c *gin.Context) {
		// tasks sliceにPost Dataを追加
		tasks := append(*tasksPtr, c.PostForm("task"))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	return r
}

func main() {
	// taskリスト
	var tasks []string
	tasks = append(tasks, "Goの勉強をする")
	tasks = append(tasks, "バスのチケットを取る")
	tasks = append(tasks, "履歴書を書く")

	r := setupRouter(&tasks)
	r.Run(":8080")
}
