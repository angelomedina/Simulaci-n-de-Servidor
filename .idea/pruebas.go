package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	//operation that takes 20 milliseconds ...
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(((elapsed)+5)*time.Minute)
}