package main
import "fmt"
func main(){
	var umur int
	fmt.Scan(&umur)
	switch {
	case umur<13:
		fmt.Println("anak-anak")
	case umur<20:
		fmt.Println("remaja")
	case umur<60:
		fmt.Println("dewasa")
	case umur>60:
		fmt.Println("lansia")
	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}