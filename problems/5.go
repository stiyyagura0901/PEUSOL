package problems

import (
	"fmt"
)

// Five ...
func Five() {
	fmt.Println("2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
	numToCheck := 20

	for {
		// fmt.Println("Checking Number: ", numToCheck)
		count := 0
		for i := 20; i >= 1; i-- {
			if (numToCheck % i != 0) {
				break
			}
			count++
		}
		if (count == 20) {
			fmt.Println(numToCheck)
			break
		}
		numToCheck++
	}
}
