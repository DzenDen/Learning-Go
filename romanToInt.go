package main

import "fmt"

func main () {
rome := map [rune] int {
	'I' : 1,
	'V' : 5,
	'X' : 10,
	'L' : 50,
	'C' : 100,
	'D' : 500,
	'M' : 1000,
}
s := "MMDCCLXXXIX"
arab := make([]int, 0)
sum := 0
		for _, c := range s {
			x := rome [c]
			arab = append(arab, x)
		} 
		fmt.Println (arab)

		max := arab [0]	
		for _, z := range arab {
			if z > max {
			sum = sum + (z - 2*max)
			max = z
			} else {
				sum = sum + z
				max = z
			}
		}
	fmt.Println (sum)
}