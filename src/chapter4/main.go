package main

import "fmt"

func main() {
	var x = "Hello, World"
	fmt.Println(x)

	y := "ramadan"
	fmt.Println(y)

	const PI float64 = 3.14
	fmt.Println(PI)

	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)

	fmt.Println("Please Enter a number: ")
	var num float64
	fmt.Scanf("%f", &num)
	println("the number doubled will be", double(num))
}

func double(num float64) float64 {
	return 2 * num
}

func fromFahrenheitIntoCelsius(temprature float64) float64 {
	// (C = (F - 32) * 5/9)
	return (temprature - 32) * 5 / 9
}

func fromFeetToMeter(feet float64) float64 {
	// (1 ft = 0.3048 m)
	return feet * 0.3048
}
