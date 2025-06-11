package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Json    string            `arg:"-j,--json"`
		Strings map[string]string `arg:"-s,--string"`
		Numbers map[string]int    `arg:"-n,--number"`
		Bool    map[string]bool   `arg:"-b,--bool"`
	}

	arg.MustParse(&args)

	if args.Json != "" {
		var dat map[string]any
		dat = FromJson(args.Json)

		jsonStr := ToJson(dat)
		fmt.Println(jsonStr)
		args.Json = ""
	}

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

func FromJson(jsonStr string) map[string]any {
	byt := []byte(jsonStr)
	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	return dat
}
