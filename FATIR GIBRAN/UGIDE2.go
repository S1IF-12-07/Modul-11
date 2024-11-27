package main

import "fmt"

func main() {
	var kendaraan string
	var durasi float64

	fmt.Print("Masukkan jenis kendaraan (motor, mobil, truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	var tarif float64
	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	totalBiaya := tarif * durasi
	fmt.Printf("Total biaya parkir: Rp %d\n", int(totalBiaya))
}
