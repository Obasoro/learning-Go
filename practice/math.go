package main

import {
	"fmt"
	"unicode/utf8"
}

func main() {
	s := "G"
	fmt.Println(s)

	fmt.Println("len gives the number of bytest\t", len(s))

	fmt.Println("RuneCountInString gives us the number of runses, eg character\t", utf8.RunesCountInString(s))

	fmt.Printf("%c \t %T\n", s[0], s[0])
	fmt.Printf("%c \t %T\n", s[1], s[1])

	for i, v := range s {
		fmt.Println(i, v)
	}
}