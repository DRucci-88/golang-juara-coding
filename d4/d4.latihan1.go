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
Latihan 1: Sistem Upgrade Level Loyalty Member E-Commerce
*/
type LoyaltyTierEnum string

const (
	SILVER   LoyaltyTierEnum = "SILVER"   // Default
	GOLD     LoyaltyTierEnum = "GOLD"     // >= 10_000_000
	PLATINUM LoyaltyTierEnum = "PLATINUM" // >= 50_000_000
)

type ShippingAddress1 struct {
	Street     string
	City       string
	PostalCode string
}

func NewShippingAddress1(
	street string,
	city string,
	postalCode string,
) (*ShippingAddress1, error) {
	if street == "" || city == "" || postalCode == "" {
		return nil, errors.New("Shipping Address must not have a Zero-Value")
	}

	return &ShippingAddress1{
		Street:     street,
		City:       city,
		PostalCode: postalCode,
	}, nil
}

type CustomerProfile1 struct {
	ID          string
	Name        string
	LoyaltyTier LoyaltyTierEnum
	TotalSpent  float64
	Address     *ShippingAddress1
}

func NewCustomerProfile1(
	id string, name string, address *ShippingAddress1,
) (*CustomerProfile1, error) {
	if address == nil {
		return nil, errors.New("Shipping Address is nil")
	}

	return &CustomerProfile1{
		ID:          id,
		Name:        name,
		LoyaltyTier: SILVER,
		TotalSpent:  0.0,
		Address:     address,
	}, nil
}

func (profile *CustomerProfile1) RecordTransaction(amout float64) error {
	if profile == nil {
		return errors.New("Shipping is is nil")
	}

	profile.TotalSpent += amout
	if totalSpent := profile.TotalSpent; totalSpent >= 50_000_000 {
		profile.LoyaltyTier = PLATINUM
	} else if totalSpent >= 10_000_000 {
		profile.LoyaltyTier = GOLD
	}
	fmt.Println(profile)

	return nil
}

func (profile *CustomerProfile1) UpdateAddress(newAddress *ShippingAddress1) error {
	if profile == nil {
		return errors.New("Shipping is not defined")
	}

	profile.Address = newAddress
	return nil
}

func (profile CustomerProfile1) PrintProfile() {
	fmt.Println("Print Profile:")
	fmt.Printf("%+v\n", profile)
	fmt.Printf("%#v\n", profile)
	fmt.Printf("%+v\n\n", profile.Address)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("==================================================")
	fmt.Println(" Sistem Upgrade Level Loyalty Member E-Commerce")
	fmt.Println("==================================================")

	address, errAddress := NewShippingAddress1("Jalan Sudirman", "Jakarta Selatan", "14292")
	if errAddress != nil {
		fmt.Println(errAddress)
	}
	profile, errProfile := NewCustomerProfile1("Profile 1", "Le Rucco", address)
	if errAddress != nil {
		fmt.Println(errProfile)
	}
	var menu int

MainLoop:
	for {
		fmt.Println("\n0. Keluar aplikasi")
		fmt.Println("1. Record Transaction")
		fmt.Println("2. Update Address")
		fmt.Println("3. Print Customer Profile")
		menu = readInt1(reader, "Masukan menu: ")

		switch menu {
		case 0:
			break MainLoop
		case 1:
			menuRecordTransaction(reader, profile)
		case 2:
			menuUpdateAddress(reader, profile)
		case 3:
			menuPrintProfile(profile)
		default:
			fmt.Println("Salah pilih menu")
		}
	}
	fmt.Println("Selamat Tinggal")
}

func menuRecordTransaction(reader *bufio.Reader, profile *CustomerProfile1) {
	amount := readFloat1(reader, "Masukan Amount: ")
	err := profile.RecordTransaction(amount)
	if err != nil {
		fmt.Printf("%w", err)
	}
}

func menuUpdateAddress(reader *bufio.Reader, profile *CustomerProfile1) {
	street := readString1(reader, "Masukan Jalan: ")
	city := readString1(reader, "Masukan Kota: ")
	postalCode := readString1(reader, "Masukan Kode Pos: ")

	address, errAddress := NewShippingAddress1(street, city, postalCode)
	if errAddress != nil {
		fmt.Printf("%w", errAddress)
		return
	}
	fmt.Printf("New Address %+v", address)
	err := profile.UpdateAddress(address)
	if err != nil {
		fmt.Printf("%w", err)
	}
}

func menuPrintProfile(profile *CustomerProfile1) {
	profile.PrintProfile()
}

// Helper untuk membaca input bertipe string secara bersih
func readString1(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper untuk membaca input bertipe integer secara bersih
func readInt1(reader *bufio.Reader, prompt string) int {
	for {
		inputStr := readString1(reader, prompt)
		val, err := strconv.Atoi(inputStr)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka bulat.")
	}
}

// Helper untuk membaca input bertipe float64 secara bersih
func readFloat1(reader *bufio.Reader, prompt string) float64 {
	for {
		inputStr := readString1(reader, prompt)
		val, err := strconv.ParseFloat(inputStr, 64)
		if err == nil {
			return val
		}
		fmt.Println("Input tidak valid, harap masukkan angka desimal/nominal.")
	}
}

func clearConsole1() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
