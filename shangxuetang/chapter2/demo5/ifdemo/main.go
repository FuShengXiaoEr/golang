package main

import "fmt"

func main(){
	// var age int
	// fmt.Scanln(&age)
	// if age > 18 {
	// 	fmt.Println("年量超了18岁")
	// }else {
	// 	fmt.Println("年龄合适")
	// }


	//打印九九惩罚口诀表
	for i := 1;i <= 9 ;i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t",j,i,i*j)
		}
		fmt.Println()
	}
}