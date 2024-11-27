package main

import "fmt"

func main() {
	var jm, price int
	var unit, kondisi string
	fmt.Scan(&unit, &jm)
	switch {
	case unit == "motor":
		price = 2000 * jm
	case unit == "mobil":
		price = 5000 * jm
	case unit == "truk":
		price = 8000 * jm
	default:
		kondisi = "input unit atau jm salah"
	}
	fmt.Println("Rp. ", price, " ", kondisi)
}
