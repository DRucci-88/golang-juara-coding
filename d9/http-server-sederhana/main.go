package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Definisikan Router / Multiplexer bawaan
	mux := http.NewServeMux()

	// 2. Hubungkan pola URL dengan fungsi Handler
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Menulis HTTP Response Body
		fmt.Fprintf(w, "Halo, selamat datang di Web Server Go!")
	})

	fmt.Println("Server berjalan di http://localhost:8080")

	// 3. Jalankan server pada port 8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
	
}
