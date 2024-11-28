package main

import (
	"fmt"
)

func main() {
	var x, y int
	var totalWeight [1000]float64
	var weights []float64

	// Input jumlah ikan dan kapasitas wadah
	fmt.Scan(&x, &y)

	// Input berat ikan
	weights = make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	// Menghitung total berat ikan di setiap wadah
	for i := 0; i < x; i++ {
		totalWeight[i/y] += weights[i]
	}

	// Output total berat ikan di setiap wadah
	for i := 0; i < (x+y-1)/y; i++ {
		fmt.Printf("%.2f ", totalWeight[i])
	}
	fmt.Println()

	// Menghitung dan output berat rata-rata ikan di setiap wadah
	for i := 0; i < (x+y-1)/y; i++ {
		count := 0
		for j := 0; j < y && (i*y+j) < x; j++ {
			count++
		}
		if count > 0 {
			average := totalWeight[i] / float64(count)
			fmt.Printf("%.2f ", average)
		}
	}
	fmt.Println()
}
