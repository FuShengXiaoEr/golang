package main

import (
	"fmt"
	"sort"
	"sync"
)

var (
	myMap = make(map[int]int,10)
	// 互斥锁，sync同步的意思
	lock sync.Mutex
	wg sync.WaitGroup
)

func test( n int) {
	defer wg.Done()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock() 
}


func main(){
	var keys []int
	wg.Add(200)
	for i := 0; i < 50; i++ {
		go test(i)
	}
	wg.Wait()

	lock.Lock()
	
	for k := range myMap {
		keys = append(keys, k)
	}
	lock.Unlock()

	sort.Ints(keys)
	
	for _,k := range keys {
		lock.Lock()
		v := myMap[k]
		lock.Unlock()
		fmt.Printf("map[%d]=%d\n",k,v)
	}

}
