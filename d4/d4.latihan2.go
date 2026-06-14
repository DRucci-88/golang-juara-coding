package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/*
Sistem Checkout & Potong Saldo e-Wallet (Mutasi Pointer)
*/

type EWallet2 struct {
	OwnerName string
	Balance   float64
}

type OrderTransaction2 struct {
	OrderID   string
	ItemTotal float64
	Status    string
}

func (order *OrderTransaction2) Pay(wallet *EWallet2) error {
	if wallet.Balance < order.ItemTotal {
		order.Status = "FAILED"
		return errors.New("pembayaran gagal: saldo e-wallet tidak mencukupi")
	}
	wallet.Balance -= order.ItemTotal
	order.Status = "PAID"
	return nil
}

func NewOrderTransaction(id string, total float64) (*OrderTransaction2, error) {
	if total <= 0 {
		return nil, errors.New("Total belanja tidak boleh kurang atau sama dengan nol")
	}
	return &OrderTransaction2{
		OrderID:   id,
		ItemTotal: total,
		Status:    "PENDING",
	}, nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" Sistem Checkout & Potong Saldo e-Wallet (Mutasi Pointer)")
	fmt.Println("==================================================")

	wallet := EWallet2{OwnerName: "LeRucco", Balance: 5_000_000.0}

	/// Transaksi A
	orderA, errOrderA := NewOrderTransaction("ORD-A", 1_500_000.0)
	if errOrderA != nil {
		fmt.Println(errOrderA)
	}
	errPayA := orderA.Pay(&wallet)
	if errPayA != nil {
		fmt.Println(errPayA)
	}
	fmt.Printf("OrderID [%s] Status [%s] Sisa Balance [%.2f]\n", orderA.OrderID, orderA.Status, wallet.Balance)

	/// Transaksi B
	orderB, errOrderB := NewOrderTransaction("ORD-B", 4_000_000.0)
	if errOrderB != nil {
		fmt.Println(errOrderB)
	}
	errPayB := orderB.Pay(&wallet)
	if errPayB != nil {
		fmt.Println(errPayB)
	}
	fmt.Printf("OrderID [%s] Status [%s] Sisa Balance [%.2f]\n", orderB.OrderID, orderB.Status, wallet.Balance)

}

// Helper untuk membaca input bertipe string secara bersih
func readString2(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt2(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString2(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat2(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString2(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}

func clearConsole2() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
