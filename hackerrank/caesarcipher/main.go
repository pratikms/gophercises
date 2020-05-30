package main

import "fmt"

func rotate(s rune, delta int, key []rune) rune {
	idx := -1
	for i, r := range key {
		if r == s {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic("Rune not found in key")
	}
	for i := 0; i < delta; i++ {
		idx++
		if idx >= len(key) {
			idx = 0
		}
	}
	return key[idx]
}

func main() {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	newRune := rotate('z', 2, alphabet)
	fmt.Println(string(newRune))
}
