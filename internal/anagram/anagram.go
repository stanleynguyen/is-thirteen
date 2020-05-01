// Package anagram Courtesy of https://siongui.github.io/2017/05/06/go-check-if-two-string-are-anagram/
package anagram

import (
	"sort"
)

type byRune []rune

func (r byRune) Len() int {
	return len(r)
}

func (r byRune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r byRune) Less(i, j int) bool {
	return r[i] < r[j]
}

func stringToRuneSlice(s string) byRune {
	var r byRune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// AreAnagram Check if 2 strings are anagram of each other
func AreAnagram(s1, s2 string) bool {
	var r1 byRune = stringToRuneSlice(s1)
	var r2 byRune = stringToRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}
