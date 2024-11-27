package main

import "fmt"

func main() {
	var kode string

	fmt.Print("Masukkan kode film (G, PG, PG-13, R): ")
	fmt.Scan(&kode)

	switch kode {
	case "G":
		fmt.Println("Semua usia (G)")
	case "PG":
		fmt.Println("Anak-anak di atas 7 tahun (PG)")
	case "PG-13":
		fmt.Println("Remaja di atas 13 tahun (PG-13)")
	case "R":
		fmt.Println("Dewasa (R)")
	default:
		fmt.Println("Kode tidak valid.")
	}
}
