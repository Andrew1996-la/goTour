package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Attack struct {
	Mage   string
	Damage int
}

func mage(damageCh chan Attack, ctx context.Context, wg *sync.WaitGroup, mageName string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			randomAttckTime := rand.Intn(800-400+1) + 400
			randomDamage := rand.Intn(25-10+1) + 10
			time.Sleep(time.Duration(randomAttckTime) * time.Millisecond)
			damageCh <- Attack{Mage: mageName, Damage: randomDamage}
		}
	}
}

// Задача: «Арена магов»
func MageAttack() {
	context, cancelContext := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	damageCh := make(chan Attack)
	bossHealth := 300

	wg.Add(3)
	go mage(damageCh, context, wg, "Fire")
	go mage(damageCh, context, wg, "Cold")
	go mage(damageCh, context, wg, "Electirc")

	for {
		if bossHealth <= 0 {
			cancelContext()
			break
		} else {
			msg := <-damageCh

			bossHealth -= msg.Damage
			fmt.Println(msg.Mage, "атаковал босс на", msg.Damage, "урона")
		}
	}

	fmt.Println("Бой окончен")
}
