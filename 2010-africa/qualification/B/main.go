package main

import (
	"./cdjm"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	inArr, _ := cdjm.Input("B-small-practice.in")

	//a == a[0:] // its a reminder how it includes number on the left

	size, _ := strconv.Atoi(inArr[0])
	inArr = inArr[1:]

	for i := 0; i < size; i++ {
		contains := strings.Contains(inArr[i], " ")
		fmt.Println(i+1, contains, "---------")
		if contains {
			// Reverse a
			a := strings.Split(inArr[i], " ")
			for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
				a[i], a[j] = a[j], a[i]
			}
			inArr[i] = strings.Join(a, " ")
		}

		fmt.Println(i+1, inArr[i])
	}

	cdjm.Output("output", inArr)

}
