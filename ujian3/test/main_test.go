package test

import (
	"fmt"
	"testing"
)

func TestApalah(m *testing.T) {
	input := "3a2b1c"

	var results string
	var num int
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			num = int(input[i])
		} else {
			for j := 0; j < num; j++ {
				results += string(input[i])
			}
		}
	}
	fmt.Println(results)
	// return result
}
