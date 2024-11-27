//Nama: Salsabilla Nurul Hassanah
//Kelas: 12-IF-07
//NIM: 103112430256

package main

import "fmt"

func main() {
	var hari int
	fmt.Print("Masukkan angka hari: ")
	fmt.Scan(&hari)

	switch {
	case hari == 1:
		fmt.Print("Senin")
	case hari == 2:
		fmt.Print("Selasa")
	case hari == 3:
		fmt.Print("Rabu")
	case hari == 4:
		fmt.Print("Kamis")
	case hari == 5:
		fmt.Print("Jumat")
	case hari == 6:
		fmt.Print("Sabtu")
	case hari == 7:
		fmt.Print("Minggu")
	default:
		fmt.Print("Tidak terdaftar")
	}
}