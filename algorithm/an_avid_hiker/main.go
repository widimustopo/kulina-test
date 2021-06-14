package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func countingValleys(travel []string) int {

	var valley int
	var resAltitude int

	for _, i := range travel {
		if i == "U" {
			resAltitude += 1
		} else {
			resAltitude -= 1
			if resAltitude == -1 {
				valley += 1
			}
		}
	}

	return valley
}

func process() int {
	var n int
	var altitudes string
	var arrAltitude []string
	fmt.Println("================================")
	fmt.Printf("Input how many Uphill and Downhill you've traversed : ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Input traversed No.%v : ", i+1)
		fmt.Scanln(&altitudes)
		newString := strings.ToUpper(altitudes)
		regInput := regexp.MustCompile("^[DU]*$")
		if !regInput.MatchString(newString) {
			fmt.Println("input must be U for Up Hill AND D for Down Hill")
			return process()
		}

		arrAltitude = append(arrAltitude, newString)
	}
	fmt.Println("================================")

	result := countingValleys(arrAltitude)

	res := fmt.Sprintf("You had %v valley traversed", strconv.Itoa(result))
	fmt.Println(res)
	return process()
}

func main() {
	process()
}
