package problems

import (
	"fmt"
	"sync"
)

const concurrentSummers = 10

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sumChannel(ch chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}

// One ...
func One() {
	fmt.Println("If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23. \nFind the sum of all the multiples of 3 or 5 below 1000.")
	limit := 10
	nums := []int{}
	input := make(chan int)
	output := make(chan int, concurrentSummers)

	for k := 1; k < limit; k++ {
		if (k%3 == 0) || (k%5 == 0) {
				nums = append(nums, k)
		}
	}

	go func() {
		for _, n := range nums {
			input <- n
		}
		close(input)
	}()

	var wg sync.WaitGroup
	wg.Add(concurrentSummers)

	for i := 0; i < concurrentSummers; i++ {
		go func () {
			output <- sumChannel(input)
			wg.Done()
		}()
	}

	wg.Wait()
	close(output)

	fmt.Println(sumChannel(output))
}
