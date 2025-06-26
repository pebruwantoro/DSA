package main

import "fmt"

type Case struct {
	input []int
}

func main() {
	cases1 := []Case{
		{
			input: []int{1, 10, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{1123, 100, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{17, 60, 39, 200, 27, 54, 62, 12, 33, 99},
		},
	}

	cases2 := []Case{
		{
			input: []int{1, 10, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{1123, 100, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{17, 60, 39, 200, 27, 54, 62, 12, 33, 99},
		},
	}

	cases3 := []Case{
		{
			input: []int{1, 10, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{1123, 100, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{17, 60, 39, 200, 27, 54, 62, 12, 33, 99},
		},
	}

	cases4 := []Case{
		{
			input: []int{1, 10, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{1123, 100, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{17, 60, 39, 200, 27, 54, 62, 12, 33, 99},
		},
	}

	cases5 := []Case{
		{
			input: []int{1, 10, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{1123, 100, 19, 200, 2, 5, 6, 12, 33, 99},
		},
		{
			input: []int{17, 60, 39, 200, 27, 54, 62, 12, 33, 99},
		},
	}

	fmt.Println("BUBBLE SORT")
	for _, v := range cases1 {
		fmt.Println("arr before sorting: ", v.input)
		bubbleSort(v.input)
		fmt.Println("arr after sorting: ", v.input)
	}
	fmt.Println("=====================================================")

	fmt.Println("SELECTION SORT")
	for _, v := range cases2 {
		fmt.Println("arr before sorting: ", v.input)
		selectionSort(v.input)
		fmt.Println("arr after sorting: ", v.input)
	}
	fmt.Println("=====================================================")

	fmt.Println("INSERTION SORT")
	for _, v := range cases3 {
		fmt.Println("arr before sorting: ", v.input)
		insertionSort(v.input)
		fmt.Println("arr after sorting: ", v.input)
	}
	fmt.Println("=====================================================")

	fmt.Println("MERGE SORT")
	for _, v := range cases4 {
		fmt.Println("arr before sorting: ", v.input)
		mergeSort(v.input, 0, len(v.input)-1)
		fmt.Println("arr after sorting: ", v.input)
	}
	fmt.Println("=====================================================")

	fmt.Println("QUICK SORT")
	for _, v := range cases5 {
		fmt.Println("arr before sorting: ", v.input)
		quickSort(v.input, 0, len(v.input)-1)
		fmt.Println("arr after sorting: ", v.input)
	}
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		isSwapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	/*
		Bubble Sort is a simple comparison-based sorting algorithm. It repeatedly steps through the list, compares adjacent items, and swaps them if they are in the wrong order.
		•	Outer Loop (i): Iterates through the entire array.
		•	Inner Loop (j): Compares adjacent elements and swaps them if necessary.
		Algorithm:
		•	Compare each pair of adjacent items.
		•	Swap them if they are in the wrong order.
		•	Repeat the process until no swaps are needed.

		Time Complexity:
		•	Worst-case: O(n^2)
		•	Best-case: O(n) (when the array is already sorted)
		•	Space Complexity: O(1)
	*/
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j // Find the minimum element
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // Swap
	}

	/*
		Selection Sort works by repeatedly finding the minimum element from the unsorted part and swapping it with the first unsorted element.
		•	Outer Loop (i): Sets the position for the next smallest element.
		•	Inner Loop (j): Finds the smallest element in the unsorted portion.

		Algorithm:
			•	Iterate through the array.
			•	Find the minimum element from the unsorted part and swap it with the first unsorted element.

		Time Complexity:
			•	Worst-case: O(n^2)
			•	Best-case: O(n^2)
			•	Space Complexity: O(1)
	*/
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Shift elements of arr[0..i-1] that are greater than key
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key // Insert key at the correct position
	}

	/*
		Insertion Sort builds the sorted array one element at a time by inserting each element into its correct position.

		Algorithm:
			•	Start from the second element.
			•	Compare it with the previous elements and insert it in the correct position by shifting elements.

		Time Complexity:
			•	Worst-case: O(n^2)
			•	Best-case: O(n) (when the array is already sorted)
			•	Space Complexity: O(1)
	*/
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		// Recursively sort the left half
		mergeSort(arr, left, mid)
		// Recursively sort the right half
		mergeSort(arr, mid+1, right)

		// Merge the two sorted halves
		n1 := mid - left + 1
		n2 := right - mid
		leftArr := make([]int, n1)
		rightArr := make([]int, n2)

		copy(leftArr, arr[left:mid+1])
		copy(rightArr, arr[mid+1:right+1])

		i, j, k := 0, 0, left
		for i < n1 && j < n2 {
			if leftArr[i] <= rightArr[j] {
				arr[k] = leftArr[i]
				i++
			} else {
				arr[k] = rightArr[j]
				j++
			}
			k++
		}

		for i < n1 {
			arr[k] = leftArr[i]
			i++
			k++
		}

		for j < n2 {
			arr[k] = rightArr[j]
			j++
			k++
		}
	}

	/*
		Merge Sort is a divide-and-conquer algorithm that splits the array into halves, recursively sorts each half, and then merges the two sorted halves.

		Algorithm:
			•	Divide the array into two halves.
			•	Recursively sort each half.
			•	Merge the sorted halves.

		Time Complexity:
			•	Worst-case: O(n log n)
			•	Best-case: O(n log n)
			•	Space Complexity: O(n)

	*/
}

func quickSort(arr []int, low, height int) {

	if low < height {
		pivot := arr[height]
		i := low - 1

		for j := low; j < height; j++ {
			if arr[j] <= pivot {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

		arr[i+1], arr[height] = arr[height], arr[i+1]

		next := i + 1
		quickSort(arr, low, next-1)
		quickSort(arr, next+1, height)
	}

	/*
		Quick Sort is another divide-and-conquer algorithm. It works by selecting a “pivot” element and partitioning the array into two subarrays, such that elements smaller than the pivot go to the left, and elements greater go to the right.

		Algorithm:
			•	Select a pivot element.
			•	Partition the array into two subarrays.
			•	Recursively apply Quick Sort to the subarrays.

		Time Complexity:
			•	Worst-case: O(n^2)
			•	Best-case: O(n log n)
			•	Space Complexity: O(log n) (due to recursion)

	*/
}
