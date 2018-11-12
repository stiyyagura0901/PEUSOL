package problems

import (
	"fmt"
)

func isPrime(value int) bool {
	if value == 2 { return true }
	if value == 3 { return true }
	if value % 2 == 0 { return false }
	if value % 3 == 0 { return false }

	i := 5;
	w := 2;

	for {
		if (i*i <= value) {
			if value % i == 0 {
				return false
			}
			i += w
			w = 6 - w
		} else {
			break
		}
	}

	return true
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
