// Exercise 5.15: Write variadic functions max and min, analogous to sum. What
// show these functions do when called with no arguments? Write variants that
// require at least one argument.

package main

import "fmt"

func Max(vals ...int) (val int, err error) {
	if len(vals) == 0 {
		err = fmt.Errorf("max error: no values given")
		return
	}
	val = vals[0]
	for _, v := range vals[1:] {
		if v > val {
			val = v
		}
	}
	return
}

func Min(vals ...int) (val int, err error) {
	if len(vals) == 0 {
		err = fmt.Errorf("min error: no values given")
		return
	}
	val = vals[0]
	for _, v := range vals[1:] {
		if v < val {
			val = v
		}
	}
	return
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	fmt.Println("max:", Must(Max(1)))
	fmt.Println("max:", Must(Max(1, 2, 3, 4, 5)))
	fmt.Println("min:", Must(Min(1)))
	fmt.Println("min:", Must(Min(1, 2, 3, 4, 5)))
	fmt.Println("max:", Must(Max()))
}
