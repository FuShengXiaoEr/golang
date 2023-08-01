package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Machine struct {
	Name string
}

type Phone struct {
	Machine
}

func (phone Phone) Start() {
	fmt.Println(phone.Name,"手机开始工作")
}

func (phone Phone) Stop(){
	fmt.Println(phone.Name,"手机停止工作")
}

type Computer struct {
	Machine
}

func (computer Computer)Start(){
	fmt.Println(computer.Name,"电脑开始工作")
}

func (computer Computer)Stop(){
	fmt.Println(computer.Name,"电脑停止工作")
}

func main() {
	var machine [3]Usb // 将数组元素类型声明为 Usb 接口类型
	machine[0] = Phone{Machine: Machine{Name: "小米"}}  // 使用指针赋值给 Usb 接口类型的变量
	machine[1] = Phone{Machine: Machine{Name: "VIVO"}}
	machine[2] = Computer{Machine: Machine{Name: "联想"}}

	for _,val := range machine {
		fmt.Println(val)
	}
	
}