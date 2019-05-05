package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println("EVEN", i)
		} else {
			fmt.Println("ODD", i)
		}

	}

	i := 0

	switch i {
	case 0:
		fmt.Println("ZERO")
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	case 4:
		fmt.Println("FOUR")
	case 5:
		fmt.Println("FIVE")
	default:
		fmt.Println("UNKNOWN")
	}

	fmt.Println("---------------------------------")
	divisableByThree(1, 20)
	fmt.Println("--------------------------------")
	problem3()
}

func getNearestDivisable(num int, by int) int {
	for i := num; i <= num+by; i++ {
		if i%by == 0 {
			return i
		}
	}
	return -1
}

func divisableByThree(min int, max int) {
	for i := getNearestDivisable(min, 3); i <= max; i += 3 {
		fmt.Println(i)
	}
}

func problem3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
