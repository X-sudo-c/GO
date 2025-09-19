package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){

	fmt.Println("Enter a Series of Strings")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Input := scanner.Text()

	fmt.Printf("You typed: %q", Input)
}
//what does the scanner mean
