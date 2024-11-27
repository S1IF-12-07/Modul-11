package main

import "fmt"

func main() {
	var kode int
	var hari string
	fmt.Print("Masukkan kode hari:  ")
	fmt.Scan(&kode)
	switch {
	case kode == 1:
		hari = "Senin"
	case kode == 2:
		hari = "Selasa"
	case kode == 3:
		hari = "Rabu"
	case kode == 4:
		hari = "Kamis"
	case kode == 5:
		hari = "Jumat"
	case kode == 6:
		hari = "Sabtu"
	case kode == 7:
		hari = "Minggu"
	default:
		fmt.Println("Kode Tidak di Temukan")
	}
	fmt.Print(hari)

}
