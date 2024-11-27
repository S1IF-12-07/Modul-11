package main
import "fmt"

func main()  {
	var hari int
	fmt.Scan(&hari)

	switch hari {
	case 1 :
		fmt.Println("hari senin")
	case 2 :
		fmt.Println("hari selasa")
	case 3 :
		fmt.Println("hari rabu")
	case 4 :
		fmt.Println("hari kamis")
	case 5 :
		fmt.Println("hari jum'at")
	case 6 :
		fmt.Println("hari sabtu")
	case 7 :
		fmt.Println("hari minggu")	
	default :
		fmt.Println("bukan hari jirr")		
	}
}