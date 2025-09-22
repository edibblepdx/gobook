// Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that
// represents a UTF-8 encoded string in place. Can you do it without allocating
// new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

// O(m*n) where m=runes n=bytes
func reverseShift(b []byte) []byte {
	bb := b[:]
	for len(bb) > 0 {
		rune, size := utf8.DecodeRune(bb)
		copy(bb[:], bb[size:])
		utf8.EncodeRune(bb[len(bb)-size:], rune)
		bb = bb[:len(bb)-size]
	}

	return b
}

// Swaps first two runes and returns their sizes
func swap(b []byte) (int, int) {
	rune1, size1 := utf8.DecodeRune(b)
	rune2, size2 := utf8.DecodeRune(b[size1:])

	utf8.EncodeRune(b, rune2)
	utf8.EncodeRune(b[size2:], rune1)

	return size1, size2
}

// O(m^2) where m=runes
func reverseBubble(b []byte) []byte {
	var fst, snd int
	count, b1 := utf8.RuneCount(b)-1, b[:]

	for range count {
		b2 := b1[:]

		for range count {
			fst, snd = swap(b2)
			b2 = b2[snd:]
		}

		b1 = b1[:len(b1)-fst]
		count--
	}

	return b
}

func main() {
	s1 := "◤₥⛚⋴┨⧕␺ⱍ⩽⍛⋘∀⥵⌐⧹⻡Ⅼ⏏⠍⃗ℴ⠲◇⸊⥱⩵⸮⑪⪖☴⣶⁗⪏⠭␆⩟⚹∝⬑╸ⴀ⾢ⴡⵍⷝ⻺ℤ⬗Ⰰ⦅⎂⦤⻜⏵⌻€⧬⡺⠛⎐⋚"
	s1 = string(reverseShift([]byte(s1)))
	fmt.Println(s1)

	fmt.Println()

	s2 := "◤₥⛚⋴┨⧕␺ⱍ⩽⍛⋘∀⥵⌐⧹⻡Ⅼ⏏⠍⃗ℴ⠲◇⸊⥱⩵⸮⑪⪖☴⣶⁗⪏⠭␆⩟⚹∝⬑╸ⴀ⾢ⴡⵍⷝ⻺ℤ⬗Ⰰ⦅⎂⦤⻜⏵⌻€⧬⡺⠛⎐⋚"
	s2 = string(reverseBubble([]byte(s2)))
	fmt.Println(s2)
}
