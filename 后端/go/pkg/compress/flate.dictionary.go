package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 预设字典可用于提高压缩率。使用字典的缺点是，压缩器和解压缩器必须事先商定使用什么字典。
func main() {
	// 字典是一串字节。在压缩某些输入数据时，
	// 压缩器会尝试用字典中找到的匹配字符串替换
	// 字典中的字符串。因此，字典只应包含预计会在实际数据流中找到的子串。
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	// 要压缩的数据应该（但不是必须）包含与字典中匹配的频繁 
	// 子串。
	const data = `<?xml version="1.0"?>
	<book>
		<meta name="title" content="The Go Programming Language" />
		<meta name="authors" content="Alan Donovan and Brian Kernighan" />
		<meta name="published" content="2015-10-26" />
		<meta name="isbn" content="978-0134190440" />
		<data>...</data>
	</book>
	`

	var b bytes.Buffer

	// 使用特制字典压缩数据。
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(zw, strings.NewReader(data)); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// 减压器必须使用与压缩机相同的字典。
	// 否则，输入可能会出现损坏。
	fmt.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	// 用 "#"替换字典中的所有字节，以显示
	// 证明了使用预设词典的大致效果。
	fmt.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}
/* Decompressed output using the dictionary:
<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>

Substrings matched by the dictionary are marked with #:
#####################
######
	############title###########The Go Programming Language"/#
	############authors###########Alan Donovan and Brian Kernighan"/#
	############published###########2015-10-26"/#
	############isbn###########978-0134190440"/#
	######...</#####
</##### */
}