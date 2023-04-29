package main

import "fmt"

func main () {
	data := map [rune] int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
	}
	x := "abdacef"
	sum := 0
	for _, c := range x {
			sum = sum + data [c]
	}
	fmt.Println (sum)
}