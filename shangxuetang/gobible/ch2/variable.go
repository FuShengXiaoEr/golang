package main

import "fmt"

func main(){
	x,y := 1,2 
	fmt.Printf("x,y=%d,%d",x,y)

	x,y = y,x 
	fmt.Printf("x,y=%d,%d",x,y)
}