package logger

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("[AUDIT SYSTEM] Menginisialisasi modul pencatatan transaksi terpusat...")
}

func LogTransaction(orderID string, amount float64) {
	fmt.Printf("[AUDIT SYSTEM - %s] Transaksi berhasil tercatat. Nilai Rp %.2f\n", time.Now().Format("2006-01-02 15:04:05"), amount)
}
