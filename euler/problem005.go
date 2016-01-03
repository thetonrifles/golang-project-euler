package euler

import "fmt"

func Problem005() int {
  i := 1
  for {
    num := 20*i
    fmt.Printf("processing %v\n", num)
    divisible := divisibleByNumbers(num,
      1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
    if divisible {
      return num
    }
    i += 1
  }
}

func divisibleByNumbers(input int, numbers ...int) bool {
  divisible := true
  for _, number := range numbers {
    divisible = divisible && input%number==0
  }
  return divisible
}
