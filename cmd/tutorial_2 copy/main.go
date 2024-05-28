package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var dominator int = 2
	var result, reminder, err = intDivision(numerator, dominator)
	// if err != nil {
	// 	fmt.Println(err)
	// } else if reminder == 0 {
	// 	fmt.Printf("The result of integer devision is %v", result)
	// } else {
	// 	fmt.Printf("The result of integer devision is %v with reminder %v", result, reminder)
	// }
	switch {
	case err != nil:
		fmt.Println(err)
	case reminder == 0:
		fmt.Printf("The result of integer devision is %v", result)
	default:
		fmt.Printf("The result of integer devision is %v with reminder %v", result, reminder)
	}

	switch reminder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, dominator int) (int, int, error) {
	var err error
	if dominator == 0 {
		err = errors.New("CannotDivideByZero")
		return 0, 0, err
	}
	var result int = numerator / dominator
	var reminder int = numerator % dominator
	return result, reminder, err
}
