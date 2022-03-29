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
	for jerseyNumber, selectedNumTimes := range jerseys {
		if jerseyNumber == 99 || jerseyNumber == 10 {
			percent = (float64(selectedNumTimes) / float64(plays)) * 100
			fmt.Println("#" + strconv.Itoa(jerseyNumber) + " gets " + fmt.Sprintf("%f", percent) + "%")
		}
	}
}
