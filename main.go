package main

import (
	"context"
	"database/sql"
	"fmt"
	"gin-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
	// . "github.com/volatiletech/sqlboiler/queries/qm"
)

//ルーティングの設定
func setupRouter() *gin.Engine {
	r := gin.Default()

	// 静的ファイル読み込み
	r.Static("/assets", "./app/assets")
	r.LoadHTMLGlob("app/templates/*.html")

	// 初期表示
	r.GET("/", func(c *gin.Context) {
		ctx := context.Background()
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3300)/todos?parseTime=true")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		boil.DebugMode = true
		boil.SetDB(db)
		tasks, err := models.Tasks().All(ctx, db)
		if err != nil {
			fmt.Printf("Get tasks error: %v", err)
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	// タスク作成
	r.POST("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
