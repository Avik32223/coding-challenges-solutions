package sort

import "github.com/Avik32223/coding-challenges-solutions/internal/heap"

type heapItem struct {
	val string
}

func (h heapItem) Compare(c heap.Node) int {
	if ch, ok := c.(heapItem); ok {
		if h.val > ch.val {
			return 1
		} else if h.val < ch.val {
			return -1
		}
	}
	return 0
}

func heapSort(arr []string) []string {
	results := make([]string, 0)
	mHeap := heap.NewMinHeap()
	for _, i := range arr {
		a := heapItem{
			val: i,
		}
		mHeap.Insert(a)
	}
	for mHeap.Size() > 0 {
		a, _ := mHeap.Extract()
		results = append(results, a.(heapItem).val)
	}
	return results
}

func HeapSort(arr []string) (c_arr []string) {
	c_arr = heapSort(arr)
	return
}
