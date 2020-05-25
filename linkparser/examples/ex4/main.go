package main

import (
	"fmt"
	"strings"

	"github.com/pratikms/gophercises/linkparser"
)

var exampleHTML = `<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --> after comment</a>
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
