package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 本示例使用解码器对 JSON 对象流数组进行解码。
func main() {
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		var m Message
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	// json.Delim: [
	// 	Ed: Knock knock.
	// 	Sam: Who's there?
	// 	Ed: Go fmt.
	// 	Sam: Go fmt who?
	// 	Ed: Go fmt yourself!
	// 	json.Delim: ]
}