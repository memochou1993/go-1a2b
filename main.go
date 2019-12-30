package main

import (
	"fmt"
)

func explode(number int) []int {
	var array []int

	for number%10 != 0 {
		array = append([]int{number % 10}, array...)

		number /= 10
	}

	return array
}

func compare(expected int, actual int) (int, int) {
	var a int
	var b int

	expectedArray := explode(expected)
	actualArray := explode(actual)

	for expectedIndex, expectedValue := range expectedArray {
		for actualIndex, actualValue := range actualArray {
			if expectedIndex == actualIndex {
				if expectedValue == actualValue {
					a++
					break
				}
				b++
				break
			}
		}
	}

	return a, b
}

func main() {
	expected := 4579
	actual := 7945

	a, b := compare(expected, actual)

	fmt.Printf("%dA%dB", a, b)
}
