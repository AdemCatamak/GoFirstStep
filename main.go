package main

import (
	"fmt"
)

const pi = 3.1415926

func main() {
	var str1 string
	str1 = "str1"
	fmt.Println("Str1: " + str1)

	var str2 = "str2"
	fmt.Println("Str2: ", str2)

	fmt.Println("Const pi value: %s", pi)

	type User struct {
		Id        int
		FirstName string
		LastName  string
	}

	var defaultUser User
	fmt.Println("defaultUser -> ", defaultUser)

	var user1 = User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Println("user1 -> ", user1)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println("arr1 -> ", arr1)

	var slice1 = []int{1, 2, 3}
	fmt.Println("slice1 -> ", slice1)

	slice2 := append(slice1, 5, 6, 7)
	fmt.Println("slice2 -> ", slice2)
}
