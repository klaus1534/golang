package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//html := template.Must(template.ParseFiles("c.html","a.txt"))
	//r.SetHTMLTemplate(html)
	r.Static("static","./static/")
	r.LoadHTMLGlob("static/**")
	r.Run()
}
