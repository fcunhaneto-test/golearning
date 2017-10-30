package main

import (
	"fmt"
	"math"
)

func main() {
	m := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	if age, ok := m["bob"]; ok {
		fmt.Println("idade:", age)

	} else {
		fmt.Println("bob não esta em map")
	}

}

func maxis(a [][]float64) {
	var i int
	var mapxl map[int]int
	for l := 0; l < (len(a) - 1); l++ {
		var max float64
		max = math.Abs(a[l][l])
		for i = (l + 1); i < len(a); i++ {
			if max < math.Abs(a[i][l]) {
				max = math.Abs(a[i][l])
				fmt.Printf("line %d: %.1f, %d %d\n", l, max, l, i)
				//if arrmax[l] = i
			}
			if _, ok := mapxl[l]; !ok {
				mapxl[l] = i
			}
		}
		fmt.Println("resultado:", max, l, i)
	}
}
