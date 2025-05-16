//this is the else if statement in go similar to the elif statement in python
package main

import "fmt"

func main(){

	position := 2 

	if position == 1{
		fmt.Println("You won the gold!")
	}else if position == 2 {
		fmt.Println("You got the silver  medal")
	}else if position == 3{
		fmt.Println("Great Job on bronze.")
	}


}