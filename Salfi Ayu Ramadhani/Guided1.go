package main

import "fmt"

func main() {
	var usia int
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	switch {
	case usia < 0:
		fmt.Print("Umur tidak masuk kategori")
	case usia < 13:
		fmt.Print("Anak-anak")
	case usia < 20:
		fmt.Print("Remaja")
	case usia < 60:
		fmt.Print("Dewasa")
	case usia >= 60:
		fmt.Print("Lansia")
	default:
		fmt.Print("Umur tidak masuk kategori")
	}
}
