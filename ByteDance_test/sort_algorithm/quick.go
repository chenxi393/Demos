package main

// https://www.cnblogs.com/MAKISE004/p/16909610.html
// 可以参考这里的写法 当时感觉差不多 下面第一种 简洁一点
// 快速排序
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {
		if arr[left] <= pivot {
			left++
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}

	arr[0], arr[left-1] = arr[left-1], arr[0]
	quickSort(arr[:left-1])
	quickSort(arr[left:])
}
// 1 2 3 4 5

// ！！！！！！ 注意到 上述排序在有序的情况下 会打乱后面有序的排序
// 下面的不会 会退化为冒泡排序

// this quickSort practice is uausally used by myself
// 时间和上面的差不多 在有序情况下上面的更优 原因在上面（会打乱后面的序列）
// 所以上面在有序的情况下仍然为O(n*logn) 但下面为n^2
func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1
	// 注意数组越界的问题    left < right
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		if left < right {
			arr[right], arr[left] = arr[left], arr[right]
			left++
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[right], arr[left] = arr[left], arr[right]
			right--
		}
	}
	arr[left] = pivot
	quickSort2(arr[:left])
	quickSort2(arr[left+1:])
}
