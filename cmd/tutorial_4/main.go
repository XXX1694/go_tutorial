package main

import (
	"fmt"
)

func main() {
	var myString = "resume"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
}
