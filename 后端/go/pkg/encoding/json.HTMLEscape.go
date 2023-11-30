package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout) // {"Name":"\u003cb\u003eHTML content\u003c/b\u003e"}
}
