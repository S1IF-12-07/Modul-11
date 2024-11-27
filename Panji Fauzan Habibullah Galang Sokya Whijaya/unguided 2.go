package main

import (
	"fmt"
	"math"
)

func main() {
	var kendaraan string
	var durasi float64
	var TBiayaParkir int
	fmt.Scan(&kendaraan, &durasi)
	fmt.Println(kendaraan, durasi, "jam")
	if durasi < 1 {
		durasi = 1
	} else {
		durasi = math.Round(durasi)
	}
	switch kendaraan {
	case "Motor", "motor":
		TBiayaParkir = 2000 * int(durasi)
	case "Mobil", "mobil":
		TBiayaParkir = 5000 * int(durasi)
	case "Truk", "truk", "Truck", "truck":
		TBiayaParkir = 8000 * int(durasi)
	}
	fmt.Print("Rp ", TBiayaParkir)
}
