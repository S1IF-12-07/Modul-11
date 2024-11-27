package main

import "fmt"

func main() {
	var usia int

	fmt.Print("Masukkan usia: ")
	fmt.Scanln(&usia)

	switch {
	case usia < 13:
		fmt.Println("Kategori: Anak-anak")
	case usia < 20:
		fmt.Println("Kategori: Remaja")
	case usia < 60:
		fmt.Println("Kategori: Dewasa")
	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}
