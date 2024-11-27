package main

import "fmt"

func main() {
	var kendaraan string
	var parkir, harga, jam int
	fmt.Scan(&kendaraan)
	fmt.Scan(&jam)
	switch {
	case jam < 1:
		jam = 1
	case kendaraan == "Motor" || kendaraan == "motor":
		parkir = 2000
		harga = parkir * jam
		break
	case kendaraan == "Mobil" || kendaraan == "mobil":
		parkir = 5000
		harga = parkir * jam
		break
	case kendaraan == "Truk" || kendaraan == "truk":
		parkir = 8000
		harga = parkir * jam
	default:
		fmt.Println("Kendaraan tidak terdaftar")
	}
	fmt.Printf("Rp. %d", harga)
}
