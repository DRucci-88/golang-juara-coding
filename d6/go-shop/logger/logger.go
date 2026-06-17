package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Inisialisasi variable logget dengan formatting warna
var (
	greenPrint  = color.New(color.FgGreen, color.Bold).PrintlnFunc()
	redPrint    = color.New(color.FgRed, color.Bold).PrintlnFunc()
	yellowPrint = color.New(color.FgYellow).PrintlnFunc()
)

// Fungsi init otomatis dieksekusi sekali saat paket di-load pertama kali
func init() {
	yellowPrint("[AUDIT SYSTEM] Menginisialisasi modul pencatatan transaksi terpusat...")
}

// LogSuccess mencetak audit sukses berwarna hijau (Exported)
func LogSuccess(orderID string, amount float64) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	greenPrint(fmt.Sprintf("[AUDIT SUCCESS - %s] Order %s berhasil diproses. Nilai: %.2f", timestamp, orderID, amount))
}

// LogFailure mencetak audit gagal berwarna merah (Exported)
func LogFailure(reason string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	redPrint(fmt.Sprintf("[AUDIT FAILURE - %s] Transaksi ditolak. Alasan: %s", timestamp, reason))
}
