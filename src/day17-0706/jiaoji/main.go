package main

import "fmt"

func main(){
	A := [3]string{"aa", "bb", "cc"}
	B := [3]string{"ac", "bc", "cc"}
	for _,v1 := range A {
		for _,v2 := range B{
			if v1 == v2{
			fmt.Println(v1)
			}
		}
	}
}