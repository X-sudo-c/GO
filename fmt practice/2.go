/*
Task: Use fmt.Printf() to print:

The integer 42 with the format: "The answer is: %d".

The float 3.14159 as: "Pi is approximately: %.2f" (rounded to 2 decimal places).
Expected Output:

The answer is: 42  
Pi is approximately: 3.14  

*/

package main

import "fmt"


func main(){
	

	var num int = 42
	var pi float32 = 3.14159



	fmt.Printf("The answer is %d \n", num)
	fmt.Printf("Pi is aproximately: %.2f", pi)


}