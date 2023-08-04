package test2

import (
	"shangxuetang/test/test2/model"
)

func serialMonster(){
	monster := model.MonsterModel{
		Name:"aa",
		Age:12,
		Skill:"shot",
	}
	monster.Store()
	monster.ReStore()
	
}

