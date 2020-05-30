package main

import (
	"fmt"
	"strings"
)

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("Rune not found in key")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
}

func main() {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	newRune := rotate('z', 2, alphabet)
	fmt.Println(string(newRune))
}
