package main

import "fmt"

func main() {
	var angka int

	fmt.Println("Masukan angka : ")
	fmt.Scan(&angka)

	switch angka {
	case 1:
		fmt.Println("senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")

	}
}