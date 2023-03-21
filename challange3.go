package main

import "fmt"

func main() {
	text := "selamat malam"
	counter := make(map[string]int)
	for _, c := range text {
		fmt.Printf("%v\n", string(c))
		counter[string(c)]++
	}
	fmt.Println(counter)
}