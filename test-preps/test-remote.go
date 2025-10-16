package main

import (
	"fmt"

	"github.com/dgempiuc/go-journey/remote_repo"
)

func main() {
	original := "Hello, World!"
	reversed := remote_repo.Reverse(original)
	fmt.Println(reversed)
}
