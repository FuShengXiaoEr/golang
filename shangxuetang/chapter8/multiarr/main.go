package main

import "fmt"

func main(){

	var arr [4][4]int
	arr[0][0] = 1
	arr[1][1] = 1
	arr[2][2] = 1
	arr[3][3] = 1

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}
}
