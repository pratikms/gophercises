package main

import (
	"fmt"
	"strings"

	"github.com/pratikms/gophercises/linkparser"
)

var exampleHTML = `<html>
<body>
  <h1>Hello world!</h1>
  <a href="/some-page">A link to some page</a>
  <a href="/some-other-page">A link to some other page</a>
</body>
</html>`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := linkparser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}