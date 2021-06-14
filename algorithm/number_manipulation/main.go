package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(items []string) []string {
	var result []string
	for i := 0; i < len(items); i++ {

		zeroCount := len(items) - i - 1
		multipler := 0
		strItems := items[i]
		if items[i] != "0" {
			for multipler < zeroCount {
				strItems = strItems + "0"
				multipler++
			}
			result = append(result, strItems)
		}
	}

	return result
}

func process() int {
	var number int
	var container []string

	fmt.Println("================================")
	fmt.Printf("Input numbers: ")
	fmt.Scanln(&number)
	fmt.Println("================================")
	fmt.Println("Result for ", number)
	fmt.Println("================================")

	numberStr := strconv.Itoa(number)

	numberArray := strings.Split(numberStr, "")

	for _, i := range numberArray {
		container = append(container, i)
	}

	resNum := convert(container)

	for i := 0; i < len(resNum); i++ {
		fmt.Println(resNum[i])
	}

	return process()
}

func main() {
	process()
}
