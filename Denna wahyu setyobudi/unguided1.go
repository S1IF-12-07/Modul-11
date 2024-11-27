package main
import "fmt"
func main() {
	var ph float64
	fmt.Scanln(&ph)
	switch {
	case ph < 0 || ph > 14:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")
		break
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air Layak Minum")
		break
	case ph < 6.5 || ph > 8.6:
		fmt.Println("Air Tidak Layak Minum")
		break
	}
}
