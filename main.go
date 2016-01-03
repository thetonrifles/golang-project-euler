package main

import (
  "fmt"
  "github.com/thetonrifles/golang-project-euler/euler"
)

func main() {
  fmt.Printf("largest prime for %v = %v\n", 600851475143, euler.Problem003(600851475143))
  fmt.Printf("%v\n", euler.Problem004())
  fmt.Printf("smallest number = %v\n", euler.Problem005())
}
