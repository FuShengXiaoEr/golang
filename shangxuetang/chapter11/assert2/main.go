package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (phone Phone) Start() {
	fmt.Println(phone.Name,"手机开始工作")
}

func (phone Phone) Stop(){
	fmt.Println(phone.Name,"手机停止工作")
}

func (phone Phone) Call(){
	fmt.Println(phone.Name,"打电话")
}
type Computer struct {
	Name string
}

func (computer Computer)Start(){
	fmt.Println(computer.Name,"电脑开始工作")
}

func (computer Computer)Stop(){
	fmt.Println(computer.Name,"电脑停止工作")
}

type UseMachine struct {

}
func (useMachine UseMachine) working (usb Usb) {
	usb.Start()
	if phone ,ok := usb.(Phone) ;ok{
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var machine [3]Usb // 将数组元素类型声明为 Usb 接口类型
	machine[0] = Phone{Name: "小米"}// 使用指针赋值给 Usb 接口类型的变量
	machine[1] = Phone{Name: "VIVO"}
	machine[2] = Computer{Name: "联想"}
	
	var useMachine UseMachine
	for _,val := range machine {
		useMachine.working(val)
		fmt.Println()
	}
	
}