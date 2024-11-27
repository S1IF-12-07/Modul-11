package main

import "fmt"

func main() {
	var usia int
	fmt.Print("masukkan usia: ")
	fmt.Scan(&usia)

	switch {
	case usia <= 0:
		fmt.Print("Umur tidak terdaftar dalam kaategori")

	case usia < 13:
		fmt.Print("anak-anak")

	case usia < 20:
		fmt.Print("remaja")

	case usia < 60:
		fmt.Print("dewasa")

	case usia >= 60:
		fmt.Print("lansia")

	}
}
