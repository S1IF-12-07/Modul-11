package main

import "fmt"

func main() {
	var ph float64
	var keterangan string
	fmt.Scan(&ph)
	switch {
	case ph < 6.5:
		keterangan = "air tidak layak minum"
	case ph >= 6.5 && ph <= 8.6:
		keterangan = "air layak diminum"
	case ph > 8.6 && ph <= 14:
		keterangan = "air tidak layak minum"
	default: //punya syahdan...
		keterangan = "nilai ph tidak valid"
	}
	fmt.Print(keterangan)

}
