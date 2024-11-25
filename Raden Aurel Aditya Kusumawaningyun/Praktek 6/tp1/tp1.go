package main

import "fmt"

func main() {

	var usia int

	fmt.Println("Masukan Usia :")
	fmt.Scan(&usia)

	switch {

	case usia < 13:
		fmt.Println("Anak-Anak")
	case usia < 20:
		fmt.Println("Remaja")
	case usia < 60:
		fmt.Println("Dewasa")
	case usia >= 60:
		fmt.Println("Lansia")
	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}
