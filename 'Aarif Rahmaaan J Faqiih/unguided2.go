package main

import (
	"fmt"
	"strings"
)

func main() {
	var kendaraan string; fmt.Print("kendaraan: "); fmt.Scan(&kendaraan)
	kendaraan = strings.ToLower(kendaraan)

	var jam int; fmt.Print("jam: "); fmt.Scan(&jam)

	switch kendaraan {
	case "motor": fmt.Printf("Rp %d\n", 2000*jam)
	case "mobil": fmt.Printf("Rp %d\n", 5000*jam)
	case "truk": fmt.Printf("Rp %d\n", 8000*jam)
	default:
		fmt.Println("maaf, lahan parkir untuk kendaraan tersebut tidak tersedia")
	}
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	MotorPerJam := 2000
// 	MobilPerJam := 5000
// 	TrukPerJam := 8000

// 	var kendaraan string; fmt.Print("kendaraan: "); fmt.Scan(&kendaraan)
// 	kendaraan = strings.ToLower(kendaraan)
// 	var jam int; fmt.Print("jam: "); fmt.Scan(&jam)

// 	TarifMotor := MotorPerJam * jam
// 	TarifMobil := MobilPerJam * jam
// 	TarifTruk := TrukPerJam * jam

// 	switch{
// 	case kendaraan == "motor": fmt.Print("Rp ", TarifMotor)
// 	case kendaraan == "mobil": fmt.Print("Rp ", TarifMobil)
// 	case kendaraan == "truk": fmt.Print("Rp ", TarifTruk)
// 	default: fmt.Print("maaf, lahan parkir untuk kendaraan tersebut tidak tersedia")
// 	}
// }

