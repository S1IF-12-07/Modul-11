//package main

//import "fmt"

//func main() {
//	var angka, hasil int
//	var kategori string
//	fmt.Scan(&angka)
//	switch {
//	case angka%2 != 0:
//		kategori = "bilangan ganjil"
//		hasil = angka + (angka + 1)
//	case angka%2 == 0:
//		kategori = "bilangan genap"
//		hasil = (angka * 1)
//	}
//	fmt.Println(kategori)
//	fmt.Printf("%d + %d = %d", angka, (angka + 1), hasil)
//}

// ///////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func main() {
	var angka, angka1, hasil int
	var kategori string //kondisi string
	fmt.Scan(&angka)

	genapganjil := angka % 2
	lima := angka % 5
	sepuluh := angka % 10

	switch {
	case lima == 0 && angka != 5 && sepuluh != 0:
		kategori = "bilangan kelipatan 5"
		angka1 = angka
		hasil = angka * angka
		//kondisi = "^2"
		fmt.Println(kategori)
		fmt.Printf("%d^2 = %d", angka, hasil)

	case sepuluh == 0:
		kategori = "bilangan kelipatan 10"
		angka1 = 10
		hasil = angka / 10
		//kondisi = "/ 10"
		fmt.Println(kategori)
		fmt.Printf("%d / %d = %d", angka, angka1, hasil)

	case genapganjil != 0:
		kategori = "bilangan ganjil"
		angka1 = angka + 1
		hasil = angka + angka1
		//kondisi = "+"//, angka1
		fmt.Println(kategori)
		fmt.Printf("%d + %d = %d", angka, angka1, hasil)

	case genapganjil == 0:
		kategori = "bilangan genap"
		angka1 = angka + 1
		hasil = angka * angka1
		//kondisi = "*"//, angka1
		fmt.Println(kategori)
		fmt.Printf("%d * %d = %d", angka, angka1, hasil)
	}
	//fmt.Println(kategori)
	//fmt.Printf("%d %s %d = %d", angka, kondisi, angka1, hasil)
}
