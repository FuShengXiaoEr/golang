package main

import (
	"fmt"
	_ "strconv"
)

func putNum(intChan chan int ){
	for i := 1 ; i <= 80;i++ {
		intChan <- i
	}

	// 关闭intChan
	close(intChan)
}

func primeNUm (intChan chan int,  primeChan chan int, exitChan chan bool) {
	
	for {
		var flag bool
		num,ok := <- intChan
		if !ok {
			break
		}
		for i := 2; i < num; i++ {
			if num % i == 0 { // 不是素数
				flag = true			
				break
			}
		}
		if !flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个协程取完了数据")
	// 这里还不能关闭primeChan
	// 想exit写入
	exitChan <- true
}

func main(){
	intChan := make(chan int,1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	// 开启4个协程，取数判断
	for i := 0; i < 4; i++ {
		go primeNUm(intChan,primeChan,exitChan)
		
	}

	go func() {
		for i := 0;i < 4;i++ {
			<-exitChan
		}
		// 当从exitChan取出4个结果，就可以关闭primeNUm
		close(primeChan)
	}()
	
	//遍历取值
	for{
		res,ok := <-primeChan 
		if !ok {
			break
		}
		fmt.Printf("素数为%d\n",res)
	}
	
	fmt.Println("主线程退出")

}