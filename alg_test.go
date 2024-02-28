package breakfastalg

import (
	"slices"
	"testing"
)


func Mul(l []int) int {
	out := 1
	for _, v := range l {
		out *= v
	}
	return out
}

func AnyRepeat(a [][]int) (bool, int, int) {
	for i := 0; i < len(a); i++ {
		for k := i + 1; k < len(a); k++ {
			if slices.Equal(a[i], a[k]) {
				return true, i, k
			}
		}
	}
	return false, -1, -1
}

func BetterAnyAdjEqual(ll [][]int) (bool, int) {
	for i := 0; i < len(ll)-1; i++ {
		if slices.Equal(ll[i], ll[i+1]) {
			return true, i
		}
	}
	return false, -1
}

func CoprimeTest(t *testing.T, alg func([]int) [][]int) {
	lvls := [][]int{
		{3, 2},
		{2, 3},
		{5, 9},
		{9, 5},
		{2, 3, 5},
		{2, 3, 5, 7},
		{7, 5, 3, 2},
		{7, 3, 2, 5},
	}
	for _, lvl := range lvls {
		res := alg(lvl)
		if len(res) != Mul(lvl) {
			t.Errorf("wrong length for result with test %v, res: %v", lvl, res)
		}
		yes, i := BetterAnyAdjEqual(res)
		if yes {
			t.Errorf("there are at least two adjacent indicies (%d, %d) where any of the subindicies are equal, %v", i, i+1, res)
		}

		yes, i, k := AnyRepeat(res)
		if yes {
			t.Errorf("there are at least two indicies (%d, %d) that are the same, %v", i, k, res)
		}
	}
}

func NonCoprimeTest(t *testing.T, alg func([]int) [][]int) {
	lvls := [][]int{
		{3, 3},
		{4, 4},
		{5, 5},
		{2, 4, 6},
		{3, 6, 9},
		{9, 6, 4, 2},
		{3, 15},
		{15, 3},
	}
	for _, lvl := range lvls {
		res := alg(lvl)
		if len(res) != Mul(lvl) {
			t.Error("wrong length for result")
		}
		yes, i := BetterAnyAdjEqual(res)
		if yes {
			t.Errorf("there are at least two adjacent indicies (%d, %d) where any of the subindicies are equal, %v", i, i+1, res)
		}

		yes, i, k := AnyRepeat(res)
		if yes {
			t.Errorf("there are at least two indicies (%d, %d) that are the same, %v", i, k, res)
		}
	}
}

func TestWrongSimpleAlg(t *testing.T) {
	CoprimeTest(t, WrongSimpleAlg)
}

func TestRightSimpleAlg(t *testing.T) {
	CoprimeTest(t, RightSimpleAlg)
	NonCoprimeTest(t, RightSimpleAlg)
}
func TestRightFastAlg(t *testing.T) {
	CoprimeTest(t, RightFastAlg)
	NonCoprimeTest(t, RightFastAlg)
}

func ConvertedRightFasterAlg(lvls []int) [][]int {
	t := RightFasterAlg(lvls)
	return t.ConvertToRegular()
}
func TestRightFasterAlg(t *testing.T) {
	CoprimeTest(t, ConvertedRightFasterAlg)
	NonCoprimeTest(t, ConvertedRightFasterAlg)
}

func ConvertedRightEvenFasterAlg(lvls []int) [][]int {
	t := RightEvenFasterAlg(lvls)
	return t.ConvertToRegular()
}
func TestRightEvenFasterAlg(t *testing.T) {
	CoprimeTest(t, ConvertedRightEvenFasterAlg)
	NonCoprimeTest(t, ConvertedRightEvenFasterAlg)
}

func BenchmarkRightSimpleAlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lvls := []int{10, 20, 30}
		_ = RightSimpleAlg(lvls)
	}
}

func BenchmarkRightFastAlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lvls := []int{10, 20, 30}
		_ = RightFastAlg(lvls)
	}
}

func BenchmarkRightFasterAlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lvls := []int{10, 20, 30}
		_ = RightFasterAlg(lvls)
	}
}
func BenchmarkRightEvenFasterAlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lvls := []int{10, 20, 30}
		_ = RightEvenFasterAlg(lvls)
	}
}
