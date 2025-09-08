package main

import "fmt"

func CompLetters(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			count++
		}
	}
	return count
}
func main() {
	text := "mot"
	fmt.Print(CompLetters(text))
}
