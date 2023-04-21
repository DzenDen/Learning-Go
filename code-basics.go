package main

import ("strings"
"fmt")

func main () {
	fmt.Println(ModifySpaces("Hello world", "vsda"))
}

func ModifySpaces(s, mode string) string {
	switch {
	case mode == "dash":
		return strings.ReplaceAll(s, " ", "-")
	case mode == "underscore":
		return strings.ReplaceAll(s, " ", "_")
	case mode == "" || mode == "unknown":
		return strings.ReplaceAll(s, " ", "*")
	default:
		return "idk"
}	
	}