package main

import "fmt"

func main() {
	var PH float64
	fmt.Scan(&PH)
	switch {
	case PH < 0 || PH > 14:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
	case PH >= 6.5 && PH <= 8.6:
		fmt.Println("Air layak minum")
	default:
		fmt.Println("Air Tidak Layak Minum")
	}
}
