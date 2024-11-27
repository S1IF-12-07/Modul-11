package main
import "fmt"
func main() {
	var kode string
	fmt.Scanln(&kode)
	switch kode {
	case "G":
		fmt.Println("Kategori: Semua usia")
	case "PG":
		fmt.Println("Kategori: Anak-anak di atas 7 tahun")
	case "PG-13":
		fmt.Println("Kategori: Remaja di atas 13 tahun")
	case "R":
		fmt.Println("Kategori: Dewasa")
	default:
		fmt.Println("Kode tidak valid")
	}
}