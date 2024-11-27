package main

import "fmt"

func main() {
	var ph float64
	fmt.Print("Masukkan kadar pH: ")
	fmt.Scanln(&ph)

	switch {
	case ph < 0 || ph > 14:
		fmt.Println("Input tidak valid")
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak diminum")
	default:
		fmt.Println("Air tidak layak diminum")
	}
}
