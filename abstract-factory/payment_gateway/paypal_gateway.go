package payment_gateway

import "fmt"

type Paypal struct {
}

func (p Paypal) ProcessPayment(customerID int, amount float64) {
	fmt.Println(fmt.Sprintf("Paypal: customer(%d), amount(%f) -> processed", customerID, amount))
}

func (p Paypal) RefundPayment(customerID int, amount float64) {
	fmt.Println(fmt.Sprintf("Paypal: customer(%d), amount(%f) -> refunded", customerID, amount))
}
