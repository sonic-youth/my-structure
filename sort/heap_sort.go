package sort

func HeapSort(data []int) {
	//TODO: implement it
}

// buildHeap build the origin slice to a heap slice
func buildHeap(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		adjustHeap(data, i)
	}
}

// adjustHeap will
func adjustHeap(data []int, i int) {
	temp := data[i]
	// from the left subNode of i, which is 2*i+1
	for k := i*2 + 1; k < len(data); k = k*2 + 1 {
		// if left subNode is less than right subNode, k will point to right subNode
		if k+1 < len(data) && data[k] < data[k+1] {
			k++
		}
		// if subNode is larger than parent i, will should update data[k] to parent
		if data[k] > temp {
			data[i] = data[k]
			i = k
		} else {
			break
		}
	}
	// the last subNode place is i, set temp to here
	data[i] = temp
}
