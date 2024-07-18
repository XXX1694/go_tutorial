package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// type owner struct {
// 	name string
// }

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type engine interface {
	milesLeft() uint8
}

func main() {
	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 25}
	var myEngine1 electricEngine = electricEngine{mpkwh: 12, kwh: 45}
	// myEngine.mpg = 12
	// fmt.Println(myEngine)

	// fmt.Println("Total miles legt in tank: ", myEngine.milesLeft())

	canMakeIt(myEngine1, 50)
}
