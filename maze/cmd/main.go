package main

import (
	"fmt"
	"maze/algo"
)

func main() {
	maze := [][]int{
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 0},
	}

	solution := algo.SolveMaze(maze)
	if solution != nil {
		fmt.Println("Jalur solusi:")
		for _, step := range solution {
			fmt.Println(step)
		}
	} else {
		fmt.Println("Tidak ada solusi ditemukan.")
	}
}
