package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// 	// fmt.Println(os.Args[1:])
// }

// exercise 1
// print index 0 of args as well

// func main() {
// 	fmt.Println(strings.Join(os.Args[0:], " "))
// }

// exercise 2
// print index and value on new line
// func main() {
// 	for index, value := range(os.Args[1:]) {
// 		fmt.Println(index, value)
// 	}
// }

// exercise 3
// try to measure time diff between versions 1, 2, 3
// 0.000000000s elapsed
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.9fs elapsed", time.Since(start).Seconds())
}