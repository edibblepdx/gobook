// Exercise 5.16: Write a variadic version of strings.Join.

package main

import "fmt"

func Join(sep string, elems ...string) string {
	var res string
	n := len(elems)
	for _, s := range elems[:n-1] {
		res += s + sep
	}
	res += elems[n-1]
	return res
}

func main() {
	fmt.Println(Join("-", "fish"))
	fmt.Println(Join("-", "cats", "dogs", "rats", "dolphins"))
}
