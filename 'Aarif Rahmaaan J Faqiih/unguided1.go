package main

import "fmt"
func main() {
	fmt.Println("pH 6.5 - 8.6       : air layak minum")
	fmt.Println("pH kurang dari 6.5 : air tidak layak minum")
	fmt.Println("pH lebih dari 8.6  : air tidak layak minum")

	var pH float32; fmt.Print("pH: "); fmt.Scan(&pH)
	var keterangan string

	switch {
	case pH < 0 || pH > 14:
		keterangan = "Nilai pH tidak valid. Nilai pH harus antara 0 dan 14"
	case pH >= 6.5 && pH <= 8.6:
		keterangan = "Air layak minum"
	default:
		keterangan = "Air tidak layak minum"
	}
	fmt.Print(keterangan)
}
