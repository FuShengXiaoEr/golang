package main

import (
	"fmt"
	"os"
	"shangxuetang/utils"
)

/**
二分查找
**/
func main(){
	inputStr := os.Args[1:]
	intStr := utils.ConversStrToInt(inputStr)
	leftIndex := 0
	rightIndex := len(intStr)
	BinaryFind(&intStr,leftIndex,rightIndex,9)
	
}

func BinaryFind(arr *[]int,leftIndex int,rightIndex int,findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没有查找到该数")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex,middle - 1,findVal)
	}else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle + 1, rightIndex,findVal)
	}else{
		fmt.Printf("找到了，位于第%d位置",middle + 1)
	}
}