package main

import "fmt"
func main() {
	var usia int; fmt.Print("usia: "); fmt.Scan(&usia)
	
	switch {
	case usia<0: fmt.Print("belum lahir")
	case usia>60: fmt.Print("lansia")
	case usia>=20: fmt.Print("dewasa")
	case usia>=13: fmt.Print("remaja")
	case usia>=0: fmt.Print("anak-anak")
	default: fmt.Print("umur tidak terdaftar")
	}
}
