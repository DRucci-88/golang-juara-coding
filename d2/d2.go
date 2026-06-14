package main

import (
	"fmt"
	"strconv"
)

/*
Kontrol Alur Program: Mengatur Logika Percabangan dan Perulangan di Go
*/

func main() {
	//// Sesi 1: Logika Percabangan Tingkat Lanjut (If-Else) (45Menit)
	fmt.Println("============")
	fmt.Println("Sesi 1: Logika Percabangan Tingkat Lanjut (If-Else) (45Menit)")
	fmt.Println("============")

	/// 1. Sintaks Dasar If-Else
	fmt.Println("1. Sintaks Dasar If-Else")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	/// 2. If dengan Temporary Variable (Go-Specific Pattern)
	fmt.Println("2. If dengan Temporary Variable (Go-Specific Pattern)")

	// Format: if inisialisasi_variabel; kondisi_pengecekan { ::.
	if age, err := strconv.Atoi("25"); err != nil {
		fmt.Println("Gagal mengonversi umur:", err)
	} else {
		// Variabel 'age' hanya dapat diakses di dalam blok ini dan else-nya
		fmt.Printf("Konversi berhasil! Umur: %d tahun\n", age)
	}
	// Compile Error: fmt.Println(age) => age tidak dikenal di luar scope 'if'

	//// Sesi 2: Logika Percabangan dengan Switch (35 Menit)
	fmt.Println("============")
	fmt.Println("Sesi 2: Logika Percabangan dengan Switch (35 Menit)")
	fmt.Println("============")

	/// 1. Switch Standar & Multiple Expressions
	fmt.Println("1. Switch Standar & Multiple Expressions")

	day := "Sabtu"
	switch day {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari Kerja (Workday)")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur (Weekend)")
	default:
		fmt.Println("Hari tidak valid")
	}

	/// 2. Switch Tanpa Ekspresi (Pengganti If-Else Rantai)
	fmt.Println("2. Switch Tanpa Ekspresi (Pengganti If-Else Rantai)")

	backlogCount := 12
	switch {
	case backlogCount == 0:
		fmt.Println("Sistem aman, tidak ada antrean kerja.")
	case backlogCount > 0 && backlogCount <= 5:
		fmt.Println("Beban sistem: Ringan.")
	case backlogCount > 5 && backlogCount <= 10:
		fmt.Println("Beban sistem: Sedang. Perlu dipantau.")
	default:
		fmt.Println("⚠ BEBAN SISTEM TINGGI! Antrean menumpuk!")
	}

	/// 3. Kata Kunci  fallthrough
	fmt.Println("3. Kata Kunci  fallthrough")
	step := 1
	switch step {
	case 1:
		fmt.Println("Langkah 1: Inisialisasi Database")
		fallthrough
	case 2:
		fmt.Println("Langkah 2: Migrasi Skema Database")
		fallthrough
	case 3:
		fmt.Println("Langkah 3: Sinkronisasi Data Selesai")
	case 4:
		fmt.Println("Wait and See")
	}
	// Output jika step = 2:
	// Langkah 2: Migrasi Skema Database
	// Langkah 3: Sinkronisasi Data Selesai

	//// Sesi 3: Perulangan dengan Kata Kunci Tunggal:  for  (45Menit)
	fmt.Println("============")
	fmt.Println("Sesi 3: Perulangan dengan Kata Kunci Tunggal:  for  (45Menit)")
	fmt.Println("============")

	/// 1. Tiga Gaya Perluangan
	fmt.Println("1. Tiga Gaya Perluangan")

	// A. Gaya Standart
	for i := 0; i <= 5; i++ {
		fmt.Println("Iterasi ke ", i)
	}

	// B. Gaya While (Single Condition Loop)
	counter := 1
	for counter <= 5 {
		fmt.Println("Nilai counter ", counter)
		counter++
	}

	// C. Gaya Infinite Loop (Tanpa Kondisi)
	counter = 1
	for {
		fmt.Println("Loop berjalan ...", counter)
		if counter >= 5 {
			break
		}
		counter++
	}

	/// 2. Iterasi Kumpulan Data dengan  for range
	fmt.Println("2. Iterasi Kumpulan Data dengan  for range")
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

	/// 3. Pengendalian Loop:  break ,  continue , dan Label
	fmt.Println("3. Pengendalian Loop:  break ,  continue , dan Label")

	// Loop Bersarang & Label (Nested Loop Break)
OuterLoop: // Mendefinisikan label luar
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Println("Kombinasi kriteria ditemukan, menghentikan seluruh iterasi")
				break OuterLoop
			}
			fmt.Printf("Menyinkronkan Node %d - Partisi %d\n", i, j)
		}
	}
	fmt.Println("Selesai OuterLoop.")

}
