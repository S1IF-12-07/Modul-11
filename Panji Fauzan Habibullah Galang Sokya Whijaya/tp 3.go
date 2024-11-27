package main

import "fmt"

func main() {
	var rating, kpi string
	fmt.Scan(&rating)
	switch rating {
	case "G", "g":
		kpi = "Semua usia"
	case "PG", "pg":
		kpi = "Anak-anak di atas 7 tahun"
	case "PG-13", "pg-13":
		kpi = "Remaja di atas 13 tahun"
	case "R", "r":
		kpi = "Dewasa"
	default:
		kpi = "Kode Tidak Valid"
	}
	fmt.Print(kpi)
}
