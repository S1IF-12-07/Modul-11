package main
import "fmt"

func main() {
	var code string
	fmt.Scan(&code)
	switch code {
	case "G" :
		fmt.Println("semua umur")
	case "PG" : 
		fmt.Println("anak-anak diatas 7 tahun")
	case "PG-13" : 
		fmt.Println("remaja diatas 13 tahun")
	case "R" : 
	fmt.Println("dewasa")
	default :
	fmt.Println("kode tidak valid")
	}
}