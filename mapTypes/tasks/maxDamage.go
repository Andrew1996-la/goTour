package tasks

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Attack struct {
	Player string
	Damage int
	Target string
}

func MaxDamage() {
	generalPlayrDamage := make(map[string]int)
	generalMonstersDamage := make(map[string]int)

	attacks := []Attack{
		{"Mage", 30, "Orc"},
		{"Warrior", 20, "Orc"},
		{"Mage", 25, "Goblin"},
		{"Archer", 15, "Orc"},
		{"Mage", 10, "Orc"},
		{"Warrior", 40, "Goblin"},
	}

	var maxDamageName string
	var maxDamageCount int

	// Посчитать общий урон по каждому игроку
	// Посчитать общий урон по каждому монстру
	// Найти игрока, который нанёс максимальный суммарный урон
	for _, player := range attacks {
		generalPlayrDamage[player.Player] += player.Damage
		generalMonstersDamage[player.Target] += player.Damage
	}

	for preson, damage := range generalPlayrDamage {
		if damage > maxDamageCount {
			maxDamageName = preson
			maxDamageCount = damage
		}
	}

	pp.Println(generalPlayrDamage)
	pp.Println(generalMonstersDamage)
	fmt.Printf("Больше всего урона нанес %s:%d\n", maxDamageName, maxDamageCount)
}
