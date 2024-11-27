package main

import "fmt"

func main() {
	var kodeFilm string
	var usia int

	fmt.Print("Masukkan kode film (G, PG, PG-13, R): ")
	fmt.Scanln(&kodeFilm)

	switch kodeFilm {
	case "G":
		fmt.Println("Semua usia")
	case "PG":
		fmt.Print("Masukkan usia anak: ")
		fmt.Scanln(&usia)
		if usia > 7 {
			fmt.Println("Boleh menonton")
		} else {
			fmt.Println("Terlalu muda")
		}
	case "PG-13":
		fmt.Print("Masukkan usia remaja: ")
		fmt.Scanln(&usia)
		if usia > 13 {
			fmt.Println("Boleh menonton")
		} else {
			fmt.Println("Terlalu muda")
		}
	case "R":
		fmt.Println("Dewasa saja")
	default:
		fmt.Println("Kode film tidak valid")
	}
}
