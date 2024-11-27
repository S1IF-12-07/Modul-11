package main

import "fmt"

func main() {
	var kode string
	fmt.Print("Masukkan kode film : ")
	fmt.Scan(&kode)
	switch kode {
	case "G":
		fmt.Println("untuk semua umur")
	case "PG":
		fmt.Println("untuk anak-anak diatas 7 tahun")
	case "PG-13":
		fmt.Println("untuk remaja diatas 13 tahun")
	case "R":
		fmt.Println("untuk dewasa")
	default:
		fmt.Println("kode tidak ditemukan")
	}
}
