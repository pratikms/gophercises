package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	numberOfWords := 1
	min, max := 'A', 'Z'
	fmt.Println("min: ", min, "max: ", max)
	fmt.Printf("type(min): %T\n", min)
	fmt.Printf("type(max): %T\n", max)
	for _, ch := range input {
		if ch >= min && ch <= max {
			numberOfWords++
		}
	}

	fmt.Println("Number of words: ", numberOfWords)
}
