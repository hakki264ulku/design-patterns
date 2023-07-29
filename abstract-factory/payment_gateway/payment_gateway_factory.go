package payment_gateway

type PaymentGatewayFactory struct {
}

func (p PaymentGatewayFactory) GetPaymentGateway(gatewayName GatewayName) PaymentGateway {
	switch gatewayName {
	case PaypalName:
		return Paypal{}
	case IyzicoName:
		return Iyzico{}
	default:
		return nil
	}
}
