package main

import (
	"fmt"
)

func pangkat(a float64, n int) float64 {
	hasil := 1.0
	for i := 1; i <= n; i++ {
		hasil *= a
	}
	return hasil
}

func main() {
	var a float64
	var n int

	fmt.Print("Masukkan Bilangan Utama: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan nilai pangkat: ")
	fmt.Scanln(&n)

	fmt.Printf("%.2f pangkat %d = %.2f\n", a, n, pangkat(a, n))
}
