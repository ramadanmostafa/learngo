package main

import "fmt"

func main() {
	x := 10
	zero(&x)
	fmt.Println(x)

	x = 100
	xPtr := new(int)
	zero(xPtr)
	fmt.Println(*xPtr)

	y := 1.5
	square(&y)
	fmt.Println(y)

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func zero(num *int) {
	*num = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
