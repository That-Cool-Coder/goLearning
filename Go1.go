// the first Go experiment

package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputTest()
}

func intts(intVal int) string {
	return strconv.Itoa(intVal)
}

func f64ts(float float64) string {
	str := strconv.FormatFloat(float, 'f', 6, 64)
	return str
}

func inputTest() {
	fmt.Print("CM TO IMPERIAL UNITS CONVERTER\n\n")
	for true {
		fmt.Println("\nEnter cm")

		var cm int;
		fmt.Scan(&cm)

		inches, feet := cmToImperial(float64(cm))
		var toWrite string = intts(cm) + "cm is the same as " + 
			f64ts(inches) + " inches, or " + f64ts(feet) + " feet"
		fmt.Println(toWrite)
	}
}

func cmToImperial(cm float64) (float64, float64) {
	var inches float64 = cm / 2.54
	var feet float64 = cm / 30.48
	return inches, feet
}

func powerTest() {
	for i := 0; i < 5; i ++ {
		var result int = power(i + 1, i + 2)
		var text string = strconv.Itoa(i + 1) + " to the power of " + strconv.Itoa(i + 2) + 
			" is " + strconv.Itoa(result)

		fmt.Println(text)
	}
}

func power(base int, power int) int {
	var total int = base
	for i := 1; i < power; i ++ {
		total *= base
	}
	return total
}