package main

import "fmt"

func main() {
	var kadarpH float64
	fmt.Scan(&kadarpH)
	switch {
	case kadarpH >= 6.5 && kadarpH <= 8.6:
		fmt.Print("Air layak minum")
	case kadarpH < 6.5 || kadarpH > 8.6 && kadarpH < 14:
		fmt.Print("Air tidak layak minum")
	case kadarpH > 14:
		fmt.Print("Nilai pH tidak valid, Nilai harus antara 0 dan 14 ")
	}
}
