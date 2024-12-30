package main
import (
	"fmt"
	"time"
	"os"
)

// 0.000000000s elapsed 
func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.11fs elapsed", time.Since(start).Seconds())
}