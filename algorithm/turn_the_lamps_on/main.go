package main

import "fmt"

func main() {
	var multiTrip []int
	numberA := 0
	numberB := 100

	var resTrip []int

	var result int

	for numberA < numberB {

		tries := (numberA + 1)
		trip := numberB / tries
		if tries <= 3 {
			multiTrip = append(multiTrip, trip)

			tries++
		}

		for _, i := range multiTrip {

			resTrip = append(resTrip, i)
		}

		numberA++
	}

	for i := 0; i < numberB; i++ {
		result += resTrip[i]
	}

	fmt.Println(result)
}
