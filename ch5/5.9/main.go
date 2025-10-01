// Exercise 5.9: Write a function expand(s string, f func(string) string) string
// that replaces each substring "$foo" within s by the text returned by f("foo").

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(expand("Hello, $World", func(in string) (out string) {
		if in == "World" {
			out = "世界"
		}
		return
	}))

	fmt.Println(expand("$a $b\t$c\n$d   $e", func(in string) (out string) {
		switch in {
		case "a":
			out = "1"
		case "b":
			out = "2"
		case "c":
			out = "3"
		case "d":
			out = "4"
		case "e":
			out = "5"
		}
		return
	}))
}

func expand(s string, f func(string) string) string {
	// I want to preserve white space, so I won't scan words
	// or split the string then join. And iterating over bytes
	// is a trap with unicode.

	var builder strings.Builder
	reader := strings.NewReader(s)

	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if ch == '$' {
			var rs []rune
			for {
				ch, _, err = reader.ReadRune()
				if err != nil || unicode.IsSpace(ch) {
					builder.WriteString(f(string(rs)))
					break // fall through (ch will be written after)
				}
				rs = append(rs, ch)
			}
		}
		builder.WriteRune(ch)
	}

	return builder.String()
}
