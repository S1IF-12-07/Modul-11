package main

import "fmt"

func getTarifPerJam(jenisKendaraan string) (int, error) {
    switch jenisKendaraan {
    case "motor":
        return 2000, nil
    case "mobil":
        return 5000, nil
    case "truk":
        return 8000, nil
    default:
        return 0, fmt.Errorf("jenis kendaraan tidak valid")
    }
}

func hitungBiayaParkir(tarifPerJam, durasiParkir int) int {
    if durasiParkir < 1 {
        durasiParkir = 1
    }
    return tarifPerJam * durasiParkir
}

func main() {
    var jenisKendaraan string
    var durasiParkir int
    fmt.Print("Masukkan jenis kendaraan (motor, mobil, truk): ")
    fmt.Scanln(&jenisKendaraan)
    fmt.Print("Masukkan durasi parkir (dalam jam): ")
    fmt.Scanln(&durasiParkir)
    tarifPerJam, err := getTarifPerJam(jenisKendaraan)
    if err != nil {
        fmt.Println(err)
        return
    }

    totalBiaya := hitungBiayaParkir(tarifPerJam, durasiParkir)
    fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}