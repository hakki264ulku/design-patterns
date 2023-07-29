package main

import "abstract_factory/payment_gateway"

type PaymentProcessor struct {
	Instructions          []Instruction
	PaymentGatewayFactory payment_gateway.PaymentGatewayFactory
}

func (p *PaymentProcessor) ProcessPaymentInstructions() {
	for _, instruction := range p.Instructions {
		gateway := p.PaymentGatewayFactory.GetPaymentGateway(payment_gateway.GatewayName(instruction.GatewayName))
		switch instruction.ProcessType {
		case "process":
			gateway.ProcessPayment(instruction.CustomerID, instruction.Amount)
		case "refund":
			gateway.RefundPayment(instruction.CustomerID, instruction.Amount)
		}
	}
}
