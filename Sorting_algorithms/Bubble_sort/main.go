// Stable: Yes
// Best Case: O(n^2)
// Worst Case: O(n^2)
// Average Case: O(n^2)
// Space Complexity: O(1)

package main

import (
	"fmt"
)

// Explanation: https://www.programiz.com/dsa/bubble-sort

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	a := []int{1, 5, 3, 2, 4}

	bubbleSort(a)

	fmt.Printf("Sorted array: %v\n", a)
}

// Output: "Sorted array: [1 2 3 4 5]"
