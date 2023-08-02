package main

import (
	"fmt"
	"shangxuetang/customerManagerSystem/model"
	"shangxuetang/customerManagerSystem/service"
)

type customerView struct {
	// 定义必要的字段
	key string // 接受输入
	loop bool //是否继续循环
	service *service.CustomerService
}

func (view *customerView) list(){
	customers := view.service.List()
	fmt.Println("********************客户列表********************")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, customer := range customers {
		  fmt.Println(customer.PrintToTerminal())
	}
	fmt.Println("********************客户列表展示完成********************\n\n")
}
 
// 得到用户新的输入，并构建新的用户
func (view *customerView) add(){
	fmt.Println("********************添加客户********************")
	fmt.Println("姓名：")
	name :=""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	tel := ""
	fmt.Scanln(&tel)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	newCustome := model.NewCustomer2(name,gender,age,tel,email)

	view.service.Add(newCustome)
	fmt.Println("********************添加客户完成********************")
}
 
// 得到用户输入，删除用户
func (cv *customerView) delete(){
	fmt.Println("********************删除客户********************")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "N" || choice == "n"{
		return
	}else if choice == "Y" || choice == "y" {
		cv.service.Delete(id)
		fmt.Println("********************删除客户完成********************")
	}else{
		return
	}
	
}

func(cv *customerView) mainMenu(){
	for {
		fmt.Println("--------------------客户信息管理软件-------------------------")
		fmt.Println("                    1 添 加 客 户")
		fmt.Println("                    2 修 改 客 户")
		fmt.Println("                    3 删 除 客 户")
		fmt.Println("                    4 客 户 列 表")
		fmt.Println("                    5 退      出")
		fmt.Println("请选择（1-5）： ")

		fmt.Scanln(&cv.key)
		switch cv.key {
			case "1":
				cv.add()
			case "2":
				fmt.Println("修改客户")
			case "3":
				cv.delete()
			case "4":
				cv.list()
			case"5":
				cv.loop = false
			default:
				fmt.Println("请输入正确选项")
		}

		if !cv.loop {
			fmt.Println("您推出了客户管理系统")
			fmt.Println()
			break
		}
	}
}

func main(){
	cus := customerView {
		key: "",
		loop: true,
	}
	// 完成对结构体service字段初始化
	cus.service = service.NewCustomerService()
	cus.mainMenu()
}