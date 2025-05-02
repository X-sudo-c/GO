package main 

import "fmt"


func main(){
	specialNum := 42

	var y int = 23
	fmt.Println("The value of y is:", y)

	randomtext := "adklmalda"
	fmt.Printf("This values type is %T \n", specialNum)
	//Prints: This values type is int

	//the %T verb stands for the data type

	fmt.Println("This is a random string called %v", randomtext)


	quote := "To do or not to do"

	fmt.Printf("This value type is a %T", quote)

	//Prints this value type is a string 

	fmt.Println("\n")

	votingAge := 18

	fmt.Printf("You must be %d years old to vote.", votingAge)

	//Prints you must be 18 years to vote but 
	//Then syntax looks like c bcus in C you use percentage signs  and their data type to be able to to print them out %d stands for d


	gpa := 3.8

	fmt.Println("You are averaging f%.", gpa)
	//This prints you are averaging : 3.8 bcus %f stands for floats 


	yearsOfExp := 3
	reqYears := 15


	fmt.Printf("This company said they require %d years but i have about %d years in the bag", reqYears, yearsOfExp)
	//I just realised you can only used PrintF for stuff like this using Println wouldnt output the results of the %d formatter it would litterally print "%d"


	stockPrice := 3.50

	fmt.Printf("Each share odf Gopher is $%.2f! \n", stockPrice)








}