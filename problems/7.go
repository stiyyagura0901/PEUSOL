package problems

import (
	"fmt"
	"math"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// Seven ...
func Seven() {
	fmt.Println(`By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

	What is the 10 001st prime number?`)

	nthPrime := 10001
	i := 1
	count := 0

	for {
		fmt.Println(i)
		if isPrime(i) {
			count++
			if (count == nthPrime) {
				fmt.Println("Found:", i)
				break
			}
		}
		
		i++
	}
}
