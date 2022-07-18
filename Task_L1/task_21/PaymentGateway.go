package main

import "fmt"

type Account struct {
	accNumber int
	name      string
	phone     string
	balance   int
}

type PaymentGateway interface {
	doPayment(account1, account2 Account)
}

type PaymentGatewayImpl struct {
}

func (acc PaymentGatewayImpl) doPayment(account1, account2 Account) {
	fmt.Printf("Payment Gateway: Перевод с аккаунта %s на аккаунт %s", account1.name, account2.name)
}
