package main

import "fmt"

func main() {
	var phAir float64
	var n string
	fmt.Scan(&phAir)

	switch {
	case phAir < 6.5:
		n = "air tidak layak minum"
	case phAir >= 6.5 && phAir <= 8.6:
		n = "air layak minum"
	case phAir > 8.6 && phAir <= 14:
		n = "air tidak layak minum"
	default:
		n = "nilai phAir tidak valid"
	}
	fmt.Println(n)
}