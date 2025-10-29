package main

import (
	"errors"
	"fmt"
	"log"
)

// PaymentProcessor определяет новый интерфейс, который ожидает клиент.
type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) error
}

// LegacyPaymentService обрабатывает платежи в устаревшем строковом формате.
type LegacyPaymentService struct{}

// Pay выполняет платёж, получая данные в виде одной строки(не подходящий формат для клиента).
func (s *LegacyPaymentService) Pay(data string) error {
	if len(data) > 0 {
		log.Printf("Обработка старого платежа: %s\n", data)
		return nil
	} else {
		return errors.New("empty string")
	}
}

// PaymentAdapter преобразует интерфейс PaymentProcessor к формату LegacyPaymentService.
type PaymentAdapter struct {
	service *LegacyPaymentService
}

func NewPaymentAdapter(service *LegacyPaymentService) *PaymentAdapter {
	return &PaymentAdapter{service: service}
}

// ProcessPayment реализует интерфейс PaymentProcessor, адаптируя вызовы к LegacyPaymentService.
func (p *PaymentAdapter) ProcessPayment(amount float64, currency string) error {
	data := fmt.Sprintf("%.2f|%s", amount, currency)
	return p.service.Pay(data)
}

// ProcessOrder - клиентский код, использует PaymentProcessor для обработки оплаты заказа.
func ProcessOrder(processor PaymentProcessor, amount float64, currency string) {
	err := processor.ProcessPayment(amount, currency)
	if err != nil {
		log.Printf("Ошибка оплаты: %v\n", err)
	} else {
		log.Println("Оплата прошла успешно")
	}
}

func main() {
	adapter := NewPaymentAdapter(&LegacyPaymentService{})

	ProcessOrder(adapter, 99.99, "USD")
}
