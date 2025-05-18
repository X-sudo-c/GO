package main


import (
	"fmt"
	"time"
)




func isItLateInNewYork(){
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5{
		lateMessage = "Goodness its late"

	}else if nyHour < 16 {
		lateMessage = "Its not late at all"
	}else if nyHour{
		lateMessage = "I guess its getting Late"
	} 
}