package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukkan umur : ")
	fmt.Scan(&umur)
	switch {
	case umur == 0 :
		fmt.Println("tidak terdefinisi dalam umur")
	case umur < 13 : 
	fmt.Println("anak-anak")
	case umur < 20 : 
	fmt.Println("remaja")
	case umur < 60 : 
	fmt.Println("dewasa")
	case umur > 60 :
	fmt.Println("lansia")
	}
} 