package main

import "fmt"

func main() {
	// Starcase1(10)
	// Starcase2(10)
	data := []int{1, 3, 5, 6, 2}
	Barchar(data)
}

func Starcase1(lenght int) {
	for i := 0; i < lenght; i++ {
		for s := lenght - i; s > 0; s-- {
			fmt.Print("  ")
		}
		for j := lenght - i; j < lenght; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func Starcase2(lenght int) {
	for i := 0; i < lenght; i++ {
		for s := lenght - i; s < lenght; s++ {
			fmt.Print("  ")
		}
		for j := lenght - i; j > 0; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func Barchar(data []int) {
	max := func() int {
		highest := 0
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				highest = data[i]
			}
		}
		return highest
	}
	for i := 0; i < max(); i++ {
		fmt.Print(max() - i)
		for j := 0; j < len(data); j++ {
			if data[j]+1 > max()-i {
				fmt.Print("-|")
			} else {
				fmt.Print("--")
			}
		}
		fmt.Println()
	}
	fmt.Print("0 ")
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}
}
