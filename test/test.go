package main

import "fmt"

type A interface {
	StartUnpack()
}

type B struct {
	b string
}

func (b *B) StartUnpack() {
	fmt.Println("B: Start" + b.b)
	// B 的 Start 方法逻辑
}

type C struct {
	// C 的字段
	c string
}

func (c *C) StartUnpack() {
	fmt.Println("C: Start" + c.c)
	c.ss()
	// C 的 Start 方法逻辑
}

func (c *C) ss() {
	fmt.Println("ss")
}

func main() {
	objMap := make(map[string]A)

	b := &B{"123"}
	c := &C{"456"}

	objMap["B"] = b
	objMap["C"] = c

	for _, obj := range objMap {
		obj.StartUnpack()
	}
}
