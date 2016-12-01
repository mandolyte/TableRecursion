// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

func (t *TR) Len() int {
	return len(t.Table)
}

func (t *TR) Swap(i, j int) {
	t.Table[i], t.Table[j] = t.Table[j], t.Table[i]
}

func (t *TR) Less(i, j int) bool {
	return t.Table[i][0] < t.Table[j][0]
}
