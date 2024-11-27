package main

import "fmt"

func main() {
	var usia int

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	switch {
	case usia < 0:
		fmt.Println("Usia tidak valid.")
	case usia < 13:
		fmt.Println("Kategori: Anak-anak")
	case usia < 20:
		fmt.Println("Kategori: Remaja")
	case usia < 60:
		fmt.Println("Kategori: Dewasa")
	default:
		fmt.Println("Kategori: Lansia")
	}
}
