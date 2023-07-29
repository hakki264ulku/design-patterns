package main

import "abstract_factory/payment_gateway"

// This is a project for studying abstract-factory pattern
// Gateway implementation with abstract_factory

func main() {
	instructions := []Instruction{
		{
			CustomerID:  1234,
			Amount:      99.9,
			GatewayName: "paypal",
			ProcessType: "process",
		},
		{
			CustomerID:  3214,
			Amount:      49.9,
			GatewayName: "iyzico",
			ProcessType: "refund",
		},
	}

	paymentGatewayFactory := payment_gateway.PaymentGatewayFactory{}

	paymentProcessor := PaymentProcessor{
		Instructions:          instructions,
		PaymentGatewayFactory: paymentGatewayFactory,
	}

	paymentProcessor.ProcessPaymentInstructions()
}
