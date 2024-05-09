package main

import "fmt"

func SupermanChickenRescue() {
	area1 := 5
	chickens1 := []int{2, 5, 10, 12, 15}
	fmt.Println("Example 1")
	fmt.Println("Rescue Chicken Number:", rescue(area1, chickens1))
}

func rescue(roof int, chickens []int) int {
	num_rescued := 0

	for i := 0; i < len(chickens); i++ {
		reach := chickens[i] + roof - 1
		count := 1

		for j := i; j < len(chickens); j++ {
			if reach >= chickens[j] {
				count += 1
			} else {
				break
			}
		}

		if count > num_rescued {
			num_rescued = count
		}
	}

	return num_rescued
}

func main() {
	SupermanChickenRescue()
}
