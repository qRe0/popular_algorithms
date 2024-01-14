// Time Complexity: O(log(n))
// Space Complexity: O(1)

package main

import (
	"fmt"
)

func binarySearch(arr []int, key int) []int {
	var ans []int

	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == key {
			ans = append(ans, mid)
		}

		if arr[mid] > key {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ans := binarySearch(a[:], 2)

	fmt.Printf("Found at index: %v\n", ans)
}
