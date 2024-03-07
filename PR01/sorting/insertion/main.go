package main

import "fmt"

func main() {
	// Array yang akan diurutkan
	arr := []int{5, 3, 8, 4, 2}

	// Jumlah elemen dalam array
	n := len(arr)

	// Lakukan pass sebanyak n - 1 kali
	for i := n - 1; i >= 1; i-- {
		imaks := 0
		// Temukan indeks maksimum dari elemen array
		for j := 1; j <= i; j++ {
			if arr[j] > arr[imaks] {
				imaks = j
			}
		}

		// Tukar nilai antara arr[imaks] dan arr[i]
		temp := arr[i]
		arr[i] = arr[imaks]
		arr[imaks] = temp
	}

	// Cetak array yang sudah diurutkan
	fmt.Println("Array setelah diurutkan:")
	fmt.Println(arr)
}
