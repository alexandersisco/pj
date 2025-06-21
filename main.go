package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"maps"

	"github.com/alexflint/go-arg"
)

type args struct {
	Json     string            `arg:"-j,--json" help:"json string"`
	Strings  map[string]string `arg:"-s,--string" help:"key-value pairs where the value is a string"`
	Numbers  map[string]int    `arg:"-n,--number" help:"key-value pairs where the value is a number"`
	Booleans map[string]bool   `arg:"-b,--bool" help:"key-value pairs where the value is a bool."`
}

func (args) Description() string {
	return "\nThe easiest way to produce json strings on the command line\n"
}

func main() {

	var args args
	arg.MustParse(&args)

	// Handle stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	jsonIn := "{}"
	if (fi.Mode() & os.ModeNamedPipe) != 0 {
		// Stdin is connected to a pipe
		jsonIn, _ = ReadStdIn()
	}

	jsonArg := args.Json
	if args.Json == "" {
		jsonArg = "{}"
	}
	dat := MergeJson(
		jsonIn,
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
		log.Fatal(err)
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

func ReadStdIn() (string, error) {
	rdr := bufio.NewReader(os.Stdin)

	switch line, err := rdr.ReadString('\n'); err {
	case nil:
		return line, nil
	case io.EOF:
		return "", err
	default:
		fmt.Fprintln(os.Stderr, "error:", err)
		return "", err
	}
}
