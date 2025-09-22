// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in
// a []string slice.

package main

import "fmt"

func elimdups(s []string) []string {
	if len(s) == 0 {
		return s
	}

	prev := s[0]
	pos := 1

	for _, v := range s[1:] {
		if v != prev {
			s[pos] = v
			pos++
		}
		prev = v
	}

	return s[:pos]
}

func main() {
	s1 := []string{
		"a", "a", "b", "c", "c", "c", "a", "a",
		"b", "a", "b", "b", "b", "c", "a", "a",
	}

	s1 = elimdups(s1)
	fmt.Println(s1)

	s2 := []string{"a", "b", "c", "d", "e", "f", "g"}

	s2 = elimdups(s2)
	fmt.Println(s2)

	s3 := []string{"a", "a", "b", "c", "c", "c", "a", "b"}

	s3 = elimdups(s3)
	fmt.Println(s3)
}
