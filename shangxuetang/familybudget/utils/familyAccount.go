package utils

import (
	"fmt"
)

type Account struct  {
	// 保存用户的输入
	key string
	// 控制是否退出
	loop bool

	// 余额
	balance float64
	// 每次收支的金额
	money float64
	details string 

}

func (this *Account) ShowDetails(flag bool) {
	fmt.Println("-------------家庭收支记账明细-----------------")
	if flag {
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有明细，请先记录一条")
	}
}

func (this *Account) Input() {
	fmt.Println("本次收入金额：")
	var money float64
	fmt.Scanln(&money)
	
}

