package demo5
import "fmt"

func main(){
	var num int = 9
	fmt.Printf("num address = %v",&num)
	fmt.Println()

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Printf("new num value = %v",*ptr)
}