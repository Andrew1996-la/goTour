package errorhanlers

import (
	"errors"
	"math/rand"

	"github.com/k0kubun/pp"
)

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	kmch := rand.Intn(150)

	if c.Armor-10 <= 0 {
		return 0, errors.New("у машины недостачно прочености. обратитесь в сервис")
	}

	c.Armor -= 10

	return kmch, nil
}

func CarHandler() {
	zhiga := Car{Armor: 23}

	for {
		kmch, err := zhiga.Gas()
		if err != nil {
			pp.Println("Разгон 0.", err.Error())
			break
		} else {
			pp.Println("Машина разогналась до: ", kmch)
		}
	}

}
