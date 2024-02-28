package breakfastalg

import (
	"slices"
)

func WrongSimpleAlg(lvls []int) [][]int {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	prev := make([]int, len(lvls))
	out := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, prev)
		next := make([]int, len(lvls))
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}
	}
	return out
}

func RightSimpleAlg(lvls []int) [][]int {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	prev := make([]int, len(lvls))
	out := make([][]int, 0, n)
	it := NewIncrementalTree(lvls)
	for i := 0; i < n; i++ {
		out = append(out, prev)
		next := make([]int, len(lvls))
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}
	restart:
		for _, v := range out {
			if slices.Equal(prev, v) {
				prev = it.Path(it.Index(prev) + 1)
				goto restart
			}
		}
	}
	return out
}

func RightFastAlg(lvls []int) [][]int {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	prev := make([]int, len(lvls))
	out := make([][]int, 0, n)
	out = append(out, prev)
	it := NewIncrementalTree(lvls)
	alreadydone := make([]bool, n)
	alreadydone[0] = true
	for i := 1; i < n; i++ {
		next := make([]int, len(lvls))
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}

		for alreadydone[it.Index(prev)] {
			it.Increment(prev)
		}
		out = append(out, prev)
		alreadydone[it.Index(prev)] = true
	}
	return out
}

type FlatArr[T any] struct {
	backing []T
	width   int
}

func (f *FlatArr[T]) Append(v []T) {
	// assert len(v) == width
	f.backing = append(f.backing, v...)
}
func (f *FlatArr[T]) ConvertToRegular() [][]T {
	out := make([][]T, 0, len(f.backing)/f.width)
	for i := 0; i < len(f.backing); i += f.width {
		t := make([]T, f.width)
		copy(t, f.backing[i:i+f.width])
		out = append(out, t)
	}
	return out
}

func RightFasterAlg(lvls []int) FlatArr[int] {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	width := len(lvls)
	out := make([]int, n*width)
	previndex := 0
	prev := out[previndex : previndex+width]
	it := NewIncrementalTree(lvls)
	alreadydone := make([]bool, n)
	alreadydone[0] = true
	for i := 1; i < n; i++ {
		previndex += width
		next := out[previndex : previndex+width]
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}

		for alreadydone[it.Index(prev)] {
			it.Increment(prev)
		}
		alreadydone[it.Index(prev)] = true
	}
	return FlatArr[int]{out, width}
}

// actually slower :(
// but might be a good idea
func RightEvenFasterAlg(lvls []int) FlatArr[int] {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	width := len(lvls)
	out := make([]int, n*width)
	previndex := 0
	prev := out[previndex : previndex+width]
	it := NewIncrementalTree(lvls)
	alreadydone := make([]int, n)
	alreadydone[0] = 1
	for i := 1; i < n; i++ {
		previndex += width
		next := out[previndex : previndex+width]
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}

		idx := it.Index(prev)
		firstindex := idx
		for alreadydone[idx] != 0 {
			idx = (idx + alreadydone[idx]) % n
		}
		copy(prev, it.Path(idx))
		alreadydone[idx] = 1
		alreadydone[firstindex] = (idx - firstindex) + 1
	}
	return FlatArr[int]{out, width}
}
