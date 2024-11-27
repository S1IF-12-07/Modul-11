package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a%10 == 0:
		perhitungan := a / 10
		fmt.Println("kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d", a, perhitungan)
		break
	case a%5 == 0 && a != 5:
		perhitungan := a * a
		fmt.Println("kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil perpangkatan antara %d ^ 2 = %d", a, perhitungan)
		break
	case a%2 == 0:
		bilanganberikutnya := a + 1
		perhitungan := a * bilanganberikutnya
		fmt.Println("kategori: Bilangan Genap")
		fmt.Printf("Hasil dengan bilangan berikutnya %d * %d = %d", a, bilanganberikutnya, perhitungan)
		break
	case a%2 !=0:
		bilanganberikutnya := a + 1
		perhitungan := a + bilanganberikutnya
		fmt.Println("kategori: Bilangan Ganjil")
		fmt.Printf("Hasil  penjumlahan dengan bilangan berikutnya %d + %d = %d", a, bilanganberikutnya, perhitungan)
		break
	}
}