package main

import (
	"fmt"
)

// BubbleSort 冒泡排序，稳定，O(n^2)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				//如果当前元素大于下一个元素，交换他们
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort 选择排序，不稳定，O(n^2)
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//将找到的最小的元素交换到正确的位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 插入排序，稳定，O(n^2)
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			// 将元素前移
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

// quickSort 快速排序，不稳定，平均O(n*logn)，最坏O(n^2)
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := nums[r]
	i := l
	j := r - 1

	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for j >= i && nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[i], nums[r] = nums[r], nums[i]

	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

// mergeSort 归并排序，稳定，O(n*logn)
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// heapSort 堆排序，O(n*logn)，不稳定
func heapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		//把堆顶元素交换到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// shellSort 希尔排序，最好O(n)最坏O(n^2)，不稳定
func shellSort(arr []int) {
	n := len(arr)
	//步长
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2
	}
}

// radixSort 基数排序，O(kn)，稳定
func radixSort(arr []int) {
	max := findMax(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}

func findMax(arr []int) int {
	max := arr[0]

	for _, val := range arr[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) //输出数组
	count := make([]int, 10) //计数数组

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	// 累加计数器，计数数组转化为排序后的下标数组
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 填入排序后的下标
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	//quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
