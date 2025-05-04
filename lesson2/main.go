package main

import (
	"fmt"
	"math"
)

func main() {
	// Используем константы из пакета math
	var b byte = math.MaxUint8       // 255
	var smallI int32 = math.MaxInt32 //  2147483647
	var bigI uint64 = math.MaxUint64 //  18446744073709551615

	// Переполнение приведёт к «wrap around»
	b++
	smallI++
	bigI++

	fmt.Printf("b:       %v\n", b)     // 0
	fmt.Printf("smallI: %v\n", smallI) // -2147483648
	fmt.Printf("bigI:   %v\n", bigI)   // 0
}
