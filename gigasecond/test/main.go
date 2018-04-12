package main

import (
	"fmt"
	"time"
)

func main() {
	timey := time.Now()

	newtime := timey.Add(time.Second * 500000000)

	fmt.Println(timey)
	fmt.Println(newtime)
}
