package main

import "fmt"

func tit() {
	n := 6

	for i := n/2; i <= n; i += 2 {
		for j := 1; j < n - i; j += 2 {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for j := 1; j < n - i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n; i >= 1; i-- {
		for j := i; j < n; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (i*2) - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}