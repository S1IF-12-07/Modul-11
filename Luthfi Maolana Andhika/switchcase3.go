package main

import "fmt"

func main() {
	var kode, ketentuan string
	fmt.Scan(&kode)
	switch {
	case kode == "G" || kode == "g":
		ketentuan = "untuk semua umur"
	case kode == "PG" || kode == "pg":
		ketentuan = "untuk anak-anak diatas 7 tahun"
	case kode == "PG-13" || kode == "pg-13":
		ketentuan = "untuk remaja diatas 13 tahun"
	case kode == "R" || kode == "r":
		ketentuan = "untuk dewasa"
	default:
		ketentuan = "kode tidak ditemukan"
	}
	fmt.Println(ketentuan)
}
