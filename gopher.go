package main

import (
	"fmt"
	"time"
)


func main(){

	x:= `      .-::::::-.
          .:-::::::::::::::-:.
           _:::    ::    :::_
           .:( ^   :: ^   ):.
            :::   (..)   :::.
            :::::::UU:::::::
           .::::::::::::::::.
           O::::::::::::::::O
           -::::::::::::::::-
            ::::::::::::::::
           .::::::::::::::.
              oO:::::::Oo `
          




	time.Sleep(5 * time.Second)

	fmt.Print("multilined", x)

	y := time.Now()

	format := y.Format("2006-01-02 15:04:05")


	fmt.Println("The new format:",format)






}