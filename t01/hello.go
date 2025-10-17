package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"journey/denizg/tutorial01/morepackage"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(morepackage.Reverse("Hello, world!"))
	fmt.Println(morepackage.ConvertUpper("Hello, world!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
