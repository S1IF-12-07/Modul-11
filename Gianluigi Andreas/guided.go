package main

import "fmt"

func main() {
	var u int
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&u)

	switch {
	case u < 0:
		fmt.Println("Umur tidak terdaftar dalam kategori")
		break
	case u < 13:
		fmt.Println("Anak-anak")
		break
	case u < 20:
		fmt.Println("Remaja")
		break
	case u < 60:
		fmt.Println("Dewasa")
	case u >= 60:
		fmt.Println("Lansia")

	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}