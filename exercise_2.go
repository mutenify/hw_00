package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64)
}

type Card struct {
	cardNumber int
	cardBank   string
	deposit    float64
}

func (c *Card) Pay(amount float64) {
	if amount <= c.deposit {
		fmt.Printf("Списано %f рублей с карты %d (%s) \n", amount, c.cardNumber, c.cardBank)
		c.deposit -= amount
	} else {
		fmt.Println("Недостаточно средств!")
	}
}

type Crypto struct {
	walletNumber int
	currency     string
	deposit      float64
	worth        float64
}

func (c *Crypto) Pay(amount float64) {
	amount /= c.worth
	if amount <= c.deposit {
		fmt.Printf("Списано %f %s с кошелька %d \n", amount, c.currency, c.walletNumber)
		c.deposit -= amount
	} else {
		fmt.Println("Недостаточно средств!")
	}
}

type Order struct {
	ID            int
	cost          float64
	PaymentMethod PaymentMethod
}

func ProcessOrder(o Order) {
	fmt.Println("Обработка заказа", o.ID)
	o.PaymentMethod.Pay(o.cost)
}

func main() {
	wallet1 := Crypto{walletNumber: 1, currency: "ETH", deposit: 14, worth: 4424}
	order1 := Order{ID: 1, cost: 42888, PaymentMethod: &wallet1}
	ProcessOrder(order1)
	fmt.Println(wallet1.deposit)
}
