package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka hari (1-7): ")
	fmt.Scan(&angka)

	switch angka {
	case 1:
		fmt.Println("Hari: Minggu")
	case 2:
		fmt.Println("Hari: Senin")
	case 3:
		fmt.Println("Hari: Selasa")
	case 4:
		fmt.Println("Hari: Rabu")
	case 5:
		fmt.Println("Hari: Kamis")
	case 6:
		fmt.Println("Hari: Jumat")
	case 7:
		fmt.Println("Hari: Sabtu")
	default:
		fmt.Println("Angka yang dimasukkan tidak valid. Masukkan angka antara 1 hingga 7.")
	}
}
