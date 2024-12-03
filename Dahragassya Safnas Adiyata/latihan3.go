package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	multipleOfTen := number%10 == 0 && number > 10
	square := number%5 == 0 && number > 5
	odd := number%2 != 0
	even := number%2 == 0

	switch {
	case multipleOfTen:
		result := number / 10
		fmt.Println("Kategori = Bilangan kelipatan 10")
		fmt.Printf("Hasil pembagian bilangan berikutnya %d / 10 = %d", number, result)
	case square:
		result := number * number
		fmt.Println("Kategori = Bilangan kelipatan 5")
		fmt.Printf("Hasil pangkat bilangan berikutnya %d^2 = %d", number, result)
	case odd:
		nextNumber := number + 1
		result := number + nextNumber
		fmt.Println("Kategori = Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d + %d = %d", number, nextNumber, result)
	case even:
		nextNumber := number + 1
		result := number * nextNumber
		fmt.Println("Kategori = Bilangan genap")
		fmt.Printf("Hasil perkalian bilangan berikutnya %d * %d = %d", number, nextNumber, result)
	}
}
