package main

/*
Konverter Durasi Detik ke Jam-Menit-Detik
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=================")
	fmt.Println("Konverter Durasi Detik ke Jam-Menit-Detik")
	fmt.Println("=================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan detik: ")
	secondInput, _ := reader.ReadString('\n')
	secondInput = strings.TrimSpace(secondInput)
	totalSecond, errSecond := strconv.ParseInt(secondInput, 10, 0)
	if errSecond != nil || totalSecond < 0 {
		fmt.Println("Salah input detik")
		return
	}

	hour := int64(math.Floor(float64(totalSecond) / 3600.0))
	minute := int64(math.Floor(float64(totalSecond%3600) / 60.0))
	second := totalSecond - (hour * 3600) - (minute * 60)
	fmt.Println(hour, minute, second)
	fmt.Printf("%d detik setara dengan: %d jam, %d menit, %d detik", totalSecond, hour, minute, second)
}
