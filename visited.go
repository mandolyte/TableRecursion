// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import "sync"

var mutexVisited = &sync.Mutex{}

type visited struct {
	m map[string]int
}

// NewVisited initializes the map for parts found for root
func newVisited() *visited {
	x := new(visited)
	x.m = make(map[string]int)
	return x
}

func (v *visited) exists(part string) bool {
	mutexVisited.Lock()
	_, ok := v.m[part]
	if !ok {
		v.m[part] = 1
	}
	mutexVisited.Unlock()
	return ok
}

func (v *visited) getMap() map[string]int {
	return v.m
}
