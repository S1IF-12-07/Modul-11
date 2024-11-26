package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	fmt.Println("Jenis kendaraan & durasi parkir")
	fmt.Scan(&kendaraan, &durasi)

	motor := 2000
	mobil := 5000
	truck := 8000

	switch {
	case kendaraan == "motor" && durasi <= 1:
		fmt.Println(motor)
	case kendaraan == "motor" && durasi >= 2:
		fmt.Println(motor * durasi)
	case kendaraan == "mobil" && durasi <= 1:
		fmt.Println(mobil)
	case kendaraan == "mboil" && durasi >= 2:
		fmt.Println(mobil * durasi)
	case kendaraan == "truck" && durasi <= 1:
		fmt.Println(truck)
	case kendaraan == "truck" && durasi >= 2:
		fmt.Println(truck * durasi)

	}

}
