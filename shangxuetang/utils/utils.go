package utils

import (
	"fmt"
	"strconv"
)

func ConversStrToInt(arr []string) []int {
	intArr := make([]int,len(arr))

	for i := 0; i < len(arr); i++ {
		num,err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Printf("第 %d 数 %s 转换错误",i,arr[i])
		}
		intArr[i] = num
	}
	return intArr
} 