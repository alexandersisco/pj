package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
	"maps"
)

func main() {
	var args struct {
		Json     string            `arg:"-j,--json"`
		Strings  map[string]string `arg:"-s,--string"`
		Numbers  map[string]int    `arg:"-n,--number"`
		Booleans map[string]bool   `arg:"-b,--bool"`
	}

	arg.MustParse(&args)

	jsonArg := args.Json
	if args.Json == "" {
		jsonArg = "{}"
	}
	dat := MergeJson(
		jsonArg,
		ToJson(args.Strings),
		ToJson(args.Numbers),
		ToJson(args.Booleans))

	jsonStr := ToJson(dat)
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

func MergeJson(jsonStrings ...string) map[string]any {
	merged := FromJson("{}")
	for _, jsonStr := range jsonStrings {
		d := FromJson(jsonStr)
		maps.Copy(merged, d)
	}

	return merged
}
