package main

import "fmt"

func main() {
	var intNum int64 = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var myString string = "huh"
	fmt.Println(myString)

	var1, var2, var3 := 1, 2, 3
	fmt.Println(var1, var2, var3)

	const pi float64 = 3.14567
	fmt.Println(pi)
}
