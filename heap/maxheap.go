package heap

import "cmp"

type MaxHeap[T cmp.Ordered] struct {
	data []T
	size int
}

func (h *MaxHeap[T]) Insert(val T) {
	h.data = append(h.data, val)
	h.size++

	// sift up
	currIdx := h.size - 1
	for currIdx > 0 && h.data[currIdx] > h.data[(currIdx-1)/2] {
		h.data[currIdx], h.data[(currIdx-1)/2] = h.data[(currIdx-1)/2], h.data[currIdx]
		currIdx = (currIdx - 1) / 2
	}
}

func (h *MaxHeap[T]) RemoveMax() {
	if h.size == 0 {
		panic("Attempt to read value from an empty heap")
	}

	h.data[0] = h.data[h.size-1]
	h.size--

	// sift down
	currIdx := 0
	for currIdx*2+1 < h.size {

		// swap with left child by default
		changeIdx := 2*currIdx + 1

		if 2*currIdx+2 < h.size && h.data[2*currIdx+2] > h.data[2*currIdx+1] {
			// if right child exists and is smaller
			changeIdx = 2*currIdx + 2
		}

		if h.data[currIdx] > h.data[changeIdx] {
			break
		}

		h.data[currIdx], h.data[changeIdx] = h.data[changeIdx], h.data[currIdx]
		currIdx = changeIdx
	}
}

func (h *MaxHeap[T]) GetMax() T {
	if h.size == 0 {
		panic("Attempt to read value from an empty heap")
	}

	return h.data[0]
}

func (h *MaxHeap[T]) GetSize() int {
	return h.size
}
