package main

import "fmt"

func main() {
	var kode string
	fmt.Println("Masukkan kode")
	fmt.Scan(&kode)

	switch {
	case kode == "G":
		fmt.Println("Semua usia")
	case kode == "PG":
		fmt.Println("Anak-Anak diatas usia 7 tahun")
	case kode == "PG13":
		fmt.Println("Anak-Anak diatas 13 tahun")
	case kode == "R":
		fmt.Println("Dewasa")

	default:
		fmt.Println("Kode tidak valid")
	}
}
