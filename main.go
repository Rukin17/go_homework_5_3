package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) float64 {
	b.Balance += amount
	return b.Balance
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
	var err error
	if amount > b.Balance {
		err = errors.New("На вашем счете недостаточно средств")
	} else {
		b.Balance -= amount
	}
	return b.Balance, err
}

func main() {
	client := BankAccount{
		Owner:   "John Doe",
		Balance: 100.0,
	}

	// Увеличиваем баланс
	fmt.Println("Пополнение счета: ", client.Deposit(100.0))

	// Позитивный тест снятия со счета
	checkCorrectWithdraw(client, 100)

	// Негативный тест снятия со счета
	checkCorrectWithdraw(client, 250)

}

func checkCorrectWithdraw(client BankAccount, amount float64) {
	if balance, err := client.Withdraw(amount); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Успешно! Баланс: ", balance)
	}
}
