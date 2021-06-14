package main

import (
	"fmt"
	"math"
)

type Lamps struct {
	number int
	turn   bool
}

func main() {
	var result int
	var lamps []Lamps
	number := 100

	//initial
	for i := 1; i <= number; i++ {
		lamps = append(lamps, Lamps{
			number: i,
			turn:   true,
		})
	}

	for i := 1; i < len(lamps); i++ {
		secondTrip := math.Mod(float64(lamps[i].number), 2)
		if secondTrip == 0 {
			if lamps[i].turn {
				lamps[i].turn = false
			} else {
				lamps[i].turn = true
			}
		}

		thirdTrip := math.Mod(float64(lamps[i].number), 3)
		if thirdTrip == 0 {
			if lamps[i].turn {
				lamps[i].turn = false
			} else {
				lamps[i].turn = true
			}
		}

		allTrip := math.Mod(float64(lamps[i].number), 1)
		if allTrip == 0 {
			if lamps[i].turn {
				lamps[i].turn = false
			} else {
				lamps[i].turn = true
			}
		}
	}

	for _, lamp := range lamps {
		if lamp.turn {
			result += 1
		}
	}

	fmt.Println(" Lamps currently turn on are : ", result, " lamps")
}
