package main

import (
	"fmt"
	"github.com/aabeben/vini-vidi-vici-go/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

}
