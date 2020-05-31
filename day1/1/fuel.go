package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println(calc(data()))
}

// data : Get the data from the file and parse it into []int
func data() []int {
	data, _ := ioutil.ReadFile("./FuelRequirement.txt")
	array := strings.SplitAfter(string(data), "\n")
	var val int
	var value []int
	for _, v := range array {
		fmt.Sscanf(v, "%d", &val)
		value = append(value, val)
	}
	return value
}

// calc : Calculate the fuel needed for first stage
func calc(data []int) int {
	var sum int
	for _, v := range data {
		v = v/3 - 2
		sum += v
	}
	return sum
}