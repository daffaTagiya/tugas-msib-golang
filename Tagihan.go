package main

import "fmt"

func main() {
	// Data uji
	tagihan1 := 275.0
	tagihan2 := 40.0
	tagihan3 := 430.0

	// Menghitung tip untuk tagihan1
	tip1 := 0.15 * tagihan1
	if tagihan1 < 50 || tagihan1 > 300 {
		tip1 = 0.20 * tagihan1
	}
	total1 := tagihan1 + tip1

	// Menghitung tip untuk tagihan2
	tip2 := 0.15 * tagihan2
	if tagihan2 < 50 || tagihan2 > 300 {
		tip2 = 0.20 * tagihan2
	}
	total2 := tagihan2 + tip2

	// Menghitung tip untuk tagihan3
	tip3 := 0.15 * tagihan3
	if tagihan3 < 50 || tagihan3 > 300 {
		tip3 = 0.20 * tagihan3
	}
	total3 := tagihan3 + tip3

	// Menampilkan hasil di konsol
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan1, tip1, total1)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan2, tip2, total2)
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan3, tip3, total3)
}
