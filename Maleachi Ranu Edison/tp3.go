package main

import "fmt"

func main() {
	var kode string
	fmt.Scan(&kode)
	switch {
	case kode == "G" || kode == "g":
		fmt.Println("Semua Usia")
	case kode == "PG" || kode == "pg":
		fmt.Println("Anak-anak diatas 7 tahun")
	case kode == "PG-13" || kode == "pg-13":
		fmt.Println("Anak-anak diatas 13 tahun")
	case kode == "R" || kode == "r":
		fmt.Println("Dewasa")
	default:
		fmt.Println("Kode tidak valid")
	}
}
