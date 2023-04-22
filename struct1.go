package main

import (
	"fmt"
)

//создаем кастомную структуру данных с 4мя полями
type User struct {
	name string
	sex string
	age Age
	weight int
}

//добавляем метод для нашей структуры
func (u User) PrintUserInfo (){
	fmt.Println(u.name, u.age, u.sex, u.weight, u.age.isAdult())
}

//добавляем кастомный тип данных
type Age int
//добавляем метод для проверки возраста
func (a Age) isAdult() bool {
	return a >= 18
} 

//добавляем конструктор для более удобного наполнения и инициализируем структуру
func NewUser (name, sex string, age, weight int) User {
	return User{
	name: name,
	sex: sex,
	age: Age(age),
	weight: weight,}
}
//инициализируем 3х юзеров и выводим данные на печать
func main (){
user1 := NewUser("Tom", "Male", 22, 69)
user2 := NewUser("Tim", "SheMale", 29, 58)
user3 := NewUser("Tina", "Female", 17, 45)

user1.PrintUserInfo()
user2.PrintUserInfo()
user3.PrintUserInfo()

}
