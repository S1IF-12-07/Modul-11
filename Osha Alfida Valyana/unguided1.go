package main

import "fmt"

func main() {
	var ph float64

	fmt.Print("Masukkan kadar pH : ")
	fmt.Scanln(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak diminum")
	case ph < 6.5 && ph > 0 || ph > 8.6 && ph <= 14:
		fmt.Println("air tidak layak diminum")
	case ph > 14 || ph < 0:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	}
}
