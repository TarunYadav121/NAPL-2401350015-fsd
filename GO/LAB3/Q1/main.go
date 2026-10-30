package main

import "fmt"

type Person struct {
	Name   string
	Job    string
	Age    int
	Salary float64
}

func (p Person) readdata() Person {
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scanln(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)

	return p
}

func (p Person) PrintData() {
	fmt.Println("Name   :", p.Name)
	fmt.Println("Age    :", p.Age)
	fmt.Println("Job    :", p.Job)
	fmt.Println("Salary :", p.Salary)
}

func main() {
	var p1 Person
	var p2 Person

	fmt.Println("Enter details of person 1")
	p1 = p1.readdata()

	fmt.Println("\nDetails of person 1")
	p1.PrintData()

	fmt.Println("\nEnter details of person 2")
	p2 = p2.readdata()

	fmt.Println("\nDetails of person 2")
	p2.PrintData()
}