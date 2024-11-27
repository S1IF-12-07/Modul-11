//Nama: Salsabilla Nurul Hassanah
//Kelas: 12-IF-07
//NIM: 103112430256

package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukkan umur : ")
	fmt.Scan(&umur)
	switch {
	case umur <= 0 :
	fmt.Println("Masukkan Umur >= 1")
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