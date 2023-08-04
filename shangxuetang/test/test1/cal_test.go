package test

import (
	"os"
	"strconv"
	"testing"
)

/**
test all file : go test -v
test single function  :go test -v -test.run TestXXX
test single file :go test -v xxx_test.go xxx.go
**/
func TestAdd(t *testing.T) {
	args := []string{"1","7"}
	length := len(args)
	if length <2 {
		t.Fatalf("input args length less than 2")
		os.Exit(1)
	}
	turnToNums := make([]int,length)
	for i := 0; i < length; i++ {
		num,err := strconv.Atoi(args[i])
		if err != nil {
			t.Fatalf("transfer input string to num error :%v\n",err)
		}
		turnToNums[i] = num
	}

	res := Add(turnToNums[0],turnToNums[1])
	if res != 9 {
		t.Fatalf("input error")
	}
	t.Logf("input success")
}