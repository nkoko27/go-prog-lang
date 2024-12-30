package main

import (
	"fmt"
	"os"
	"time"
)

// 0.000528300s elapsed
func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.9fs elapsed\n", time.Since(start).Seconds())
}