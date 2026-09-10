package math_util;

// func Add(a int, b int) int {
// 	Sum:=a+b
// 	return Sum
// }

func Factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}


func Power(a int, b int) int {
	result := 1

	for i := 1; i <= b; i++ {
		result *= a
	}

	return result
}