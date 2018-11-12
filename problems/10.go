package problems
import "fmt"

// Ten ...
func Ten() {
	fmt.Println(`The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

	Find the sum of all the primes below two million.`)

	val := 2000000
	sum := 0
	for i := 1; i <= val; i++ {
		if isPrime(i) {
			sum += i
		}
		fmt.Println(sum)
	}
	fmt.Println("SUM: ", sum)
}
