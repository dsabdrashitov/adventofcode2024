package main

import "fmt"

type A struct {
	a string
}

type B struct {
	a string
	b string
}

type C struct {
	A
	B
}

func printPromoted() {
	c := C{A{"afield"}, B{"aaa", "bfield"}}
	fmt.Println(c.A.a, c.B.a, c.b)
}
