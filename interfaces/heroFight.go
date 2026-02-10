package interfaces

import (
	"fmt"
	"math/rand"
)

type Fighter interface {
	Attack() int
	TakeDamage(damage int)
	IsAlive() bool
	GetName() string
	GetHealth() int
}

type BaseHero struct {
	Name   string
	Health int
	Damage int
}

type Warrior struct {
	BaseHero
}

func (w Warrior) Attack() int {
	return w.Damage
}
func (w *Warrior) TakeDamage(damage int) {
	w.Health -= damage - 2
}
func (w Warrior) IsAlive() bool {
	if w.Health > 0 {
		return true
	}
	return false
}
func (w Warrior) GetName() string {
	return w.Name
}
func (w Warrior) GetHealth() int {
	return w.Health
}

type Mage struct {
	BaseHero
}

func (m Mage) Attack() int {
	if rand.Intn(2) == 0 {
		return m.Damage * 2
	}

	return m.Damage
}
func (m *Mage) TakeDamage(damage int) {
	m.Health -= damage
}
func (m Mage) IsAlive() bool {
	if m.Health > 0 {
		return true
	}
	return false
}
func (m Mage) GetName() string {
	return m.Name
}
func (m Mage) GetHealth() int {
	return m.Health
}

func Fight(a, b Fighter) {
	round := 1
	for a.IsAlive() && b.IsAlive() {
		if round%2 == 0 {
			damage := a.Attack()

			fmt.Println(a.GetName(), "атакует", b.GetName(), "на", damage)
			b.TakeDamage(damage)
			if !b.IsAlive() {
				fmt.Println(b.GetName(), "умер в раунде", round)
				break
			}
			fmt.Println(b.GetName(), "осталось", b.GetHealth(), "HP")
		}

		damage := a.Attack()
		fmt.Println(b.GetName(), "атакует", a.GetName(), "на", damage)
		a.TakeDamage(damage)
		if !a.IsAlive() {
			fmt.Println(a.GetName(), "умер в раунде", round)
			break
		}
		fmt.Println(a.GetName(), "осталось", a.GetHealth(), "HP")

		round++
	}
}

func testFight() {
	warrior := Warrior{BaseHero{
		Name:   "Warrior",
		Health: 50,
		Damage: 10,
	}}

	mage := Mage{BaseHero{
		Name:   "Mage",
		Health: 40,
		Damage: 7,
	}}

	Fight(&warrior, &mage)
}
