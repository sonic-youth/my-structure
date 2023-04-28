package sort

type Slice interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

// QuickSort just use quick sort algorithm to sort the slice you provided
func QuickSort(data Slice) {
	quickSort(data, 0, data.Len()-1)
}
func quickSort(data Slice, start int, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	for i < j {
		// start is the partition
		for i < j && data.Less(i, j) {
			j--
		}
		if i < j {
			// swap j, partition
			data.Swap(j, i)
			i++
		}
		// now j is the partition
		for i < j && data.Less(i, j) {
			i++
		}
		if i < j {
			// swap i, partition
			data.Swap(i, j)
			j--
		}
	}
	quickSort(data, start, i-1)
	quickSort(data, i+1, end)
}
