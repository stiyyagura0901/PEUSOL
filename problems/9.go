
package problems
import "fmt"

// Nine ...
func Nine() {
	fmt.Println(`A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

	a2 + b2 = c2
	For example, 32 + 42 = 9 + 16 = 25 = 52.
	
	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.`)

	sum := 1000

	for a := 1; a <= sum/2; a++ {
		for b := a + 1; b <= sum/2; b++ {
			c := sum - a - b
			if (a*a + b*b == c*c) {
				fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
				fmt.Println("PRODUCT of abc:", a*b*c)
				break
			}
		}
	}
}
