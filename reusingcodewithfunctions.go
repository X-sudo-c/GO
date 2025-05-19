package main

import (
	"fmt"
	"math"
)

// Define specialComputation() here
func specialComputation(x float64) float64{
  return math.Log2(math.Sqrt(math.Tan(x)))
  
}


func dummy(x , y int32) int32 {
  fmt.Println("This is a dummy commit ")

  return x + y
}

func main() {
  
  var a, b, c, d float64
  a = .0214
  b = 1.02
  c = 0.312
  d = 4.001
  
  // Replace the following four lines with specialComputation()
  a = specialComputation(a)
  b = specialComputation(b)
  c = specialComputation(c)
  d = specialComputation(d)
  
  fmt.Println(a, b, c, d)


  fmt.Println(dummy(3,3))
}