package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	plays := 100000000
	jerseys := make(map[int]int)

	for i := 0; i < plays; i++ {
		num := rand.Intn(100)
		jerseys[num]++
	}

	percent := 0.0
	for k, v := range jerseys {
		if k == 99 || k == 10 {
			percent = (float64(v) / float64(plays)) * 100
			fmt.Println("#" + strconv.Itoa(k) + " gets " + fmt.Sprintf("%f", percent) + "%")
		}
	}
}
