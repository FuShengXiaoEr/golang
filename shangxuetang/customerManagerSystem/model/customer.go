package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func (customer Customer) PrintToTerminal() string{
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",customer.Id,customer.Name,
	customer.Gender,customer.Age,customer.Phone,customer.Email)
}

// 工厂模式返回一个Customer实例子
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2( name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
