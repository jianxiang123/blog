package main

import "fmt"

//表单数据
/*
func main()  {
	route:=gin.Default()
	route.POST("/form", func(c *gin.Context) {
		// 获取默认值
		type1:=c.DefaultPostForm("xxx","aaa")
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		hobby:=c.PostFormArray("hobby")
		//fmt.Println(type1,username,password,hobby)
		c.String(http.StatusOK,fmt.Sprintf("type is %s username is " +
			"%s password is %s hobby is %s",type1,username,password,hobby))
	})
	route.Run(":8080")
}
 */

//上传单个文件
var db *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}