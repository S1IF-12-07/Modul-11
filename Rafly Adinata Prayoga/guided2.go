package main

import "fmt"

func main() {
	var d int
	var hasil string
	fmt.Println("Masukkan hari")
	fmt.Scan(&d)
	switch {
	case d == 1:
		hasil = "Senin"
	case d == 2:
		hasil = "Selasa"
	case d == 3:
		hasil = "Rabu"
	case d == 4:
		hasil = "Kamis"
	case d == 5:
		hasil = "Jumat"
	case d == 6:
		hasil = "Sabtu"
	case d == 7:
		hasil = "Minggu"
	default:
		fmt.Println("masukkan yang benar")
	}
	fmt.Print(hasil)

}
