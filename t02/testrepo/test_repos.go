package main

import (
	"denizg/repos/localrepo"
	"fmt"
	"github.com/dgempiuc/go-journey/t02/remoterepo"
)

func main() {
	original := "Hello, World!"
	reversed := remoterepo.Reverse(original)
	truncated := localrepo.Truncate(original, 4, "...")
	fmt.Println(reversed)
	fmt.Println(truncated)
}
