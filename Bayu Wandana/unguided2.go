package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	fmt.Println("Jenis Kendaraan & Durasi Parkir :")
	fmt.Scan(&kendaraan, &durasi)
	motor := 2000
	mobil := 5000
	truk := 8000
	switch {
	case kendaraan == "motor" && durasi <= 1:
		fmt.Println(motor)
	case kendaraan == "motor" && durasi >= 2:
		fmt.Println(motor * durasi)
	case kendaraan == "mobil" && durasi <= 1:
		fmt.Println(mobil)
	case kendaraan == "mobil" && durasi >= 2:
		fmt.Println(mobil * durasi)
	case kendaraan == "truk" && durasi <= 1:
		fmt.Println(truk)
	case kendaraan == "truk" && durasi >= 1:
		fmt.Println(truk * durasi)
	}
}
