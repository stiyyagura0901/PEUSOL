package problems

import (
	"fmt"
	"sort"
)

func isPalin(i int) bool {
	var reverse, remainder int

	temp := i

	for {
		remainder = i % 10
		reverse = reverse * 10 + remainder
		i /= 10

		if (i == 0) {
			break
		}
	}

	return (temp == reverse)
}

// Four ...
func Four() {
	fmt.Println("A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99. Find the largest palindrome made from the product of two 3-digit numbers.")
	
	arr := []int{}

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			mul := j * i
			if (isPalin(mul)) {
				arr = append(arr, mul)
			}
		}
	}

	sort.Ints(arr)
	fmt.Println(arr[0])
}
