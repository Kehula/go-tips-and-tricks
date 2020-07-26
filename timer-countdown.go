package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}


//var tick <-chan time.Time
//tick = time.Tick(500 * time.Millisecond)
//select {
//case "something":
//	...
//case <-tick:
//	...
//case <-time.After(1 * time.Second):
//	...
//}
