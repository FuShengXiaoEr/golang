package test2

import (
	"os"
	"shangxuetang/test/test2/model"
	"testing"
)

func TestXxx(t *testing.T) {
	item := model.MonsterModel{
		Name: "aa",
		Age: 12,
		Skill: "shot",
	}

	item.Store()
	
	// 检查文件是否存在
	_,err := os.Stat("abc.txt")
	if os.IsNotExist(err) {
		t.Fatal("store fail,don't find file :")
	}else {
		t.Log("store success")
	}

	// 反序列化
	item.ReStore2()
	
	if item.Age == 12 && item.Name== "aa"&& item.Skill == "shot" {
		t.Log("success")
	}else {
		t.Fatal("restore fail")
	}
}