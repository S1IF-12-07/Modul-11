package main

import "fmt"

func main() {
	var usia int
	fmt.Print("usia: ")
	fmt.Scan(&usia)

	switch {
	case usia < 0:
		fmt.Println("Umur tidak terdaftar dalam kategori")
		break
	case usia < 13:
		fmt.Println("Anak-anak")
		break
	case usia < 20:
		fmt.Println("Remaja")
		break
	case usia < 60:
		fmt.Println("Dewasa")
	case usia >= 60:
		fmt.Println("Lansia")
	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}