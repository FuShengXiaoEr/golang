package main

import (
	"fmt"
	"time"
)


func test(){
	defer func(){
		
		if err := recover();err != nil {
			fmt.Println("err=",err)
		}
	}()
	a := 10
	b := 0
	res := a /b
	fmt.Println("res=",res)
}

func main(){
	test()
	for {
		fmt.Println("main后面的函数")
		time.Sleep(time.Second)
	}
}