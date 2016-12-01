// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import (
	"strings"
	"sync"
	"sync/atomic"
)

func (tr *TR) recurse(seed string, wg *sync.WaitGroup,
	paths chan<- string, path string, depth int) {
	//fmt.Println("Enter recurse()")
	defer wg.Done()
	tr.setMaxDepth(depth)

	// test for cycles by examining path for presence of seed
	herepath := path + tr.Delimiter
	tmp := tr.Delimiter + seed + tr.Delimiter
	if strings.Contains(herepath, tmp) {
		paths <- herepath + seed + tr.Delimiter + "#CYCLE#"
		return
	}
	herepath += seed

	// add to the visited list
	_ = tr.v.exists(seed)

	nextLevel := tr.query(seed)
	if len(nextLevel) == 0 {
		paths <- herepath + tr.Delimiter + "#LEAF#"
		return
	}

	for _, v := range nextLevel {
		wg.Add(1)
		go tr.recurse(v, wg, paths, herepath, depth+1)
		atomic.AddInt64(&tr.goroutineCount, 1)
	}
	//fmt.Println("Exit recurse()")
}
