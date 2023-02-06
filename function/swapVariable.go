package function

import "fmt"

func SwapVariable(satu int, dua int) {
	satu += dua
	dua = satu - dua
	satu = satu - dua
	fmt.Printf("Variable A = %d\nVariable B = %d ", satu, dua)
}
