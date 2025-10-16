package main

import (
	"fmt"

	"github.com/dgempiuc/go-journey/remoterepo"
)

func main() {
	original := "Hello, World!"
	reversed := remoterepo.Reverse(original)
	fmt.Println(reversed)
}
