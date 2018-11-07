package problems

import (
	"fmt"
	"math"
)

// Six ...
func Six() {
	fmt.Println(`The sum of the squares of the first ten natural numbers is,

	12 + 22 + ... + 102 = 385
	The square of the sum of the first ten natural numbers is,
	
	(1 + 2 + ... + 10)2 = 552 = 3025
	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
	
	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.`)

	limit := 100
	sumOfSquares := math.Pow(float64(limit), 3)/3 + math.Pow(float64(limit), 2)/2 + float64(limit)/6
	sqaureOfSum := math.Pow((math.Pow(float64(limit), 2)/2 + float64(limit)/2), 2)

	fmt.Println(sumOfSquares)
	fmt.Println(sqaureOfSum)

	fmt.Println(sqaureOfSum - sumOfSquares)
}
