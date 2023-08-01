package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputArr := os.Args[1:]
	fmt.Println("输入原数组:",inputArr)
	intArr := converToIntSlice(inputArr)

	for i := 0; i < len(intArr) - 1; i++ {

		for j := 0; j < len(intArr) - 1 -i; j++ {
			if intArr[j] > intArr[j+1] {
				// fmt.Println(intArr[j],intArr[j+1],j)
				intArr[j],intArr[j+1] = intArr[j+1],intArr[j]
				// fmt.Println(intArr[j],intArr[j+1],j)
			}
		
		}

	}
	fmt.Println("排序后的数组：",intArr)
}

func converToIntSlice(strSlice []string)[]int {
	intSlice := make([]int,len(strSlice))

	for i,str := range strSlice {
		num ,err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("错误：第 %d 个参数 %q 不是有效的整数\n",i+1,str)
			os.Exit(1)
		}
		intSlice[i] = num
	}
	return intSlice
}