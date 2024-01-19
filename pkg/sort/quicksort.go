package sort

func partition(arr []string, low, high int) int {
	p := arr[high]
	left, right := low, low
	for right < high {
		if arr[right] <= p {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		right++
	}
	arr[left], arr[high] = arr[high], arr[left]
	return left
}

func quickSort(arr []string, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func QuickSort(arr []string) (c_arr []string) {
	c_arr = append(c_arr, arr...)
	quickSort(c_arr, 0, len(c_arr)-1)
	return
}
