package main

import (
	"fmt"
	"time"
)

func main() {
  start := time.Now()
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
