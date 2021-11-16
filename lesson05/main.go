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
func mapAction(w http.ResponseWriter,r *http.Request) {
	// 解析模版
	 t,err := template.ParseFiles("./ts.tmpl")
	if err != nil {
		fmt.Println("the template failed,err:%v",err)
		return
	}
	m1 := map[string]interface{}{
		"Name": "名称",
		"Gender": "男",
		"Age":18,
	}
	// 渲染模版
	err = t.Execute(w,m1)
}
func main() {

	http.HandleFunc("/",mapAction)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println("http input failed,err:%v",err)
		return
	}
}
