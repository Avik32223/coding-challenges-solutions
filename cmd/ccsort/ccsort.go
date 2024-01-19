package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_sort "github.com/Avik32223/coding-challenges-solutions/pkg/sort"
)

type algorithmType int

const (
	QuickSort algorithmType = iota
	HeapSort
	RandomSort
)

type algorithm func([]string) []string

func sort(s algorithmType, source []string) ([]string, error) {
	var fn algorithm
	switch s {
	case QuickSort:
		fn = _sort.QuickSort
	case HeapSort:
		fn = _sort.HeapSort
	case RandomSort:
		fn = _sort.RandomSort
	default:
		panic("no matching sort algorithm found")
	}
	result := fn(source)
	return result, nil
}

func uniq(s []string) []string {
	result := make([]string, 0)
	for _, i := range s {
		if len(result) == 0 || result[len(result)-1] != i {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	var uniqFlag bool
	var quickSortFlag bool
	var heapSortFlag bool
	var randomSortFlag bool
	flag.BoolVar(&uniqFlag, "u", false, "unique lines only")
	flag.BoolVar(&quickSortFlag, "qsort", false, "use quick sort")
	flag.BoolVar(&heapSortFlag, "heapsort", false, "use heap sort")
	flag.BoolVar(&randomSortFlag, "randomsort", false, "use random sort")
	flag.Parse()
	var sortAlgorithm algorithmType
	switch {
	case quickSortFlag:
		sortAlgorithm = QuickSort
	case heapSortFlag:
		sortAlgorithm = HeapSort
	case randomSortFlag:
		sortAlgorithm = RandomSort
	default:
		sortAlgorithm = QuickSort
	}

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	filename := flag.Arg(0)
	if len(filename) != 0 {
		f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		reader = bufio.NewReader(f)
	}

	source := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t := scanner.Text()
		t = strings.Trim(t, " ")
		source = append(source, t)
	}

	result, err := sort(sortAlgorithm, source)
	if err != nil {
		panic(err)
	}
	if uniqFlag {
		result = uniq(result)
	}

	defer writer.Flush()
	for _, i := range result {
		if _, err := writer.WriteString(fmt.Sprintln(i)); err != nil {
			panic(err)
		}
	}
}
