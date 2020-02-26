package main

import "fmt"

func Thirt(n int) int {
	var factors = [6]int{1, 10, 9, 12, 3, 4}
	var result int

	for {
		result = 0
		tn := n

		for i := 0; tn > 0; i++ {
			result += (tn % 10) * factors[i%len(factors)]
			tn /= 10
		}

		if n == result {
			break
		}

		n = result
	}

	return result
}

func main() {
	fmt.Print(Thirt(1234567))
}
