package main

import (
	"algorithm/binarysearch"
	"fmt"
)

func main() {
	i := generateArray2()
	i2 := binarysearch.BinarySearch7(i, 8)
	fmt.Printf("i2: %v\n", i2)
}
func generateArray1() []int {
	return []int{2, 4, 5, 6, 8, 12, 14, 16, 18, 19, 21, 25, 23, 25, 27} //n = 14
}

func generateArray2() []int {
	return []int{2, 4, 5, 6, 8, 8, 8, 8, 12, 14, 16, 18, 19, 19, 19, 21, 25, 23, 25, 27}
}
