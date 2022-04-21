package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonnacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonnacci(n-1) + Fibonnacci(n-2)
}
