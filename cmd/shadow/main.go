package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(8 * time.Second)
	defer t.Stop()
	for range t.C {
		fmt.Println("this is a shadow log")
	}
}
