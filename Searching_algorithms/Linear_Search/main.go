// Time Complexity: O(n)
// Space Complexity: O(1)

package main

import "fmt"

// Explanation: https://www.programiz.com/dsa/linear-search

func linearSearch(arr []int, key int) []int {
	var ans []int
	for i, elem := range arr {
		if elem == key {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	a := [...]int{1, 0, 3, 4, 0, 6, 7, 8, 9, 0}

	ans := linearSearch(a[:], 0)

	fmt.Printf("Found at indexes: %v\n", ans)
}

// Output: "Found at indexes: [1 4 9]"
