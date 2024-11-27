package main

import "fmt"

func main() {
	var kendaraan string
	var waktu int

	fmt.Print("Masukkan tipe kendaraan dan lama parkir: ")
	fmt.Scan(&kendaraan, &waktu)

	switch kendaraan {
	case "motor":
		fmt.Printf("Tarif parkir: Rp %d\n", 2000*waktu)
	case "mobil":
		fmt.Printf("Tarif parkir: Rp %d\n", 5000*waktu)
	case "truk":
		fmt.Printf("Tarif parkir: Rp %d\n", 8000*waktu)
	default:
		fmt.Println("Input tidak valid. Tipe kendaraan tidak dikenali.")
	}
}
