package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// 本示例使用解码器对不同的 JSON 值流进行解码。
func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
		// Ed: Knock knock.
		// Sam: Who's there?
		// Ed: Go fmt.
		// Sam: Go fmt who?
		// Ed: Go fmt yourself!
	}
}