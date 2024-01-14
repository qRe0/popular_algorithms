// Stable: Yes
// Best Case: O(n)
// Worst Case: O(n^2)
// Average Case: O(n^2)
// Space Complexity: O(1)

package main

import (
	"fmt"
)

// Explanation: https://www.programiz.com/dsa/insertion-sort

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]

			j--
		}

		arr[j+1] = key
	}
}

func main() {
	a := []int{1, 5, 3, 2, 4}

	insertionSort(a)

	fmt.Printf("Sorted array: %v\n", a)
}

// Output: "Sorted array: [1 2 3 4 5]"
