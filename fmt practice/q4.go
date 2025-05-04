package main

import("fmt")

func main(){

	var name string 

	fmt.Print("Your name:")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s \\n", name)
}