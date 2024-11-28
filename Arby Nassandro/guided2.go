package main

import "fmt"

func main() {
	var hari int
	var hasil string
	fmt.Println("Masukkan hari")
	fmt.Scan(&hari)
	switch {
	case hari ==1 :
	hasil = "Senin"
	case hari ==2 : 
	hasil = "Selasa"
	case hari ==3 : 
	hasil = "Rabu"
	case hari ==4 : 
	hasil = "Kamis"
	case hari ==5 : 
	hasil = "Jumat"
	case hari ==6 : 
	hasil = "Sabtu"
	case hari ==7 : 
	hasil = "Minggu"
	default :
	fmt.Println("masukkan yang benar")
	}
	fmt.Print(hasil)

} 