package list_number

import (
"fmt"
"log"
)

func main() {
	numbers := []int{1, 2, 3, 4, 9, 5, 6}

	if len(numbers) == 0 {
		log.Fatal("Slice is empty")
	}

	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	fmt.Printf("The maximum number is: %d\n", max)
}
