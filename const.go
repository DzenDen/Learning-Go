package main

import "fmt"

//Реализуйте функцию ErrorMessageToCode(msg string) int, 
//которая возвращает числовой код для заданного значения.
//Мы для простоты ограничимся тремя ошибками. Учтите, что 
//если в функцию передать неизвестную ошибку, она должна 
//вернуть код ошибки для сообщения UNKNOWN.
func main (){
	fmt.Println(ErrorMessageToCode("vds"))
}
const (
	OkMsg = "OK"
	CanMsg = "CANCELLED"
)
const (
	OkCode = iota
	CanCode
	UnkCode
	)

func ErrorMessageToCode(msg string) int {
	switch msg {
		case OkMsg: return OkCode
		case CanMsg: return CanCode
		default: return UnkCode
	}

}

