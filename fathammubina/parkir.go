package main

import "fmt"

func main() {
	var kendaraan string
	var waktu int

	fmt.Print("tipe kendaraan dan waktu parkir : ")
	fmt.Scan(&kendaraan, &waktu)

	switch kendaraan {
	case "motor":
		x := 2000 * waktu
		fmt.Println("Tarif parkir : ", x)
	case "mobil":
		x := 5000 * waktu
		fmt.Println("Tarif parkir : ", x)
	case "truk":
		x := 8000 * waktu
		fmt.Println("Tarif parkir : ", x)
	default:
		fmt.Println("inputan tidak valid")
	}
}
