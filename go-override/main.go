package main

import "fmt"

type test1 struct {
	message string
}

type test2 struct {
	*test1
}

func (t *test1) Get() {
	fmt.Println(t.message)
}

func (t *test1) GetGet() {
	fmt.Println(t.message)
	fmt.Println(t.message)
}

func (t *test2) Get() {
	fmt.Printf("test2だよ %s \n", t.message)
}

func main() {
	t := &test1{message: "test"}
	t.Get()

	t2 := &test2{test1: t}
	t2.Get()

	t2.GetGet()
}
