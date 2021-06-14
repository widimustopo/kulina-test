package clothing_store

import (
	"fmt"
	"strconv"
)

func findPairing(items []int) int {
	var res int

	distincts := make(map[int]int)
	for _, num := range items {

		distincts[num] = distincts[num] + 1
		fmt.Println(distincts[num])
	}

	for _, distinct := range distincts {
		div := distinct / 2
		res += div
	}

	return res
}

func process() int {
	var pair int
	var pairs int
	var arrPair []int
	fmt.Println("================================")
	fmt.Printf("Input How many socks you have : ")
	fmt.Scanln(&pair)

	for i := 0; i < pair; i++ {
		fmt.Printf("Input your sock No.%v : ", i+1)
		fmt.Scanln(&pairs)
		arrPair = append(arrPair, pairs)
	}
	fmt.Println("================================")

	result := findPairing(arrPair)

	res := fmt.Sprintf("You have %v Pairs of socks", strconv.Itoa(result))
	fmt.Println(res)
	return process()
}

func main() {
	process()
}
