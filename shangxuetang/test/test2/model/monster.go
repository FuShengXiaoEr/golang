package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MonsterModel struct {
	Name  string
	Age   int
	Skill string
}

func (monster *MonsterModel) Store() {

	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Errorf("json format error :%v\n", err)
	}

	wFile, err := os.OpenFile("abc.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Errorf("创建文本失败，error：%v\n", err)
	}
	defer wFile.Close()
	write := bufio.NewWriter(wFile)

	write.WriteString(string(jsonStr) + "\n")
	write.Flush()
}

func (mon *MonsterModel) ReStore() MonsterModel {
	rFile, err := os.Open("abc.txt")
	if err != nil {
		fmt.Errorf("read file error :%v\n", err)
	}
	defer rFile.Close()
	read := bufio.NewReader(rFile)

	for {
		str, err := read.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("pharse file to string:  %v\n", str)
	}
	return MonsterModel{}

}

func (mon *MonsterModel) ReStore2() bool {
	content, err := ioutil.ReadFile("abc.txt")
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return false
	}

	// Convert the byte slice to a string and print its contents
	err = json.Unmarshal(content, mon)
	if err != nil {
		fmt.Errorf("unmarshal error :%v\n", err)
		return false
	}
	return true
}
