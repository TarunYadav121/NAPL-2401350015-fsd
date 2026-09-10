package main

import "fmt"
import "Myproject/math_util"
import "Myproject/strop"

func main()  {
	// var a int
	// var b int
	// fmt.Println("enter number 1: ")
	// fmt.Scan(&a)
	// fmt.Println("enter number 2: ")
	// fmt.Scan(&b)
	// fmt.Printf("Sum is = %d", math_util.Add(a,b))
	// var s string
	// fmt.Println("Enter your name : ")
	// fmt.Scan(&s)
	// fmt.Println("Let me guess your name 🧐🧐🧐!")
	// fmt.Printf("Your name is : %s",strop.PrintStr(s))


	var s string

	fmt.Print("Enter a string: ")
	fmt.Scan(&s)

	fmt.Println("Reversed string:", strop.Reverse(s))
	fmt.Println("Number of vowels:", strop.CountVowels(s))

	var n, a, b int
	fmt.Print("Enter a number for factorial: ")
	fmt.Scan(&n)
	fmt.Printf("Factorial is = %d \n", math_util.Factorial(n))

	fmt.Println("Enter base (a): ")
	fmt.Scan(&a)

	fmt.Println("Enter exponent(b): ")
	fmt.Scan(&b)
	fmt.Printf("Sum is = %d", math_util.Power(a,b))

	
}