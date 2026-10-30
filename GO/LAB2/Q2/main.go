package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("Slice Operations ->")

	Students := []string{"Krish", "Tarun", "Aryan", "Yuvraj"}
	fmt.Println("Original slice:", Students)

	var name string
	fmt.Print("Enter student name to add: ")
	fmt.Scanln(&name)

	Students = append(Students, name)
	fmt.Println("After adding:", Students)

	var index int
	var newName string

	fmt.Print("Enter index to update: ")
	fmt.Scanln(&index)

	fmt.Print("Enter new name: ")
	fmt.Scanln(&newName)

	Students[index] = newName
	fmt.Println("After updating:", Students)


	fmt.Print("Enter index to delete: ")
	fmt.Scanln(&index)

	Students = slices.Delete(Students, index, index+1)
	fmt.Println("After deleting:", Students)

	fmt.Println("\nMap Operations ->")

	marks := map[string]int{
		"Math":    80,
		"Science": 75,
	}

	fmt.Println("Original map:", marks)

	// Insert
	var subject string
	var mark int

	fmt.Print("Enter subject to insert: ")
	fmt.Scanln(&subject)

	fmt.Print("Enter marks: ")
	fmt.Scanln(&mark)

	marks[subject] = mark
	fmt.Println("After inserting:", marks)

	// Delete
	fmt.Print("Enter subject to delete: ")
	fmt.Scanln(&subject)

	delete(marks, subject)
	fmt.Println("After deleting:", marks)

	// Lookup
	fmt.Print("Enter subject to lookup: ")
	fmt.Scanln(&subject)

	fmt.Println("Marks:", marks[subject])
}