package main

import "fmt"

func main() {
	num1, num2, operation := entervalue()
	result, error := calcu(num1, num2, operation)
	if error != nil {
		fmt.Println("error", error)
		return
	}
	fmt.Print(result)
}

func entervalue() (float64, float64, string) {
	var num1 float64
	var num2 float64
	var operation string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter operation: ")
	fmt.Scan(&operation)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	return num1, num2, operation
}

func calcu(num1 float64, num2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Devision by 0 - impossible")
		}
		return num1 / num2, nil
	}
	return 0, fmt.Errorf("unknow operation")
}
