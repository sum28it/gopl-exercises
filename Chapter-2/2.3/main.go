package main

import (
	"fmt"
	"time"
)

var pc [256]int

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + (i & 1)
	}
}

func popCount(num uint64) int {

	start := time.Now()
	var count int
	for i := 0; i < 8; i++ {
		count += pc[byte(num>>(8*i))]
	}

	fmt.Println(time.Since(start).Nanoseconds())

	return count
}

func main() {

	fmt.Println("Enter a number:")
	var num uint64
	fmt.Scan(&num)

	fmt.Printf("Number of set bits: %d\n", popCount(uint64(num)))
}
