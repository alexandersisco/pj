package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Strings map[string]string `arg:"-s,--string"`
	}

	arg.MustParse(&args)
	fmt.Println(args)
}
