package main

import (
	"fmt"
	"time"
)

func func1(){
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		time.Sleep(time.Second)
	}
}

func func2(){
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println("func2发生错误")
		}
	}()
	var ma map[int]string
	ma[1] = "ss"
}

func main(){
	go func1()
	go func2()

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
	}
}