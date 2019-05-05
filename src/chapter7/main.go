package main

import "fmt"

func main() {
	lst := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(average(lst))
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(add())

	lstInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(add(lstInt...))

	addTwoNum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(addTwoNum(1, 9))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	fmt.Println(factorial(4))

	defer second() // will be called after the main function is done
	first()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5

	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))

	panic("PANIC")
}

func average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}

	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func makeOddGenerator() func() (ret uint) {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
