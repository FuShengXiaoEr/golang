package main

import "fmt"

func write(storeChan chan int) {
	for i := 0; i < 50; i++ {
		storeChan <- i
		fmt.Printf("write : %v\n",i)
	}
	close(storeChan)
}

func read(storeChan chan int,exitChan chan bool) {
	for {
		v,ok := <-storeChan
		if !ok {
			break
		}
		fmt.Printf("read : %v\n",v)
	}
	exitChan <- true
	close(exitChan)
}

func main(){
	storeChan := make(chan int,200)
	exitChan := make(chan bool,1)

	go write(storeChan)
	go read(storeChan,exitChan)

	for {
		_,ok := <- exitChan
		if !ok {
			fmt.Println("over")
			break
		}
	}
}