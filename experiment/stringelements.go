package main

import "fmt"

func DoStringElements() {
	for i := '0'; i < '2'; i++ {
		fmt.Printf("(%T)%v\n", i, i)
	}
	s := "01"
	fmt.Printf("(%T)%v\n", s[0], s[0])
	fmt.Printf("(%T)%v\n", s[1], s[1])
	for _, c := range s {
		fmt.Printf("(%T)%v\n", c, c)
	}
	v0 := s[1] - 1
	fmt.Printf("(%T)%v\n", v0, v0)
	v1 := '0' + 1
	fmt.Printf("(%T)%v\n", v1, v1)
}
