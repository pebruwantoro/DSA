package main

import (
	"fmt"
)

type Case struct {
	input  []int
	target int
}

func main() {
	cases1 := []Case{
		{
			input:  []int{2, 7, 11, 15},
			target: 9,
		},
		{
			input:  []int{3, 2, 4},
			target: 6,
		},
		{
			input:  []int{3, 3, 7, 8, 6, 9},
			target: 6,
		},

		{
			input:  []int{3, 3, 4, 3, 2, 1, 7, 89},
			target: 1,
		},

		{
			input:  []int{100, 21, 23, 45, 78, 88, 1000},
			target: 88,
		},

		{
			input:  []int{100, 1001, 102, 103, 1020},
			target: 103,
		},
	}

	cases2 := []Case{
		{
			input:  []int{2, 7, 11, 15},
			target: 9,
		},
		{
			input:  []int{2, 3, 4},
			target: 6,
		},
		{
			input:  []int{2, 3, 6, 7, 8, 9},
			target: 6,
		},

		{
			input:  []int{1, 3, 4, 7, 9, 10, 17, 89},
			target: 17,
		},

		{
			input:  []int{10, 21, 23, 45, 78, 88, 1000},
			target: 88,
		},

		{
			input:  []int{100, 1001, 1020, 1030, 1020},
			target: 103,
		},
	}

	for _, v := range cases1 {
		idx, found := linearSearch(v.input, v.target)
		fmt.Println("data: ", v.target, "target: ", v.target, "found on index: ", idx, "found: ", v.target == found)
	}

	for _, v := range cases2 {
		idx, found := binarySearch(v.input, v.target)
		fmt.Println("data: ", v.target, "target: ", v.target, "found on index: ", idx, "found: ", v.target == found)
	}
}

func linearSearch(input []int, target int) (int, int) {
	for i, v := range input {
		if v == target {
			return i, target
		}
	}

	return -1, -1 // target not found in the array

	/*
		Linear Search is the simplest search algorithm. It works by iterating over the entire array and checking each element.

		Algorithm:
		•	Start from the first element and check if it matches the target.
		•	Move through the list until the target is found or the list is exhausted.

		Time Complexity:
		•	Worst-case: O(n)
		•	Best-case: O(1) (if the target is found at the first position)
		•	Space Complexity: O(1)
	*/
}

func binarySearch(input []int, target int) (int, int) {
	left, right := 0, len(input)-1

	for left <= right {
		mid := left + (right-left)/2
		if input[mid] == target {
			return mid, target // Target found
		} else if input[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return -1, -1 // target not found in the array

	/*
		Binary Search is an efficient search algorithm used for sorted arrays. It works by repeatedly dividing the search interval in half.

		Algorithm:
		•	Start with the middle element of the array.
		•	If the middle element is equal to the target, return its index.
		•	If the target is smaller than the middle element, search in the left half of the array.
		•	If the target is larger, search in the right half of the array.
		•	Repeat until the target is found or the array is exhausted.

		Time Complexity:
		•	Worst-case: O(log n)
		•	Best-case: O(1) (if the target is at the middle)
		•	Space Complexity: O(1)
	*/
}
