package main

import (
	"denizg/repo/localrepo"
	"fmt"
	"github.com/dgempiuc/go-journey/remoterepo"
)

func main() {
	original := "Hello, World!"
	reversed := remoterepo.Reverse(original)
	truancated := localrepo.Truncate(reversed, 4, "aaa")
	fmt.Println(reversed)
	fmt.Println(truancated)
}
