package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Konverter Suhu Multi-Skala
*/

func main() {
	fmt.Println("=================")
	fmt.Println("Konverter Suhu Multi-Skala")
	fmt.Println("=================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Suhu dalam celcius: ")
	celciusInput, _ := reader.ReadString('\n')
	celciusInput = strings.TrimSpace(celciusInput)
	celcius, errCelcius := strconv.ParseFloat(celciusInput, 64)
	if errCelcius != nil {
		fmt.Println("Terjadi kesalahan pada suhu")
	}

	fehrenheit := (celcius * 9.0 / 5.0) + 32.0
	remur := celcius * 4.0 / 5.0
	kelvin := celcius + 273.15

	fmt.Printf("celcius %.2f, fehrenheit %.2f, remur %.2f, kelvin %.2f", celcius, fehrenheit, remur, kelvin)
}
