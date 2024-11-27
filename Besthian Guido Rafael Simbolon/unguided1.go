package main

import "fmt"

func main() {
	var ph float64
	fmt.Print("kadar pH Air : ")
	fmt.Scan(&ph)
	switch {
	case ph >= 6.5 && ph <= 8.6 :
		fmt.Println("Air layak diminum")
	case ph < 6.5 && ph > 0|| ph > 8.6 && ph <= 14:
		fmt.Println("air tidak layak diminum")
	case ph > 14 || ph < 0:
		fmt.Println("inputan harus antara 0-14")
	}
}