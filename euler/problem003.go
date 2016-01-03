package euler

import "math"

func Problem003(n int64) int64 {
  var outcome int64
  for b := sqrt(n); b >= 1; b-- {
		if n % b == 0 && isPrime(b) {
      outcome = b
			break
		}
	}
  return outcome
}

func isPrime(a int64) bool {
	if a <= 1 {
		return false
	}
	for b := sqrt(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a % b == 0 {
			return false
		}
	}

	return true
}

func sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}
