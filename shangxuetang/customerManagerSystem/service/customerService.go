package service

import (
	"shangxuetang/customerManagerSystem/model"
)

// 该服务完成对customer的操作
type CustomerService struct {
	customers []model.Customer
	customerNum int 
}

// 显示客户列表
func (service *CustomerService) List() []model.Customer{
	return service.customers
}

// 显示客户列表
func (service *CustomerService) Delete(id int) {
	index := service.FindByID(id)
	if index == -1 {
		return
	}
	if len(service.customers) == 1 {
		service.customers = nil
	}else {
		service.customers = append(service.customers[:index],service.customers[index+1:]... )
	}
}

// 删除用户
func (service *CustomerService) FindByID(id int) int{
	for i:= 0;i < len(service.customers);i++ {
		if service.customers[i].Id == id {
			return i
		}
	}
	return -1
}

// 增加用户
func (service *CustomerService) Add(customer model.Customer) {
	service.customerNum++
	customer.Id = service.customerNum
	service.customers = append(service.customers, customer) 
}

// 修改用户

// 返回service
func NewCustomerService() *CustomerService {
	service := CustomerService{}
	// service.customerNum = 1
	
	// 这里先初始化一个customer用户测试显示
	// customer := model.NewCustomer(1,"张三","男",25,"15278944789","8956@gmail.com")
	// service.customers = append(service.customers, customer)
	return &service
}