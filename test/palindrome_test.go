package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPalindrome(str string) bool {
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
	//isPalindromeRecursive(str, 0)
}

func isPalindromeRecursive(str string, current int) bool {
	if current == len(str)/2 {
		return true
	}
	if str[current] != str[len(str)-1-current] {
		return false
	}
	return isPalindromeRecursive(str, current+1)
}

func TestPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("natan"))
	assert.True(t, isPalindrome("aaadddaaa"))
	assert.False(t, isPalindrome("nathan"))
	assert.False(t, isPalindrome("aaxdddaaa"))

	assert.True(t, isPalindromeRecursive("natan", 0))
	assert.True(t, isPalindromeRecursive("aaadddaaa", 0))
	assert.False(t, isPalindromeRecursive("nathan", 0))
	assert.False(t, isPalindromeRecursive("aaxdddaaa", 0))
}
