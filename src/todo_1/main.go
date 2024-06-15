package main

import "fmt"

type gasEngine struct {
	mpg     int8
	gallons int8
}

func (gasEngine gasEngine) milesLeft() int8 {
	return gasEngine.mpg * gasEngine.gallons
}

type engine interface {
	milesLeft() int8
}

func canMakeIt(e engine, miles int8) bool {
	return e.milesLeft() >= miles
}

func main() {

	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 5}
	fmt.Println(canMakeIt(myEngine, 100))

	var intNum int = 3200
	fmt.Println(intNum)

	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	var myString string = "Hello, World!" + " I am a string!"
	fmt.Println(myString)

	myBool := true
	fmt.Println(myBool)

	const myConst string = "I am a constant"
	fmt.Println(myConst)

	printMe("Hello, World!")
	addedNum := add(1, 2)

	fmt.Println(addedNum)
}

func printMe(stringToPrint string) {
	fmt.Println(stringToPrint)
}

func add(a int, b int) int {
	return a + b
}
