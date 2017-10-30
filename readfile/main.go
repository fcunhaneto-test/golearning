package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a [][]float64
	var line []string
	fname := "teste.txt"
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(bs), "\n")
	for _, text := range lines {
		line = strings.Split(string(text), " ")
		a = append(a, lineToFloat(line))
	}

	fmt.Println(a)
}

func lineToFloat(line []string) []float64 {
	var s = []float64{}

	for _, l := range line {
		num, err := strconv.ParseFloat(l, 64)
		if err != nil {
			panic(err)
		}
		s = append(s, num)
	}

	return s
}
