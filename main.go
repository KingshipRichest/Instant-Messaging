package main

import (
	"IM/router"
	"strconv"
)

func main() {
	e := router.Router()
	e.Run(":8080")
	//isPalindrome(121)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var i = strconv.Itoa(x)
	var left = 0
	var right = len(i)
	for left < right {
		if i[left] != i[right-1] {
			return false
		}
		right--
		left++
	}
	return true
}
