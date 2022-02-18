package main

import "fmt"

func leftRotation(a []int, size int, rotation int) []int {
	var newArr []int
	for i := 0; i < rotation; i++ {
		newArr = a[1:size]
		newArr = append(newArr, a[0])
		a = newArr
	}
	return a
}

func main() {
	soal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotation := 3
	fmt.Println(leftRotation(soal, len(soal), rotation))
}
