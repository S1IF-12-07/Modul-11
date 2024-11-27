package main

import "fmt"

func main() {
	var kode string
	fmt.Print("Masukkan kode: ")
	fmt.Scan(&kode)

	switch kode {
	case "G":
		fmt.Print("Semua usia")
	case "PG":
		fmt.Print("Anak-anak di atas 7 tahun")
	case "PG-13":
		fmt.Print("Remaja di atas 13 tahun")
	case "R":
		fmt.Print("Dewasa")
	default:
		fmt.Print("Kode tidak masuk dalam daftar")
	}
}
