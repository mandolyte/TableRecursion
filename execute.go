// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import (
	"sync"
	"sync/atomic"
	"time"
)

// Execute the main method to initialize and begin the recursion
func (tr *TR) Execute(seed string) {
	now := time.Now()

	// reset variables
	tr.goroutineCount = 0
	tr.maxDepth = 0
	tr.pathcount = 0
	tr.partcount = 0
	tr.info(seed, "Started at: %v", now)
	tr.v = newVisited()

	paths := make(chan string, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go tr.recurse(seed, &wg, paths, "", 1)
	atomic.AddInt64(&tr.goroutineCount, 1)

	// this needed to actually close the channel when all the
	// go routines are done. Otherwise, this error will be
	// returned when everything is done:
	// fatal error: all goroutines are asleep - deadlock!
	go func() {
		wg.Wait()
		close(paths)
	}()

	for {
		path, ok := <-paths
		if !ok {
			break
		}
		tr.paths(seed, "%v", path)
		tr.pathcount++
	}

	for k := range tr.v.getMap() {
		if k == seed {
			if tr.IncludeSeed {
				// no op
			} else {
				continue
			}
		}
		tr.results(seed, "%v", k)
		tr.partcount++
	}
	stop := time.Since(now)

	tr.info(seed, "Total number of threads used:%v", tr.goroutineCount)
	tr.info(seed, "Maximum depth:%v", tr.maxDepth)
	tr.info(seed, "Total number of paths:%v", tr.pathcount)
	tr.info(seed, "Total number of parts:%v", tr.partcount)
	tr.info(seed, "Elapsed time:%v", stop)
	tr.Info.Flush()
	tr.Paths.Flush()
	tr.Results.Flush()
}
