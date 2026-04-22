package main

import "fmt"

func tit() {
	n := 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == i || j == n-i+1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}