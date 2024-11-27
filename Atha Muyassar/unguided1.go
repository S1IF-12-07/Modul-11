package main
import "fmt"

func main()  {
	var ph float64
	fmt.Scan(&ph)

	switch {
	case ph < 0 || ph > 14 :
		fmt.Println("input tidak valid, masukan ph 0-14 : ")
	case ph >= 6.5 && ph <= 8.6 :
		fmt.Println("air layak minum")
	default : 
		fmt.Println("air tidak layak minum")
	}
}