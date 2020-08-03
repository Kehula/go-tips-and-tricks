package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string)
	go func(ch chan string) {
		LOOP:
		for {
			select {
			case data, ok := <-ch:
				{
					if ok {
						fmt.Println(data)
					} else {
						fmt.Println("data closed")
						break LOOP
					}
				}
			}
		}
	}(ch)
	
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
	}
	close(ch)
	fmt.Println("going sleep for 5 seconds!")
	time.Sleep(5 * time.Second)
	fmt.Println("awake")
}
