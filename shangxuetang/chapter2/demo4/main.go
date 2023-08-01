package main
import "fmt"

// show char type how to use
func main(){
	var c1 byte = 'a'
	var c2 byte = 'b'

	fmt.Printf("c1 = %c,c2 = %c\n",c1,c2)

	var a int 
	var b float32
	var c float64
	var d string
	var e bool
	// %v 按照变量原值输出
	fmt.Printf("a=%d,b=%v,c=%v,e=%v,d=%v",a,b,c,e,d) 
}