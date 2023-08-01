package main

import (
	"fmt"
	// "math/rand"
	"strconv"
	// "time"
)

func main(){
	// var randInt int
	// var count int
	
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	randInt = rand.Intn(100) + 1
	// 	fmt.Println("随机数 "+strconv.Itoa(randInt))
	// 	if randInt == 99|| randInt == 50 || count == 100 {
	// 		fmt.Printf("随机生成了%v次", count)
	// 		break
	// 	}
	// 	count++
		
	// }

	// 问题一，提示登录机会
	var userName string
	var passWord string
	for i := 1; i < 4; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scan(&userName)
		fmt.Println("请输入密码：")
		fmt.Scan(&passWord)

		if userName == "张无忌" && passWord == " 888" {
			fmt.Println("登陆成功")
			break
		}else {
			surplus := 3 - i
			fmt.Println("密码或者用户名错误，还有" + strconv.Itoa(surplus) + "次机会")
			fmt.Printf("username = %v , password = %v",userName,passWord)
		}
	}
}