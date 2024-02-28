package breakfastalg

import (
	"slices"
	"testing"
)


func TestBasic(t *testing.T) {
	lvls := []int{10}
	tr := NewIncrementalTree(lvls)
	lvls[0] = 100
	if tr.levels[0] == lvls[0] {
		t.Error("the new increment tree should have its own version of levels")
	}

	lvls = []int{3, 2}
	tr = NewIncrementalTree(lvls)
	if !slices.Equal([]int{0,0}, tr.Path(0)) {
		t.Error("wrong zero path")
	}
	if !slices.Equal([]int{0,1}, tr.Path(1)) {
		t.Error("wrong one path")
	}
	if !slices.Equal([]int{1,0}, tr.Path(2)) {
		t.Errorf("wrong two path %v", tr.Path(2))
	}
}

func TestPathAndIndex(t *testing.T) {
	tr := NewIncrementalTree([]int{1,2,3,4,5})
	for i := 0; i < tr.Max(); i++ {
		if tr.Index(tr.Path(i)) != i {
			t.Errorf("error here")
		}
	}
}


func hlpr(t *testing.T, a []int, b []int) {
	t.Helper()
	if !slices.Equal(a ,b) {
		t.Errorf("slices not equal %v, %v", a, b)
	}
}
func TestIncrement(t *testing.T) {
	tr := NewIncrementalTree([]int{2, 2})
	p := tr.Path(0)
	hlpr(t, []int{0,0}, p)
	tr.Increment(p)
	hlpr(t, []int{0,1}, p)
	tr.Increment(p)
	hlpr(t, []int{1,0}, p)
	tr.Increment(p)
	hlpr(t, []int{1,1}, p)
	tr.Increment(p)
	hlpr(t, []int{0,0}, p)
}
