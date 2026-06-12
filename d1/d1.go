package main

import (
	"fmt"
	"strconv"
)

func main() {
	//// Sesi 2: Manajemen Variabel & Tipe Data Tingkat Lanjut (60 Menit)
	fmt.Println("Sesi 2: Manajemen Variabel & Tipe Data Tingkat Lanjut (60Menit")

	/// 1. Deklarasi Variabel & Type Inference
	fmt.Print("=========\n1. Deklarasi Variabel & Type Inference\n=========\n")
	// Eksplisit dengan inisialisasi nilai
	var email string = "deve"

	// Eksplisit tanpa nilai (Zero Value)
	var age int

	// Implisit (Type Inference) - Go Menebak tipe data berdasarkan nilainya
	var isManager = true

	// Menggunakan Short Declaration ( := )
	name := "John Doe"
	score := 95.5

	// Deklarasi Multi-Variable
	var x, y, z int = 1, 2, 3
	title, isPublished := "Golang Basics", false

	/// 3. Konstanta &  iota  (Enum ala Go)
	const user int = 0

	// Menentukan level otorisasi user menggunakan iota
	const (
		Guest int = iota // Bernilai 0
		User             // Bernilai 1
		Admin            // Bernilai 2
	)
	fmt.Println(Guest, User, Admin)

	/// 5. Konversi Tipe Data (Type Casting) & Library
	var a int = 10
	var b int32 = 20
	total := a + int(b)

	// Konversi Angka ke String (Integer to ASCII)
	strAge := strconv.Itoa(25)

	// Konversi String ke Angka (ASCII to Integer)
	age, err := strconv.Atoi("30")
	if err == nil {
		fmt.Println("Berhasil Konversi")
		fmt.Println(age)
	}

	/// 6. Variable Shadowing
	token := "token_utama"
	isValid := true
	if isValid {
		// Menggunakan := di sini menciptakan variabel 'token' BARU
		// yang hanya hidup di dalam blok 'if' ini.
		token := "token_rahasia"
		fmt.Println("Di dalam block:", token) // Output: token_rahasia
	}
	// Variabel token utama di luar tidak berubah
	fmt.Println("Di luar block:", token) // Output: token_utama

}
