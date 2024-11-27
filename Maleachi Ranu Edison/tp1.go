package main

import "fmt"

func main() {
	var umur int
	fmt.Scan(&umur)
	switch {
	case umur < 13:
		fmt.Println("Anak-anak")
	case umur < 20:
		fmt.Print("Remaja")
	case umur < 60:
		fmt.Println("Dewasa")
	case umur > 60:
		fmt.Println("Lansia")
	default:
		fmt.Print("Umur tidak terdaftar dalam kategori")
	}
}
