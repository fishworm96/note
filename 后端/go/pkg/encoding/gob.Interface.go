package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

// 本例展示了如何对接口值进行编码。与普通类型的关键区别在于注册实现接口的具体类型。
func main() {
	var network bytes.Buffer // 网络的替身。

	// 我们必须为编码器和解码器（通常位于编码器之外的另一台机器上）注册具体类型。在每一端，这都会告诉引擎发送的是实现接口的具体类型。
	gob.Register(Point{})

	// 创建编码器并发送一些数值。
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// 创建解码器并接收一些数值。
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
		// 5
		// 10
		// 15
	}
}

// interfaceEncode 将接口值编码到编码器中。
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// 除非具体类型已被 注册过。我们是在调用函数中注册的。

	// 通过接口指针，Encode 可以看到（并因此发送）接口类型的值。如果我们直接传递 p，它将看到的是具体类型。有关背景，请参阅博文 "反射法则"。
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// interfaceDecode 从数据流中解码下一个接口值并返回。
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	// 除非导线上的具体类型已被注册，否则解码将失败。我们在调用函数中注册了它。
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}