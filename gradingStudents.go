package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	var results []int32
	for _, value := range grades {
		rest := value % 5
		var result int32

		if (value < 38) || (rest < 3) {
			result = value
		} else {
			result = value + (5 - rest)
		}

		results = append(results, result)
	}

	return results
}

func main() {
	grades := []int32{75, 32, 43, 59, 35, 95, 43, 68}

	result := gradingStudents(grades)

	fmt.Printf("Result: %d\n", result)
}
