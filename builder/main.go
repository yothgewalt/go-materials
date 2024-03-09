package main

import (
	"fmt"
)

type Order struct {
	Items     []MenuItem
	Customer  *Customer
	Delivery  *Delivery
	Payment   *Payment
	TotalCost float64
}

type MenuItem struct {
	Name        string
	Price       float64
	Description string
}

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Delivery struct {
	Address string
	Fee     float64
}

type Payment struct {
	Method  string
	Details string
}

type OrderBuilder struct {
	order Order
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{
		order: Order{},
	}
}

func (b *OrderBuilder) WithItems(items []MenuItem) *OrderBuilder {
	b.order.Items = items
	return b
}

func (b *OrderBuilder) WithCustomer(customer Customer) *OrderBuilder {
	b.order.Customer = &customer
	return b
}

func (b *OrderBuilder) WithDelivery(delivery Delivery) *OrderBuilder {
	b.order.Delivery = &delivery
	return b
}

func (b *OrderBuilder) WithPayment(payment Payment) *OrderBuilder {
	b.order.Payment = &payment
	return b
}

func (b *OrderBuilder) Build() (Order, error) {
	if len(b.order.Items) == 0 {
		return Order{}, fmt.Errorf("order must have at least one item")
	}

	b.order.TotalCost = 0
	for _, item := range b.order.Items {
		b.order.TotalCost += item.Price
	}

	if b.order.Delivery != nil {
		b.order.TotalCost += b.order.Delivery.Fee
	}

	return b.order, nil
}

func main() {
	builder := NewOrderBuilder()
	order, err := builder.
		WithItems([]MenuItem{
			{Name: "Pizza", Price: 10.00, Description: "Large pepperoni pizza"},
			{Name: "Soda", Price: 2.00, Description: "12 oz. can of soda"},
		}).
		WithCustomer(Customer{Name: "John Doe", Address: "123 Main St", Phone: "123-456-7890"}).
		WithDelivery(Delivery{Address: "456 Elm St", Fee: 5.00}).
		WithPayment(Payment{Method: "Credit Card", Details: "1234-5678-9012-3456"}).
		Build()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order details:")
	fmt.Println("Items:", order.Items)
}
