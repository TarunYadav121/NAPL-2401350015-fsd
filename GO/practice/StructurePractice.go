package main

import (
"errors"
"fmt"
"strconv"
)

func read() (int, error) {
return 0, errors.New("failed")
}

func main() {
v, err := strconv.Atoi("42")


if err != nil {
	fmt.Println("bad input:", err)
	return
}

fmt.Println(v)


}

// type Student struct{
// 	Name string
// 	Age int
// 	Marks float64
// }

// func main(){
// 	Tarun := Student{
// 		Name : "Tarun",
// 		Age : 20,
// 		Marks : 105.0,
// 	}
// 	fmt.Printf("Name of student : %s \n", Tarun.Name)
// 	fmt.Printf("Age of student : %d \n", Tarun.Age)
// 	fmt.Printf("Marks of student : %.2f \n", Tarun.Marks)
// }