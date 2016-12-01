// Package TableRecursion is used to recursively process a 2D
// table with the first column being the "parent" and the second
// being the "child". It produces two things:
// a) all the parts to all levels
// b) all the paths to every part
// the paths are annotated with either "#LEAF#" for paths from
// root to leaf; or "#CYCLE#" for paths that are cyclic
package TableRecursion

import "fmt"

func (tr *TR) info(seed, format string, v ...interface{}) {
	if tr.Info == nil {
		return
	}
	tr.Info.Write([]string{seed, fmt.Sprintf(format, v...)})
}

func (tr *TR) paths(seed, format string, v ...interface{}) {
	if tr.Paths == nil {
		return
	}
	tr.Paths.Write([]string{seed, fmt.Sprintf(format, v...)})
}

func (tr *TR) results(seed, format string, v ...interface{}) {
	if tr.Results == nil {
		return
	}
	tr.Results.Write([]string{seed, fmt.Sprintf(format, v...)})
}
