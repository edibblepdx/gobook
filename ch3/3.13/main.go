// Exercise 3.13: Write the constants for KB, MB, up through YB compactly

package main

// You can't use iota here.

import (
	"fmt"
)

const (
	KB = 1000      // Kilobyte
	MB = 1000 * KB // Megabyte
	GB = 1000 * MB // Gigabyte
	TB = 1000 * GB // Terabyte
	PB = 1000 * TB // Petabyte
	EB = 1000 * PB // Exabyte
	ZB = 1000 * EB // Zettabyte
	YB = 1000 * ZB // Yottabyte
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(float64(ZB))
	fmt.Println(float64(YB))
}
