// Write a Go program that launches at least three goroutines concurrently to perform independent tasks (e.g., computing squares, cubes, and Fibonacci numbers).


package main

import (
	"fmt"
)

func main() {
	var n int

	
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	
	ch := make(chan string)

	go func() {
		fmt.Println("Square goroutine started")
		ch <- fmt.Sprintf("Square: %d", n*n)
		fmt.Println("Square goroutine completed")
		}()
		
		go func() {
			fmt.Println("Cube goroutine started")
			ch <- fmt.Sprintf("Cube: %d", n*n*n)
			fmt.Println("Cube goroutine completed")
		}()

	go func() {
    	fmt.Println("Fibonacci goroutine started")
		
    	a, b := 0, 1
		s := "Fibonacci: " 
		for i := 0; i < n; i++ {
			s += fmt.Sprintf("%d ", a)
			a, b = b, a+b 
			} 
		ch <- s
		fmt.Println("Fibonacci goroutine completed")

	}()
	
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) 
	}
}

