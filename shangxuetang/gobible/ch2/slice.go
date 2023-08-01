package main

import "fmt"

func main1(){

	months := [...]string{
	1: "January",
	2: "February",
	3: "March",
	4: "April",
	5: "May",
	6: "June",
	7: "July", 
	8: "August",
	9: "September",
	10: "October",
	11: "November",
	12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:12]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	reverse()
	appendtest()

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempy2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

}

func main(){
	arr := [5]string{"1","2","3","4","5"}
	fmt.Println(arr)
	reverse2(&arr)
	fmt.Println(arr)
}
func reverse(){
	a := [...]int{0,1,2,3,4,5}
	for i,j := 0,len(a)-1; i < j; i,j = i + 1,j-1 {
		a[i],a[j] = a[j],a[i]
	}
	fmt.Println(a)
}

func reverse2(arr *[5]string) {
	for i,j := 0,len(arr) - 1;i < j;i ,j = i+1,j -1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
}

func appendtest(){
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)

	}

	fmt.Printf("%q\n",runes)
}

func appendInt(x []int, y int)[]int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen,zcap)
		copy(z,x)
	}
	z[len((x))] = y
	return z
	
}

func nonempy (strings []string)[]string{
	count := 0
	for _, s := range strings{
		if s != "" {
			strings[count] = s
			count++;
		}
	}
	return strings[:count]
}

func nonempy2 (strings []string)[]string{
	count := strings[:0]
	for _,s := range strings {
		if s != "" {
			count = append(count, s)
		}
	}
	return count
}