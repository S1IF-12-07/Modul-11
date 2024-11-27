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
		keterangan = "air layak minum"
	case ph > 8.6 && ph <= 14:
		keterangan = "air tidak layak minum"
	default:
		keterangan = "nilai ph tidak valid"
	}
	fmt.Println(keterangan)
}
