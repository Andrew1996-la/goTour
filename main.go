package main

import (
	"errors"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func (u *User) Pay(sum int) error {
	if u.Balance-sum < 0 {
		return errors.New("недостаточно средств")
	}

	u.Balance -= sum
	return nil
}

func main() {
	user := User{
		Name:    "Vladimir",
		Balance: 10,
	}

	pp.Println(user)

	err := user.Pay(50)

	if err != nil {
		pp.Println(err.Error())
	} else {
		pp.Println("оплата успешно прошла")
	}

	pp.Println(user)
}
