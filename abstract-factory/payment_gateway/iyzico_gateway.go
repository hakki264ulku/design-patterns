package payment_gateway

import "fmt"

type Iyzico struct {
}

func (i Iyzico) ProcessPayment(customerID int, amount float64) {
	fmt.Println(fmt.Sprintf("Iyzico: customer(%d), amount(%f) -> processed", customerID, amount))
}

func (i Iyzico) RefundPayment(customerID int, amount float64) {
	fmt.Println(fmt.Sprintf("Iyzico: customer(%d), amount(%f) -> refunded", customerID, amount))
}
