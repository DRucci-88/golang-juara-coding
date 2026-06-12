package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Kalkulator Indeks Masa Tubuh
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=================")
	fmt.Println("Kalkulator Indeks Masa Tubuh")
	fmt.Println("=================")

	fmt.Print("Masukan nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Masukan Berat Badan (kg): ")
	weightInput, _ := reader.ReadString('\n')
	weightInput = strings.TrimSpace(weightInput)
	weight, errWeight := strconv.ParseFloat(weightInput, 64)
	if errWeight != nil || weight < 0 {
		fmt.Println("Berat badan tidak mungkin dibawah 0")
		return
	}

	fmt.Print("Masukan Tinggi Badan (cm): ")
	heightInput, _ := reader.ReadString('\n')
	heightInput = strings.TrimSpace(heightInput)
	height, errHeight := strconv.ParseFloat(heightInput, 64)
	if errHeight != nil || height < 0 {
		fmt.Println("Tinggi badan tidak mungkin dibawah 0")
		return
	}
	heightMeter := height / 100

	bmi := weight / math.Pow(heightMeter, 2)

	var status string
	switch {
	case bmi < 18.5:
		status = "Underweight"

	case bmi < 25.0:
		status = "Normal"

	case bmi < 30.0:
		status = "Overweight"

	case bmi >= 30.0:
		status = "Obese"
	}

	fmt.Printf("Name: %s, BMI: %.2f, status: %s", name, bmi, status)
}
