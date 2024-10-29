package heap_test

import (
	"cmp"
	"sort"
	"testing"

	"go-algo/heap"
)

func deepCopyAndSort[T cmp.Ordered](in []T) []T {
	out := append([]T{}, in...)

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out
}

func TestMinHeap(t *testing.T) {

	t.Run("int", func(t *testing.T) {
		in := []int{-74, 58, 63, 27, -48, -87, 49, -26, -99, -41, 32, 15, 77, -5, 68, -28, -11, 84, 6, -19}
		out := deepCopyAndSort(in)

		h := heap.MinHeap[int]{}

		for _, v := range in {
			h.Insert(v)
		}

		for i, v := range out {
			if m := h.GetMin(); m != v {
				t.Errorf("expected %v, got %v, itteration %d", v, m, i)
			}

			h.RemoveMin()

			if h.GetSize() != len(out)-i-1 {
				t.Errorf("expected size %d, got %d, itteration %d", len(out)-i-1, h.GetSize(), i)
			}
		}
	})

	t.Run("string", func(t *testing.T) {
		in := []string{
			"", "trqp", "jg", "m", "ba", "vxxe", "q", "nm", "r", "lk", "e", "oy", "pz", "i", "d", "z", "xu", "abc", "defg", "abcd", "fgt", "nopq"}
		out := deepCopyAndSort(in)

		h := heap.MinHeap[string]{}

		for _, v := range in {
			h.Insert(v)
		}

		for i, v := range out {
			if m := h.GetMin(); m != v {
				t.Errorf("expected %v, got %v, iteration %d", v, m, i)
			}

			h.RemoveMin()

			if h.GetSize() != len(out)-i-1 {
				t.Errorf("expected size %d, got %d, iteration %d", len(out)-i-1, h.GetSize(), i)
			}
		}
	})
}
