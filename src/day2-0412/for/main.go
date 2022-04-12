package main

import "fmt"

//for 循环
func main(){
  // 基本格式
  for i :=0;i<10;i++{
	  fmt.Println(i)
  }
  // 变种1
  var i = 5
  for ;i<10;i++{
	  fmt.Println(i)
  }
  //无限循环
  //for {
	//  print.Println("123")
  //}
  // for range 循环
  s := "hello"
  for i,v := range s {
	  fmt.Printf("%d %c\n",i,v)
  }
}