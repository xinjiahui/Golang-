package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTP SERVER
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}
func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求，参数都放在URL上。（query param）,请求体中是没有数据的
	queryParam := r.URL.Query() //自动识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))

}
func main() {
	http.HandleFunc("/test", f1)
	http.HandleFunc("/hello", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
