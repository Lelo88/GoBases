package main

import "testing"

func Test_checkPalindromeTrue(t *testing.T) {
	inputString := "aabaa"
	want := true
	got := checkPalindrome(inputString)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_checkPalindromeFalse(t *testing.T) {
	inputString := "abac"
	want := false
	got := checkPalindrome(inputString)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_main(t *testing.T) { 
	main()
}