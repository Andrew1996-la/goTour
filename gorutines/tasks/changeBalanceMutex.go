package tasks

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
	mtx     sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.mtx.Lock()
	a.balance += amount
	a.mtx.Unlock()
}

func (a *Account) Withdraw(amount int) {
	a.mtx.Lock()
	a.balance -= amount
	a.mtx.Unlock()
}

func ChangeBalance() {
	account := Account{
		balance: 1000,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				account.Deposit(2)
				account.Withdraw(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final balance:", account.balance)
}
