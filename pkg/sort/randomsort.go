package sort

import (
	"bufio"
	"crypto/rand"
	"hash/fnv"

	"github.com/Avik32223/coding-challenges-solutions/internal/heap"
)

type randomHeapItem struct {
	val  string
	hash uint32
}

func (h randomHeapItem) Compare(c heap.Node) int {
	if ch, ok := c.(randomHeapItem); ok {
		if h.hash > ch.hash {
			return 1
		} else if h.hash < ch.hash {
			return -1
		}
	}
	return 0
}

func randomSort(arr []string) []string {
	reader := bufio.NewReader(rand.Reader)
	seed, _ := reader.ReadBytes(0)
	hashGen := fnv.New32a()
	results := make([]string, 0)
	mHeap := heap.NewMinHeap()
	for _, i := range arr {
		hashGen.Reset()
		hashGen.Write(seed)
		hashGen.Write([]byte(i))
		h := hashGen.Sum32()
		a := randomHeapItem{
			val:  i,
			hash: h,
		}
		mHeap.Insert(a)
	}
	for mHeap.Size() > 0 {
		a, _ := mHeap.Extract()
		results = append(results, a.(randomHeapItem).val)
	}
	return results
}

func RandomSort(arr []string) (c_arr []string) {
	c_arr = randomSort(arr)
	return
}
