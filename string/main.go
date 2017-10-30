package main

import (
	"fmt"
)

func main() {
	s := "Olá coração"
	sr := `Olá coração hello mais uma mistura\n \u98`
	fmt.Println(len(s))
	fmt.Println([]byte(s))
	fmt.Println(s[2:4])
	fmt.Println(s[9:11])
	fmt.Println(sr)
}
