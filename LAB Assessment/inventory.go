package main

import (
	"errors"
	"fmt"
)


type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var items []Item

type Summary struct {
	ItemCount  int
	TotalValue float64
}

func AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be positive")
	}

	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}


	items = append(items, item)
	return nil
}

func UpdateItem(id int, name string, quantity int, price float64) error {
	if id <= 0 {
		return errors.New("item ID must be positive")
	}

	if name == "" {
		return errors.New("item name cannot be empty")
	}

	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}

	if price < 0 {
		return errors.New("price cannot be negative")
	}

	for i := range items {
		if items[i].ID == id {
			items[i].Name = name
			items[i].Quantity = quantity
			items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

func RemoveItem(id int) error {
	if id <= 0 {
		return errors.New("item ID must be positive")
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

func DisplayItems() {
	if len(items) == 0 {
		fmt.Println("Inventory is empty")
		return
	}

	fmt.Println("\n--- Inventory ---")

	for _, item := range items {
		fmt.Printf(
			"ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID,
			item.Name,
			item.Quantity,
			item.Price,
		)
	}
}

func CalculateSummary() Summary {
	var summary Summary

	for _, item := range items {
		summary.ItemCount++
		summary.TotalValue += float64(item.Quantity) * item.Price
	}

	return summary
}

func GetSummary() Summary {
	result := make(chan Summary)

	go func() {
		summary := CalculateSummary()
		result <- summary
	}()

	summary := <-result

	return summary
}

