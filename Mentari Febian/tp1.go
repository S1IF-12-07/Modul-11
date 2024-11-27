package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Termasuk dalam kategori: ")

	switch {
	case umur == 0:
		fmt.Println("Tidak Terdefinisi")
	case umur < 13:
		fmt.Println("Anak-Anak")
	case umur < 20:
		fmt.Println("Remaja")
	case umur < 60:
		fmt.Println("Dewasa")
	case umur >= 60:
		fmt.Println("Lansia")
	}
}
