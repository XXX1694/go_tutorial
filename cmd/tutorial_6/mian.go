package main

import "fmt"

func main() {
	// var p *int32 = new(int32)
	// var i int32

	// fmt.Printf("The value p points to is: %v", *p)
	// fmt.Printf("\nThe value if i is: %v", i)

	// p = &i
	// *p = 1

	// fmt.Printf("\n\nThe value p points to is: %v", *p)
	// fmt.Printf("\nThe value if i is: %v", i)

	// var k int32 = 2
	// i = k

	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice
	// sliceCopy[2] = 4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var reult [5]float64 = sqare(thing1)
	fmt.Printf("\nThe result is: %v", reult)
}

func sqare(thing2 [5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
