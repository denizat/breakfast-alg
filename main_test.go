package main

import "testing"

// not a normal tree, try to find a better name
type IncrementalTree struct {
	levels      []int
	levelWidths []int
	max         int
	// Arr []bool
}

func NewIncrementalTree(levels []int) IncrementalTree {
	levelWidths := make([]int, len(levels))
	width := 1
	for i := len(levels) - 1; i > 0; i-- {
		width *= levels[i]
		levelWidths[i-1] = width
	}
	width *= levels[0]
	levelWidths[len(levelWidths)-1] = 1
	// arr := make([]bool, width)
	return IncrementalTree{
		levels:      levels,
		levelWidths: levelWidths,
		max:         width,
		// Arr: arr
	}
}

// returns the path to get that index
func (it IncrementalTree) Path(n int) []int {
	path := make([]int, len(it.levels))
	cur := 0
	for i, lw := range it.levelWidths {
		k := 1
		for ; n >= cur+lw*k; k++ {
		}
		path[i] = k - 1
		cur += lw * (k - 1)
	}
	return path
}

func (it IncrementalTree) Index(path []int) int {
	out := 0
	for i := range path {
		out += it.levelWidths[i] * path[i]
	}
	return out
}

func AnyEqual[T comparable](al, bl []T) bool {
	top := min(len(al), len(bl))
	for i := 0; i < top; i++ {
		if al[i] == bl[i] {
			return true
		}
	}
	return false
}

func ExactlyEqual[T comparable](al, bl []T) bool {
	top := min(len(al), len(bl))
	for i := 0; i < top; i++ {
		if al[i] != bl[i] {
			return false
		}
	}
	return true
}

func levelsalg(lvls []int8) [][]int8 {
	var n int = 1
	for _, i := range lvls {
		n *= int(i)
	}
	prev := make([]int8, len(lvls))
	out := make([][]int8, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, prev)
		next := make([]int8, len(lvls))
		copy(next, prev)
		prev = next
		for i := range prev {
			prev[i] = (prev[i] + 1) % lvls[i]
		}
	}
	return out
}

func levelsalg2(lvls []int) [][]int {
	it := NewIncrementalTree(lvls)
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
	restart:
		for _, v := range out {
			if ExactlyEqual(v, prev) {
				i := it.Index(prev) + 1
				prev = it.Path(i)
				goto restart
			}
		}
		if AnyEqual(prev, out[len(out)-1]) {
			i := it.Index(prev) + 1
			prev = it.Path(i)
			goto restart
		}
	}
	return out
}

func TestMain(t *testing.T) {
	// lvls := []int8{10,10, 11}
	// res := levelsalg(lvls)
	// t.Log(res)
}

func TestLevelStructure(t *testing.T) {
	lvls := []int{3, 3, 2}
	nt := NewIncrementalTree(lvls)
	t.Logf("%v\n", nt)
	res := nt.Path(17)
	t.Logf("%v\n", res)
}

func TestLevels2(t *testing.T) {
	lvls := []int{3}
	res := levelsalg2(lvls)
	t.Log(res)
}

func AnyAdjEqual(ll [][]int) bool {
	for i := 0; i < len(ll)-1; i++ {
		if AnyEqual(ll[i], ll[i+1]) {
			return true
		}
	}
	return false
}

func Mul(l []int) int {
	out := 1
	for _, v := range l {
		out *= v
	}
	return out
}

func TestLevels3(t *testing.T) {
	for i := 1; i < 7; i++ {
		levels := make([]int, i)
		for i := range levels {
			levels[i] = i + 3
		}
		res := levelsalg2(levels)
		if len(res) != Mul(levels) {
			t.Error("result not long enough")
		}
		if AnyAdjEqual(res) {
			t.Errorf("result not correct for %v", levels)
		}
	}
}
func BenchmarkLevels2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levelsalg2([]int{3, 4, 5})
	}
}
