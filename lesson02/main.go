package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	//r.LoadHTMLFiles("template/users/index.html","template/posts/index.html")
	// url 访问入口
	r.GET("users/index", func(c *gin.Context) {
		//HTML 第二个参数是访问具体文件路径
		c.HTML(http.StatusOK,"users/index.html",gin.H{
			"title":"users html",
		})
	})

	//访问入口
	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts/index.html",gin.H{
			"title": "post 主题",
		})
	})

	r.Run(":9099")
}
