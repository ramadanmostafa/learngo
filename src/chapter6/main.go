package main

import "fmt"

func main() {
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	total := 0.0
	for _, val := range x {
		total += val
	}
	fmt.Println(total / float64(len(x)))

	slices()

	maps()

	// smallest number
	fmt.Println(smallestNumber())
}

func slices() {
	// create a slice that is associated with an underlying float64 array of length 5
	x1 := make([]float64, 5)

	// another way to create a slice
	arr := [5]float64{1, 2, 3, 4, 5}
	x2 := arr[1:4]

	fmt.Println(x1, x2)

	// append and copy slices
	slice1 := []float64{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)

	fmt.Println(slice1, slice2)

	slice3 := []float64{1, 2, 3}
	slice4 := make([]float64, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}

func maps() {
	// var x map[string]int
	x := make(map[string]int)
	x["key"] = 12
	x["key2"] = 122
	x["key3"] = 1222

	fmt.Println(x)
	delete(x, "key")
	fmt.Println(x)

	name, ok := x["Un"]
	fmt.Println(name, ok)

	if name, ok := x["Un"]; ok {
		fmt.Println(name, ok)
	}

}

func smallestNumber() int {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := 200
	for _, val := range x {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}
