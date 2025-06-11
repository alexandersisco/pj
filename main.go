package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Strings map[string]string `arg:"-s,--string"`
	}

	arg.MustParse(&args)

	jsonStr := ToJson(args)
	fmt.Println(jsonStr)
}

func ToJson(data any) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}
