/*
3. String Formatting with fmt.Sprintf()
Task: Store the formatted string "My name is %s and I am %d years old." in a variable using fmt.Sprintf(), where name = "Alice" and age = 25. Print the result.
Expected Output:

My name is Alice and I am 25 years old.  
*/





package main

import "fmt"


func main(){

	name := "Alice"
	age := 25

	info := fmt.Sprintf("My name is %s and I am %d years old." , name, age)

	fmt.Println(info)


}