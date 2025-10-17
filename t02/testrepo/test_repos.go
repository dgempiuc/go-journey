package main

import (
	"denizg/repos/localrepo"
	"fmt"

	"github.com/dgempiuc/go-journey/t02/remoterepo"
)

func main() {
	oriGinal := "Hello, World!"
	reversed := remoterepo.Reverse(oriGinal)
	truncated := localrepo.Truncate(oriGinal, 4, "...")
	fmt.Println(reversed)
	fmt.Println(truncated)
}
