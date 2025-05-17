package main

import "fmt"

func main() {
	var himpunan = [10]int{3, 23, 33, 64, 7, 26, 10, 88, 4, 5}

	n := len(himpunan)

	for i := 0; i < n; i++ {
		j := i
		temp := himpunan[i]
		for j > 0 && temp < himpunan[j-1] {
			himpunan[j] = himpunan[j-1]
			j--
		}
		himpunan[j] = temp
	}
	fmt.Print("MANDI MALAM DI WARUNG BATAK")
	fmt.Print("MMALAM KEOS INI")
	fmt.Println(himpunan)
}
