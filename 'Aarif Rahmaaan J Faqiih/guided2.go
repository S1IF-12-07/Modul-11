package main

import "fmt"
func main() {
	var urutan int
	var hari string

	fmt.Println("1 : senin")
	fmt.Println("2 : selasa")
	fmt.Println("3 : rabu")
	fmt.Println("4 : kamis")
	fmt.Println("5 : jumat")
	fmt.Println("6 : sabtu")
	fmt.Println("7 : minggu")

	fmt.Print("urutan hari: "); fmt.Scan(&urutan)

	switch urutan {
	case 1:
		hari = "senin"
	case 2:
		hari = "selasa"
	case 3:
		hari = "rabu"
	case 4:
		hari = "kamis"
	case 5:
		hari = "jumat"
	case 6:
		hari = "sabtu"
	case 7:
		hari = "minggu"
	default:
		fmt.Println("hanya ada 7 hari dalam seminggu!")
	}
	fmt.Println("nama hari: ", hari)
}
