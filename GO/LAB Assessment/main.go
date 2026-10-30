package main

import "fmt"

func main() {

	err := AddItem(Item{
		ID:       1,
		Name:     "Laptop",
		Quantity: 5,
		Price:    50000,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(Item{
		ID:       2,
		Name:     "Mouse",
		Quantity: 10,
		Price:    500,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = AddItem(Item{
		ID:       3,
		Name:     "Keyboard",
		Quantity: 4,
		Price:    1000,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	DisplayItems()

	err = UpdateItem(2, "Wireless Mouse", 8, 700)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nItem updated successfully")
	}

	DisplayItems()

	err = RemoveItem(3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nItem removed successfully")
	}

	DisplayItems()

	summary := GetSummary()

	fmt.Println("\n--- Inventory Summary ---")
	fmt.Println("Total Items:", summary.ItemCount)
	fmt.Printf("Total Inventory Value: %.2f\n", summary.TotalValue)
}