package main

import "fmt"


func main (){
	users := map[string]int{
		"Max": 21,
		"Pit": 42,
		"Ken": 19,
		"Tom": 4,
		"Jerry": 6,
	}
	users["Leen"] = 22
	delete(users, "Pit")

	for key, value := range users{
		if value < 20{
		fmt.Println("Mr",key, value,"YO")}
	}
}
