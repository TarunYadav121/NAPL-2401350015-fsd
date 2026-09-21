package main
import "fmt"

type Student struct{
	Name string
	Age int
	Marks float64
}

func main(){
	Tarun := Student{
		Name : "Tarun",
		Age : 20,
		Marks : 105.0,
	}
	fmt.Printf("Name of student : %s \n", Tarun.Name)
	fmt.Printf("Age of student : %d \n", Tarun.Age)
	fmt.Printf("Marks of student : %.2f \n", Tarun.Marks)
}