package main

import "fmt"

func main() {
	var rate, ketentuan string
	fmt.Scan(&rate)
	switch {
	case rate == "G" || rate == "g":
		ketentuan = "untuk semua umur"
	case rate == "PG" || rate == "pg":
		ketentuan = "untuk anak-anak diatas 7 tahun"
	case rate == "PG-13" || rate == "pg-13":
		ketentuan = "untuk remaja diatas 13 tahun"
	case rate == "R" || rate == "r":
		ketentuan = "untuk dewasa"
	default:
		ketentuan = "rate tidak ditemukan"
	}
	fmt.Println(ketentuan)
}
