package main

/*
Simulasi Tarif Parkir Progresi
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=================")
	fmt.Println("Simulasi Tarif Parkir Progresi")
	fmt.Println("=================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Jenis Kendaraan: ")
	jenisKendaraan, _ := reader.ReadString('\n')
	jenisKendaraan = strings.TrimSpace(jenisKendaraan)

	fmt.Print("Durasi parkir (jam): ")
	durasiInput, _ := reader.ReadString('\n')
	durasiInput = strings.TrimSpace(durasiInput)
	durasi, errDurasi := strconv.ParseInt(durasiInput, 10, 0)
	if errDurasi != nil || durasi < 0 {
		fmt.Println("Terjadi kesalahan pada tahun")
	}

	var tarif int64 = 0
	switch jenisKendaraan {
	case "mobil":
		tarif += 10_000
		tarif += durasi * 5_000
		tarif += durasi % 24 * 50_000
	case "motor":
		tarif += 5_000
		tarif += durasi * 2_000
		tarif += durasi % 24 * 50_000
	default:
		fmt.Println("Jenis kendaraan tidak terdaftar")
		return
	}
	fmt.Printf("Jenis Kendaraan: %s, Durasi: %s jam, tarif: %d", jenisKendaraan, durasi, tarif)
}
