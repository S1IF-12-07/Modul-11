package main

import "fmt"

func main() {
	var kode string
	fmt.Print("Masukan Kode :")
	fmt.Scan(&kode)
	switch {
	case kode == "G":
		fmt.Println("Semua usia")
	case kode == "PG":
		fmt.Println("Anak-anak di atas 7 tahun")
	case kode == "PG13":
		fmt.Println("Anak-anak di atas 13 tahun")
	case kode == "R":
		fmt.Println("Dewasa")
	default:
		fmt.Println("Kode tidak valid")
	}
}
