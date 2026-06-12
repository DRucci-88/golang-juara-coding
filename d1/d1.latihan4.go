package main

/*
Generator & Validator Voucher Promo
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
	fmt.Println(" Generator & Validator Voucher Promo")
	fmt.Println("=================")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Tahun bergabung member: ")
	tahunInput, _ := reader.ReadString('\n')
	tahunInput = strings.TrimSpace(tahunInput)
	tahun, errTahun := strconv.ParseInt(tahunInput, 10, 0)
	if errTahun != nil || tahun < 0 {
		fmt.Println("Terjadi kesalahan pada tahun")
	}

	fmt.Print("Total Belanja: ")
	totalBelanjaInput, _ := reader.ReadString('\n')
	totalBelanjaInput = strings.TrimSpace(totalBelanjaInput)
	totalBelanja, errTotalBelanja := strconv.ParseFloat(totalBelanjaInput, 64)
	if errTotalBelanja != nil || totalBelanja < 0 {
		fmt.Println("Terjadi kesalahan pada tahun")
	}

	fmt.Print("Status Keanggotaan: ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	statusKelayakan := status == "premium" && tahun < 2024 || status == "regular" && totalBelanja > 1_000_000

	if statusKelayakan {
		totalBelanja = totalBelanja * (1 - 0.15)
	}

	fmt.Printf("Status Kelayakan %t, total belanja akhir: %.2f", statusKelayakan, totalBelanja)

}
