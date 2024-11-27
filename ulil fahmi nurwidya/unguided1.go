package main

import "fmt"
func cekKelayakanAir(ph float64) string {
    switch {
    case ph >= 6.5 && ph <= 8.6:
        return "Air layak diminum"
    case (ph < 6.5 && ph > 0) || (ph > 8.6 && ph <= 14):
        return "Air tidak layak diminum"
    case ph > 14 || ph < 0:
        return "Inputan tidak valid, harus antara 0-14"
    default:
        return "Inputan tidak valid"
    }
}

func main() {
    var ph float64
    fmt.Print("Masukkan kadar pH: ")
    fmt.Scanln(&ph)
    fmt.Println(cekKelayakanAir(ph))
}
