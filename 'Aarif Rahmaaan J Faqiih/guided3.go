package main

import "fmt"
func main() {
	var kodeFilm, usia string

	fmt.Println("G     : semua usia")
	fmt.Println("PG    : anak-anak di atas 7 tahun")
	fmt.Println("PG-13 : remaja di atas 13 tahun")
	fmt.Println("R     : dewasa")

	fmt.Print("kode film: "); fmt.Scan(&kodeFilm)

	switch kodeFilm {
	case "G":
		usia = "semua usia"
	case "PG":
		usia = "anak-anak di atas 7 tahun"
	case "PG-13":
		usia = "remaja di atas 13 tahun"
	case "R":
		usia = "dewasa"
	default:
		usia = "kode tidak valid"
	}
	fmt.Println(usia)
}
