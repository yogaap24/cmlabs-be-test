package cmlabs_be_test

import (
	"regexp"
	"strings"
	"testing"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile("[^a-z0-9]").ReplaceAllString(s, "")
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man a plan a canal Panama", true},
		{"Was it a car or a cat I saw?", true},
		{"No x in Nixon", true},
		{"SAIPPUAKIVIKAUPPIAS", true},
		{"Aibohphobia", true},
		{"Anna", true},
		{"Civic", true},
	}

	for _, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.expected {
			t.Errorf("isPalindrome(%q) == %v, expected %v", c.input, got, c.expected)
		}
	}
}
