// Exercise 3.12: Anagrams.

package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1, r2 := []rune(s1), []rune(s2)
	m := make(map[rune]int)
	for i := range r1 {
		m[r1[i]]++
		m[r2[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("ちいかわ", "いわかち"))
	// --> true
	fmt.Println(isAnagram("はちわれ", "うさぎ"))
	// --> false
	fmt.Println(isAnagram("😭🥀🥀💔", "🥀😭💔🥀"))
	// --> true
	fmt.Println(isAnagram("wawawawa", "awawawaw"))
	// --> true
}
