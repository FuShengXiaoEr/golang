/*
*

	家庭记账小实战
	模拟实现一个基于文本界面的家庭记账软件
	掌握基本的编程技巧和调试技巧

*
*/
package main

import (
	"fmt"
)

func main(){
	choice := ""
	loop := true
	for {
		fmt.Println("-------------家庭收支记账软件-----------------")
		fmt.Println("             1.收支明细 ")
		fmt.Println("             2.登记收入")
		fmt.Println("             3.登记支出")
		fmt.Println("             4.退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&choice)
		switch choice {
			case "1":
				
			case "2":
			case "3":
				fmt.Println("登记支出：")
			case "4":
				
				loop = false
				
			default:
				fmt.Println("请输入正确的选项。。")


		}

		if !loop {
			break
		}
	}
	fmt.Println("您退出了家庭收支记账软件")
	
}