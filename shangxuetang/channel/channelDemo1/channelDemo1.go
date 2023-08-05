package main

import "fmt"
func main(){
	var intchan chan int
	intchan = make(chan int,3)
	
	intchan <- 10

	var num int
	num = <-intchan
	fmt.Println(num)
} 