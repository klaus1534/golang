package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type UserInfo struct {
	Name string
	Gender string
	Age int
}
func readHello(w http.ResponseWriter,r *http.Request)  {
	// 解析模版
	t,err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err:%v",err)
		return
	}

	// 利用给定数据渲染模版，并将结果写入w
	user := UserInfo{
		Name: "ABC",
		Gender: "女",
		Age: 12,
	}

	// 渲染模版
	//err = t.Execute(w,"tom")
	err = t.Execute(w,user)
	if err != nil {
		fmt.Println("render template failed,err:%v",err)
		return
	}
}
func main() {

	http.HandleFunc("/",readHello)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println("http access error ,%v",err)
		return
	}
}
