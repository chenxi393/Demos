package main

// 堆排序
func heapSort(arr []int) {
	buildHeap(arr)

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func buildHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
}

func heapify(arr []int, root, length int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < length && arr[left] > arr[largest] {
		largest = left
	}

	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, largest, length)
	}
}