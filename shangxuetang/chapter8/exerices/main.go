package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	随机生成10个整数（1——100）保存到数组，并倒叙打印以及求平均值，最大值，最大值下标，并查找里面是否有55
**/
func test1(){
	intArr := [5]int{21,32,11,19,0}
	for i := 0;i < len(intArr);i++ {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100)
		intArr[i] = num
	}
	fmt.Println("origion array: ",intArr)
	// 排序 bubblesort
	for i := 0;i < len(intArr);i++  {
		isSort := false
		for j := 0; j > len(intArr) - 1 - i;j++ {
			if intArr[j] > intArr[j+1] {
				intArr[j],intArr[j+1] = intArr[j+1],intArr[j]
				isSort = true
			}
		}
		if !isSort {
			return
		}
	}
	// 排序后的数组
	fmt.Println("sorted finish: ",intArr)



}

func main(){
	
	test1()
}