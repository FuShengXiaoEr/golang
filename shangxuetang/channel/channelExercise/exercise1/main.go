package main

import (
	"fmt"
	"sync"
)

func calcSum(n int, resChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	resChan <- sum
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan int, 8)
	done := make(chan struct{}) // 创建一个空的done通道
	var wg sync.WaitGroup

	// 启动一个协程，将1-2000的数放入numChan中
	wg.Add(1)
	go func() {
		for i := 1; i <= 2000; i++ {
			numChan <- i
		}
		close(numChan)
		wg.Done()
	}()

	// 启动8个协程，从numChan取出数据并计算1+2+...+n的值，并存放到resChan中
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go calcSumWorker(numChan, resChan, done, &wg) // 将done通道传递给计算协程
	}

	// 等待所有计算协程完成工作
	go func() {
		wg.Wait()
		close(done) // 关闭done通道
	}()

	// 遍历resChan，显示结果
	for res := range resChan {
		fmt.Printf("res=%d\n", res)
	}
}

func calcSumWorker(numChan <-chan int, resChan chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n, ok := <-numChan:
			if !ok {
				return // numChan 已经关闭，计算协程退出
			}
			sum := 0
			for i := 1; i <= n; i++ {
				sum += i
			}
			resChan <- sum
		case <-done:
			return // 主协程通知计算协程停止
		}
	}
}
