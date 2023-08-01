package main

import (
	"fmt"
	"strings"
)

func hasSuffix(suffix string) func (string) string  {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main(){
	f := hasSuffix(".jpg")
	fmt.Println("文件名处理后=", f("name"))
	fmt.Println("文件名处理后=",f("name.avi"))
}