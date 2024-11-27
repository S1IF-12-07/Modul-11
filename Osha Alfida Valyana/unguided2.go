package main

import "fmt"

func main() {
	var JenisKendaraan string
	var DurasiParkir int

	fmt.Print("Masukkan jenis kendaraan (Motor, Mobil, Truk): ")
	fmt.Scanln(&JenisKendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&DurasiParkir)

	if DurasiParkir < 1 {
		DurasiParkir = 1
	}

	var TarifPerJam int
	switch JenisKendaraan {
	case "Motor":
		TarifPerJam = 2000
	case "Mobil":
		TarifPerJam = 5000
	case "Truk":
		TarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	TotalBiaya := TarifPerJam * DurasiParkir
	fmt.Printf("Total biaya parkir: Rp %d\n", TotalBiaya)
}
