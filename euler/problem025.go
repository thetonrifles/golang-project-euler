package euler

import (
	"fmt"
	"math/big"
)

func Problem025() {
	// initialize two big ints with the first two numbers in the sequence
	a := big.NewInt(1)
	b := big.NewInt(1)
	// initialize limit as 10^999, the smallest integer with 100 digits
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)
	// loop while a is smaller than 1e100
	i := 1
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a
		a.Add(a, b)
		// swap a and b so that b is the next number in the sequence
		a, b = b, a
		// updating index
		i++
	}
	fmt.Printf("%v -> %v\n", i, a)
}
