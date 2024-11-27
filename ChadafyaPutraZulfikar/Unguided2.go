package main

import "fmt"

func main() {
	var kendaraan string
	var waktu int
	fmt.Print("Masukkan tipe kendaraan dan durasi parkir dalam jam: ")
	fmt.Scan(&kendaraan, &waktu)

	switch kendaraan {
	case "motor":
		motor := 2000 * waktu
		fmt.Println("Tarif parkir : ", motor)
	case "mobil":
		mobil := 5000 * waktu
		fmt.Println("Tarif parkir : ", mobil)
	case "truk":
		truk := 8000 * waktu
		fmt.Println("Tarif parkir : ", truk)
	default:
		fmt.Println("Data tidak ada")
	}
}
