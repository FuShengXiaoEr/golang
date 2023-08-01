package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var i int = 10
	i = 30
	i = 50
	fmt.Println(i)

	// format output
	var v1 int = 10
	fmt.Printf("v1 variable type is : %T\n",v1)

	var v2 int64 = 10;
	fmt.Printf("v2 variable type is :%T,occupy byte number is %d",v2,unsafe.Sizeof(v2))
}