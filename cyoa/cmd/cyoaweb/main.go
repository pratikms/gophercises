package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pratikms/gophercises/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with CYOA stories")
	flag.Parse()

	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
