package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

func largestPalindromeproduct() (int, int, int) {
	largestPalindrome := 0
	var multiplicand1, multiplicand2 int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j
			if product < largestPalindrome {
				break
			}
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
				multiplicand1 = i
				multiplicand2 = j
			}
		}
	}

	return largestPalindrome, multiplicand1, multiplicand2
}

func main() {
	result, multiplicand1, multiplicand2 := largestPalindromeproduct()
	fmt.Println("The largest palindrome product is :", result)
	fmt.Println("The multiplicand1 is : ", multiplicand1)
	fmt.Println("The multiplicand2 is : ", multiplicand2)
}