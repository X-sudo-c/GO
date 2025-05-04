/*
4. Reading Input
Task: Use fmt.Scanln() to read a user’s name from the terminal and print: "Hello, [name]!".
Example Input:

Bob  
*/

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){

	fmt.Println("You name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("Hello, %s", input) // okay so i managed to make it owrk wothout necesarily using any fmt.Scanln() but it still does its job
	





}
