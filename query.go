// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import "sort"

func (tr *TR) query(seed string) []string {
	tr.sema <- struct{}{}        // acquire token, max 20
	defer func() { <-tr.sema }() // release token

	// do the search...
	first := sort.Search(len(tr.Table),
		func(i int) bool { return tr.Table[i][0] >= seed })

	var results []string
	for i := first; i < len(tr.Table); i++ {
		if tr.Table[i][0] == seed {
			results = append(results, tr.Table[i][1])
			continue
		}
		break
	}
	return results
}
