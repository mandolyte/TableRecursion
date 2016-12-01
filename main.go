// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import (
	"encoding/csv"
	"io"
	"sync"
)

// TR is the struct with the "global" package variables
type TR struct {
	Table       [][]string
	Delimiter   string
	Threads     int
	Info        *csv.Writer
	Paths       *csv.Writer
	Results     *csv.Writer
	IncludeSeed bool
	// info vars
	goroutineCount int64
	maxDepth       int
	partcount      int
	pathcount      int
	// various other "global vars"
	sema          chan struct{}
	mutexMaxDepth *sync.Mutex
	v             *visited
}

func (tr *TR) setMaxDepth(depth int) {
	tr.mutexMaxDepth.Lock()
	if depth > tr.maxDepth {
		tr.maxDepth = depth
	}
	tr.mutexMaxDepth.Unlock()
}

// NewTR is the allocator for the package; variables that
// are settable may use the exported struct members directly
func NewTR(table [][]string, dlm string,
	t int, i, p, r io.Writer, inclseed bool) *TR {

	winfo := csv.NewWriter(i)
	wpath := csv.NewWriter(p)
	wpart := csv.NewWriter(r)

	tr := &TR{
		Table:         table,
		Delimiter:     dlm,
		Threads:       t,
		Info:          winfo,
		Paths:         wpath,
		Results:       wpart,
		IncludeSeed:   inclseed,
		sema:          make(chan struct{}, t),
		mutexMaxDepth: &sync.Mutex{},
	}

	tr.info("INFO", "Delimiter value is %v", tr.Delimiter)
	tr.info("INFO", "Number of threads to use is %v", tr.Threads)
	tr.info("INFO", "Include seed in results:%v", tr.IncludeSeed)

	return tr
}
