package problems

import "fmt"

var num = 600851475143

// Get all prime factors of a given number n
func primeFactors(n int) (pfs []int) {
	for n % 2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n % i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}
	return
}

// Three ...
func Three() {
	fmt.Println("The prime factors of 13195 are 5, 7, 13 and 29. What is the largest prime factor of the number 600851475143 ?")
	
	pfs := []int {}
	pfs = primeFactors(num)

	fmt.Println(pfs)
}
