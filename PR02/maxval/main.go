package main

import "fmt"

func main() {
	var NumArrays [5]int
	NumArrays[0] = 3
	NumArrays[1] = 7
	NumArrays[2] = 1
	NumArrays[3] = 8
	NumArrays[4] = 5

	maks := NumArrays[0]
	n := len(NumArrays)
	k := 2

	for k <= n {
		if NumArrays[k-1] > maks {
			maks = NumArrays[k-1]
		}
		k++
	}

	fmt.Println("Maximum value found:", maks)
}
