package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	// Inisialisasi bMin dan bMax
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for _, berat := range arrBerat {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

func rerata(arrBerat arrBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlah)
}

func main() {
	var arrBerat arrBalita
	var bMin, bMax float64
	var jumlah int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(arrBerat, jumlah))
}
