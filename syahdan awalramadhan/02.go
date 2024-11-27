package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, keterangan int
	fmt.Scan(&kendaraan, &durasi)
	switch kendaraan {
	case "motor":
		keterangan = durasi * 2000
	case "mobil":
		keterangan = durasi * 5000
	case "truk": //syahdan
		keterangan = durasi * 8000
	default:
		fmt.Print("kendaraan salah")
		return
	}
	fmt.Print("Rp ", keterangan)
}
