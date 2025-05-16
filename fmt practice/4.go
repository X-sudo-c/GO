/*
4. Reading Input
Task: Use fmt.Scanln() to read a user’s name from the terminal and print: "Hello, [name]!".
Example Input:

Bob
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	fmt.Println("Input Your name")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()


	fmt.Printf("Hello, %s! ", input)

}