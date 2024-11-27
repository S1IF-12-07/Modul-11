package main

import "fmt"

func main() {
	var usia int
	var kategori string
	fmt.Scan(&usia)

	switch {
	case usia < 12:
		kategori = "anak anak"
	case usia < 20:
		kategori = "remaja"
	case usia < 60:
		kategori = "dewasa"
	case usia > 61:
		kategori = "undifined"
	}
	fmt.Println(kategori)
}
