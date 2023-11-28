package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// 向量类型有未导出字段，软件包无法访问这些字段。因此，我们编写了一对 BinaryMarshal/BinaryUnmarshal 方法，以便用 gob 包发送和接收该类型。这些接口定义在 "编码 "包中。我们也可以使用本地定义的 GobEncode/GobDecoder 接口。
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	//  简单编码：纯文本。
	var b bytes.Buffer
	fmt.Println(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary 会修改接收器，因此必须使用指针接收器。
func (v *Vector) UnmarshalBinary(data []byte) error {
	//  简单编码：纯文本。
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// 此示例传输的值实现了自定义编码和解码方法。
func main() {
	var network bytes.Buffer // 网络的替身。

	// 创建编码器并发送数值。
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// 创建解码器并接收数值。
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v) // {3 4 5}
}