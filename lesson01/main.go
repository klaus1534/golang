package main

import "github.com/gin-gonic/gin"

//func sayHello(w http.ResponseWriter, r *http.Request)  {
//	b ,_ := ioutil.ReadFile("./hello.html")
//	_,_ = fmt.Fprintf(w,string(b))
//}

func getHello(c *gin.Context) {
	c.JSON(200,gin.H{
		"message": "Hello golang!",
	})
}

func main()  {
	//http.HandleFunc("/hello",sayHello)
	//err := http.ListenAndServe(":9090",nil) // if method handler is write nil
	//if err != nil {
	//	fmt.Printf("http server failed,err:%v\n",err)
	//	return
	//}
	// 返回默认的路由引擎
	r := gin.Default()

	r.GET("/hello",getHello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"DELETE",
		})
	})

	//run server
	r.Run(":9091")
}


