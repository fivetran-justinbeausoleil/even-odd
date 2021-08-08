package main

import (
	"fmt"
	"strconv"
)

type evenOddOutput []string

func main () {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultArray := evenOddOutput{}

	for _, number := range numbers {
		if number % 2 == 0 {
			resultArray = append(resultArray, strconv.Itoa(number)+" is even")
		} else {
			resultArray = append(resultArray, strconv.Itoa(number)+" is odd")
		}
	}
	for _, result := range resultArray {
		fmt.Println(result)
	}
}