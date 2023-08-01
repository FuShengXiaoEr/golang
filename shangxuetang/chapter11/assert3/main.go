package main

import "fmt"


func JudgeType(items ...interface{}) {
	for i,item := range items {
		switch item.(type){
			case bool:
				fmt.Printf("第%d个参数是bool类型，值是%v\n",i,item)
			case float64,float32:
				fmt.Printf("第%d个参数是float类型，值是%v\n",i,item)
			case int,int64:
				fmt.Printf("第%d个参数是int类型，值是%v\n",i,item)
			case string:
				fmt.Printf("第%d个参数是string类型，值是%v\n",i,item)
			case nil:
				fmt.Printf("第%d个参数是nil类型，值是%v\n",i,item)
			default :
				fmt.Printf("第%d个参数没有对应类型，值是%v\n",i,item)
			
		}
	}
}
func main(){
	var n1 float64 = 1.6
	var n2 float32 = 23.3
	var n3 int = 12
	var n4 bool = true
	var n5 string = "ss"
	JudgeType(n1,n2,n3,n4,n5)

}