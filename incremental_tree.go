package breakfastalg

// not a normal tree, try to find a better name
type IncrementalTree struct {
	levels      []int
	levelWidths []int
	max         int
}

func (t IncrementalTree) Max() int {
	return t.max
}

func NewIncrementalTree(levels []int) IncrementalTree {
	ourlevels := make([]int, len(levels))
	copy(ourlevels, levels)
	levelWidths := make([]int, len(levels))
	width := 1
	for i := len(levels) - 1; i > 0; i-- {
		width *= levels[i]
		levelWidths[i-1] = width
	}
	width *= levels[0]
	levelWidths[len(levelWidths)-1] = 1
	return IncrementalTree{
		levels:      ourlevels,
		levelWidths: levelWidths,
		max:         width,
	}
}

// returns the path to get that index
// make this round
func (it IncrementalTree) Path(n int) []int {
	n = n % (it.max + 1) // can comment this out and most tests should work ;)
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

func (it IncrementalTree) Increment(path []int) {
	for i := len(path) - 1; i >= 0; i-- {
		path[i]++
		if path[i] == it.levels[i] {
			path[i] = 0
		} else {
			break
		}
	}
}
func (it IncrementalTree) IncrementWithIndex(path []int, prevIndex int) int {
	firstinital := path[0]
	it.Increment(path)
	if firstinital != 0 && path[0] == 0 {
		return 0
	}
	return prevIndex + 1
}
