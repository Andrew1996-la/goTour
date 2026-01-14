package composition

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Add(amount int) {
	w.balance += amount
}

func (w *Wallet) Spend(amount int) {
	w.balance -= amount
}

type user struct {
	name string
	Wallet
}

func UserWithWallet() {
	user := user{name: "Boba"}

	user.Add(100)
	user.Spend(30)
	fmt.Println(user.balance)
}
