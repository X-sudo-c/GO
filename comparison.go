/*
1.
In main.go create a if statement that checks if lockCombo and robAttempt
 are the same. If the condition evaluates to true,
  print out "The vault is now opened.".

*/



package main


import "fmt"


func main(){

	lockCombo := "2-35-19"
	robAtttempt := "1-1-1"

	// Add your code below:

	if (lockCombo == robAtttempt){
		fmt.Println("The vault is now opened")
	}
}