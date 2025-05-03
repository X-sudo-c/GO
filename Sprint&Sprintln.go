package main

import "fmt"


func main(){
	quote := fmt.Sprintln("Look ma,", "no spaces!")

	fmt.Print((quote))//Prints look ma no spaces


	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer i s... %v!", correctAns)

	fmt.Print(answer)

	template := "I wish I had a %s."
	pet := "puppy"

	wish := fmt.Sprintf(template, pet)
	fmt.Println(wish)
}




