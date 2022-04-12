package main

import "fmt"

func main() {
	for i := 0;i < 10;i++{
		for j := 1;j<10;j++{
			if i >= j {
				s := i * j
		  		fmt.Printf("%d x %d = %d ",j,i,s)
			}	
		}
        fmt.Println()
	}
}