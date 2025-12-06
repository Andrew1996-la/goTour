package stuctures

import "fmt"

type BankAccount struct {
	Owner    string
	Balance  float64
	Currency string
}

func NewBankOwner(owner, currency string) *BankAccount {
	return &BankAccount{
		Owner:    owner,
		Currency: "RUB",
		Balance:  0,
	}
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		return
	}

	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance-amount < 0 {
		fmt.Println("Недостаточно средств")
	} else {
		b.Balance -= amount
	}
}

func (b BankAccount) GetBalance() {
	fmt.Printf("Owner: %s\nBalance: %.1f %s\n", b.Owner, b.Balance, b.Currency)
}

func BaknTask() {
	andrew := NewBankOwner("Andrew", "RUB")
	andrew.Deposit(150000)
	andrew.Withdraw(140000)
	andrew.GetBalance()
}
