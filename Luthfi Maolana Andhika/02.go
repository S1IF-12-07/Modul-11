package main

import "fmt"

func main() {
	var jam, tarif int
	var kendaraan, keadaan string
	fmt.Scan(&kendaraan, &jam)

	//jam := jam
	switch {
	case kendaraan == "motor": // && jam <= 24:
		tarif = 2000 * jam
	case kendaraan == "mobil": // && jam <= 24:
		tarif = 5000 * jam
	case kendaraan == "truk": // && jam <= 24:
		tarif = 8000 * jam
	default:
		keadaan = "input kendaraan atau jam salah"
	}
	fmt.Println("Rp. ", tarif, " ", keadaan)
}
