package payment_gateway

type PaymentGateway interface {
	ProcessPayment(customerID int, amount float64)
	RefundPayment(customerID int, amount float64)
}
