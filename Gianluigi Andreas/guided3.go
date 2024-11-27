package main

import "fmt"

func main() {
	var kode string
	fmt.Print("Masukkan kode film (G, PG, PG-13, R): ")
	fmt.Scan(&kode)

	switch kode {
	case "G":
		fmt.Println("G - Untuk semua umur")
	case "PG":
		fmt.Println("PG - Untuk anak-anak di atas 7 tahun dengan pengawasan orang tua")
	case "PG-13":
		fmt.Println("PG-13 - Untuk remaja di atas 13 tahun dengan pengawasan orang tua")
	case "R":
		fmt.Println("R - Untuk dewasa")
	default:
		fmt.Println("Kode tidak valid. Masukkan kode yang benar (G, PG, PG-13, R).")
	}
}
