package euler

import "strconv"

func Problem004() string {
  var largest int
  for i:=999; i>=100; i-- {
    for j:=999; j>=100; j-- {
      current := i*j
      palindrome := isPalindrome(strconv.Itoa(current))
      if palindrome && current>largest {
        largest = current
      }
    }
  }
  return strconv.Itoa(largest)
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
