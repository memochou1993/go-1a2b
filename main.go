package main

import (
	"fmt"
	"math/rand"
	"time"
)

func contains(needle int, haystack []int) bool {
	for _, ele := range haystack {
		if needle == ele {
			return true
		}
	}

	return false
}

func random(number int) []int {
	var numbers []int

	rand.Seed(time.Now().Unix())

	i := 0
	for i < number {
		number := rand.Intn(10)

		if number > 0 && !contains(number, numbers) {
			numbers = append(numbers, number)
			i++
		}
	}

	return numbers
}

func explode(number int) []int {
	var array []int

	for number >= 1 {
		array = append([]int{number % 10}, array...)

		number /= 10
	}

	return array
}

func compare(expected []int, actual []int) (int, int) {
	var a int
	var b int

	for expectedIndex, expectedValue := range expected {
		for actualIndex, actualValue := range actual {
			if expectedValue == actualValue {
				if expectedIndex == actualIndex {
					a++
					break
				}
				b++
			}
		}
	}

	return a, b
}

func main() {
	digits := 4

	actual := random(digits)

	var input int

	for {
		fmt.Scan(&input)

		expected := explode(input)

		if len(expected) != digits {
			fmt.Println("Invalid input")
			continue
		}

		a, b := compare(expected, actual)

		if a == digits {
			break
		}

		fmt.Printf("%dA%dB\n", a, b)
	}
}
