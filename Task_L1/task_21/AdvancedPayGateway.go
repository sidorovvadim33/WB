package main

type AdvancedPayGateway interface {
	makePayment(mobileSender, mobileReceiver string)
}

type AdvancedPayGatewayAdapter struct {
}

func (adapter *AdvancedPayGatewayAdapter) makePayment(mobileSender, mobileReceiver string) {
	//Получаем аккаунт по номеру телефона далее перевод стандартный
	account1 := Account{}
	account2 := Account{}
	PaymentGatewayImpl{}.doPayment(account1, account2)
}
