package main

import "fmt"

func main() {
	var usia int
	fmt.Scan(&usia)
	switch {
	case usia >= 0 && usia < 13:
		fmt.Print("Anak-anak")
		break
	case usia >= 0 && usia < 20:
		fmt.Print("Remaja")
		break
	case usia >= 0 && usia < 60:
		fmt.Print("Dewasa")
		break
	case usia >= 60:
		fmt.Print("Lansia")
	default:
		fmt.Print("Umur tidak terdaftar dalam kategori")
	}
}
