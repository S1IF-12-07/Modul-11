package main

import "fmt"

func main() {
	var r int
	var hari string
	fmt.Scan(&r)

	switch {
	case r == 1:
		hari = "senin"
	case r == 2:
		hari = "selasa"
	case r == 3:
		hari = "rabu"
	case r == 4:
		hari = "kamis"
	default:
		hari = "salah kabeh"
	}
	fmt.Println(hari)
}
