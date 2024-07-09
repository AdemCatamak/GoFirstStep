package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415926

func main() {

	// Variable definition and assigment

	var str1 string
	str1 = "str1"
	fmt.Println("Str1: " + str1)

	var str2 = "str2"
	fmt.Println("Str2: ", str2)

	fmt.Println("Const pi value: %s", pi)

	var x, y, z int
	x = 1
	y = 2
	z = 3
	fmt.Println(x, y, z)

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

	// Function & Method
	var operand1 int
	operand1 = 1
	operand2 := 3

	addResult := Add(operand1, operand2)
	fmt.Println("addResult -> ", addResult)

	var calc = Calculator{}
	calcRemoveResult := calc.Remove(8, 5)
	fmt.Println("calcRemoveResult -> ", calcRemoveResult)

	Log("Log function execution")

	// Program Flow

	var index1 int
	for {
		fmt.Println("index1 -> ", index1)
		index1++
		if index1 > 3 {
			break
		}
	}

	var index2 int
	for index2 < 3 {
		fmt.Println("index2 -> ", index2)
		index2++
	}

	for index3 := 0; index3 < 3; index3++ {
		fmt.Println("index3 -> ", index3)
	}

	arr1 = [3]int{1, 2, 3}
	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	var option = 2
	switch option {
	case 1:
		fmt.Println("Case for 1")
	case 2:
		fmt.Println("Case for 2")
		fallthrough
	case 3:
		fmt.Println("Case for 3 fall from 2")
	default:
		fmt.Println("Case", option)
	}

	_, err := Divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}

	divideResult, err := Divide(6, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Divide result:", divideResult)
	}
}

func Add(x int, y int) int {
	return x + y
}

func Log(message string) {
	fmt.Println(message)
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}
