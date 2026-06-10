package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Print("Jakarta\n")

	if false {
		fmt.Println("Statement")
	}

	fmt.Println("Next statement")

	// kalkulasi stock

	stock := 50
	quantity := 55

	if quantity > stock {
		fmt.Println("Maksimal pembelian", stock, "item")
	} else {
		fmt.Println("Lanjut checkout")
	}

	// Format: if inisialisasi_variabel; kondisi_pengecekan { . }
	if age, err := strconv.Atoi("2S"); err != nil {
		fmt.Println("Gagal mengonversi umur:", err)
	} else {
		// Variabel 'age' hanya dapat diakses di dalam blok ini dan else-nya
		fmt.Printf("Konversi berhasil! Umur: %d tahun\n", age)
	}

	fmt.Println("Lanjutan error handling")

	// convert date to weekday
	today := time.Now()
	day := today.Weekday().String()
	fmt.Println(day)

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Hari Kerja (Workday)")
	case "Saturday", "Sunday":
		fmt.Println("Hari Libur (Weekend)")
	default:
		fmt.Println("Hari tidak valid")
	}

	step := 2
	switch step {
	case 1:
		fmt.Println("Langkah 1: Inisialisasi Database")
		fallthrough
	case 2:
		fmt.Println("Langkah 2: Migrasi Skema Database")
		fallthrough
	case 3:
		fmt.Println("Langkah 3: Sinkronisasi Data Selesai")
	}

	// Looping

	// FOR LOOP
	for i := 0; i < 5; i++ { // 4 < 5
		fmt.Println(i)
	}

	// 1*3*5
	// *****
	// *****
	// *****
	// *****
	sisi := 10
	for i := 1; i <= sisi; i++ { // 6 <= 5
		for j := 1; j <= sisi; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	pin := 654321
	try := 3
	var input int

	for try >= 1 {
		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&input)
		if input == pin {
			fmt.Println("PIN Benar")
			break
		} else {
			try--
			fmt.Println("PIN Salah, sisa percobaan: ", try)
		}
	}

	// Slice data string
	logLevels := []string{"INFO", "WARN", "ERROR", "FATAL"}
	// index dan value dikembalikan secara bersamaan
	for idx, val := range logLevels {
		fmt.Printf("Indeks: %d, Level Log: %s\n", idx, val)
	}
	// Gunakan blank identifier (_) jika tidak membutuhkan indeks
	for _, val := range logLevels {
		fmt.Println("Tingkat Bahaya:", val)
	}

	// continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println("Sum:", sum)

	// sum prices from slice
	prices := []int{100000, 200000, 300000, 400000, 500000}
	total := 0
	for _, price := range prices {
		total += price // total = 0 + 100000 = 100000
	}
	fmt.Println("Total Price:", total)

}
