package problems

import (
	"fmt"
	"strconv"
)

func reverseStr (s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i + 1, j -1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isReversible(i uint64, base int) bool {
	returnReversible := false
	b64 := uint64(base)
	for ; i > 0; i /= b64 {
		if int(i % b64) % 2 == 1 {
			returnReversible = true
		} else {
			return false
		}
	}
	return returnReversible
}

// OneHundredFortyFive ...
func OneHundredFortyFive() {
	fmt.Println(`Some positive integers n have the property that the sum [ n + reverse(n) ] 
	consists entirely of odd (decimal) digits. 
	For instance, 36 + 63 = 99 and 409 + 904 = 1313. 
	We will call such numbers reversible; so 36, 63, 409, and 904 are reversible. 
	Leading zeroes are not allowed in either n or reverse(n). 
	There are 120 reversible numbers below one-thousand. 
	How many reversible numbers are there below one-billion (10^9)?`)

	billion := 1000000000
	countOfReversibles := 0

	for i := 1; i < billion; i++ {
		itoa := strconv.Itoa(i)
		if i % 10 != 0 {
			reversed := reverseStr(itoa)
			reverse, _ := strconv.ParseInt(reversed, 10, 64)

			sum := int64(i) + reverse

			reversible := isReversible(uint64(sum), 10)
			if reversible {
				countOfReversibles++
			}
		}
	}

	fmt.Println(countOfReversibles)
}
