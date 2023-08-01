package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c ", arr[i])
		
	}
	fmt.Println("\n-------------")

	// 随机生成五个数，然后反转打印
	var arr2 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr2); i++ {
		arr2[i] = rand.Intn(100) + 1
	}
	fmt.Println("原数组",arr2)
	for  i := 0;  i < len(arr2)/2;  i++ {
		arr2[i],arr2[len(arr2) - 1 - i] = arr2[len(arr2) - 1 - i],arr2[i]
	}

	fmt.Println("反转数组",arr2)
	 
	
}