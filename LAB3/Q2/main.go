package main

import "fmt"

// create a function to modify variable using pointer 

func modifyNum(x *int){
	*x =20
}
type Student struct{
	Name string
	Age int
	Job string
}

func modifyStructure(s *Student){
	fmt.Print("Enter Name: ")
	fmt.Scanln(&s.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&s.Age)

	fmt.Print("Enter Job: ")
	fmt.Scanln(&s.Job)
}

func main() {
	// Part 1:
    a := 10
    p := &a
    fmt.Println("Address:", p)

    fmt.Println("Value:", *p)



	// Part 2:
	// value before modification 
	fmt.Println("\nBefore Modification:",a)
	
	// calling function 
	modifyNum(&a)
	
	// value after modification 
	fmt.Println("\nAfter Modification:",a)
	
	// Part 3:
	
	fmt.Println("\nBefore Modification of String:")
	
	// 
	s1 :=new(Student)
	fmt.Println("Name:",s1.Name)
	fmt.Println("Age:",s1.Age)
	fmt.Println("Job:",s1.Job)
	
	modifyStructure(s1)
	
	fmt.Println("\nAfter Modification:")
	fmt.Println("Name:",s1.Name)
	fmt.Println("Age:",s1.Age)
	fmt.Println("Job:",s1.Job)




}