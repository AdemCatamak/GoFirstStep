package main

type Calculator struct{}

func (calc Calculator) Remove(a, b int) int {
	return a - b
}
