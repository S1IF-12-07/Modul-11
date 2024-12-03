package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasiParkir int
	fmt.Print("Jenis kendaraan dan durasi parkir (dalam jam): ")
	fmt.Scan(&jenisKendaraan, &durasiParkir)

	switch jenisKendaraan {
	case "motor":
		tarifMotor := 2000 * durasiParkir
		fmt.Println("Biaya parkir : ", tarifMotor)
	case "mobil":
		tarifMobil := 5000 * durasiParkir
		fmt.Println("Biaya parkir : ", tarifMobil)
	case "truk":
		tarifTruk := 8000 * durasiParkir
		fmt.Println("Biaya parkir : ", tarifTruk)
	default:
		fmt.Println("Input tidak valid")
	}
}
