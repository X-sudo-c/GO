package main


import "fmt"


 func main(){
	rightTime := true
	rightPlace := true


	if rightTime && rightPlace{
		fmt.Println("We're outta here!")
	}else{
		fmt.Println("Be patient ..")
	}

	enoughRobbers := false
	enoughBAgs := true

	//Edit this conditoin for the SECOND  checkpoint
	if enoughRobbers && enoughBAgs{
		fmt.Println("Grab everything!")
	}else{
		fmt.Println("Grab whatever you can ")
	}


}


