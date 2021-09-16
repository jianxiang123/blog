package main


// gin的基本使用
/*
func main(){
	//创建路由
	route:=gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello World")
	})
	err:=route.Run(":8080")
	if err !=nil{
		fmt.Println(err)
	}
}

// 自定义
func main(){
	//创建路由
	route:=gin.Default()
	s:=&http.Server{
		Addr: ":8080",
		Handler: route,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Hello word")
	})
	s.ListenAndServe()
}
 */

// 路由
// 1 API 参数
// 通过Context的 Parem方法获取值
/*
func main()  {
	route:=gin.Default()
	route.GET("/user/:name/*active", func(c *gin.Context) {
		name:=c.Param("name")
		actie:=c.Param("active")
		c.String(http.StatusOK,name+"-----"+actie)
	})
	route.Run(":8080")
}
 */

//URL参数
// URL 参数可以通过DefaultQuery 和Query方法获取
/*
func main()  {
	route:=gin.Default()
	route.GET("/mytest", func(c *gin.Context) {
		name:=c.DefaultQuery("name","hianx")
		c.String(http.StatusOK,name)
	})
	route.Run(":8080")
}
 */




