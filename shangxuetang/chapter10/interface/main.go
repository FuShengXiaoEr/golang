package main

import "fmt"

type AInterface interface {
	Say()
	Talk()
}

type Stu struct {
	Name string
}

type AType int;

func (stu Stu)Say(){
	fmt.Println(stu.Name)
}

func (aType AType)Say(){
	fmt.Println("aType")
}

func main(){
	stu := Stu {
		Name: "aa",
	}

	stu.Say()

	var aType AType
	aType.Say()
}