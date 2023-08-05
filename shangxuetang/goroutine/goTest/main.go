package main

import (
	"fmt"
	"strconv"
	"time"
)

func helloWorld(){
	for i := 0; i < 10; i++ {
		fmt.Println("goroutine " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	go helloWorld()
	for i := 0; i < 10; i++ {
		fmt.Println("main run" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	
}