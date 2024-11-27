package main

import "fmt"

func main() {
	var usia int
	fmt.Print("usia: ")
	fmt.Scan(&usia)

	switch {
	case usia == 0:
		fmt.Println("Usia tidak terdaftar dalam kategori")
		break
	case usia < 13:
		fmt.Println("Anak-anak")
		break
	case usia < 20:
		fmt.Println("Remaja")
		break
	case usia < 60:
		fmt.Println("Dewasa")
		break
	case usia >= 60:
		fmt.Println("Lansia")
	default:
		fmt.Println("Usia tidak terdaftar dalam kategori")
	}
}
