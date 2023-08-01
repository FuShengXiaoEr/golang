package model	

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 工厂模式返回一个Customer实例子
func NewCustomer() Customer {

}