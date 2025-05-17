//golang program to illustrate the usage of 
// Time.UnixNano() function


//including the main package 
package main   //this tells go compiler that this is a **standlone executable** program(not a library or a package).


//importing fmt and time 
import (
	"fmt"
	"time"
)

//calling main
func main(){

	// Defing t in UTC
	//for unix nano method

	t := time.Date(2019,12,15,23,59,12,5, time.UTC)



	unixnano := t.UnixNano()

	//Prints output
	fmt.Printf("%v\n", unixnano)
}