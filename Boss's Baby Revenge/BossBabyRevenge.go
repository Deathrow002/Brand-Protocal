package main

import (
	"fmt"
	"strings"
)

func BossBabyRevenge() {
	arr := []string{"SRSSRRR", "RSSRR", "SSSRRRRS", "SSRR", "SRRSSR"}
	for i := 0; i < len(arr); i++ {
		chars := arr[i]
		fmt.Println("Example ", i, ":", chars)
		rev := 0

		for j := 0; j < len(chars); j++ {
			char := string(chars[j])

			if char == "R" && strings.Index(chars, char) == 0 {
				rev = rev + 1
				break
			}

			if char == "S" {
				rev = rev + 1
			} else if char == "R" {
				rev = rev - 1
			}

			if rev < 0 {
				rev = 0
			}

			//fmt.Print(rev, ":", char, "\n")

		}

		if rev > 0 {
			fmt.Println("Output:", "Bad Boy")
		} else {
			fmt.Println("Output:", "Good Boy")
		}

	}
}
func main() {
	BossBabyRevenge()
}
