package main

import (
	"fmt"
)

func IsPat(s string) bool {
	if len(s) == 1 {
		return true
	}

	for i := 1; i < len(s); i++ {
		l, r := s[:i], s[i:]

		if SMin(l) > SMax(r) && IsPat(SReverse(l)) && IsPat(SReverse(r)) {
			return true
		}
	}

	return false
}

func SMin(s string) byte {
	minChar := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
		}
	}
	return minChar
}

func SMax(s string) byte {
	maxChar := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > maxChar {
			maxChar = s[i]
		}
	}
	return maxChar
}

func SReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
    var s1, s2 string
    fmt.Scan(&s1, &s2)

	if IsPat(s1) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
	}

    if IsPat(s2) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }

    if IsPat(s1 + s2) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}