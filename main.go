package main

import (
	"fmt"
)

const LENGTH = 10

func main() {
	var mass [LENGTH]int
	for i := 0; i < LENGTH; i++ {
		fmt.Printf("Введите %v элемент массива: ", i)
		fmt.Scan(&mass[i])
	}
	fmt.Println(reverse(mass))
}

func reverse(mass [LENGTH]int) [LENGTH]int {
	length := len(mass)-1
	reverseMass := mass
	for i := 0; i < len(mass); i++ {
		reverseMass[length-i] = mass[i]
	}
	return reverseMass
}
