package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 20
	c, err := myf(a, b)
	fmt.Println(err)
	fmt.Println(c)

}

func myf(x int, y int) (int, error) {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if x < y {
		return -1, err
	}
	return 15, nil
}
