# TableRecursion
Package TableRecursion is used to recursively process a 2D
table with the first column being the "parent" and the second
being the "child". It produces two things:
a) all the parts to all levels
b) all the paths to every part
the paths are annotated with either "#LEAF#" for paths from
root to leaf; or "#CYCLE#" for paths that are cyclic

There is a tester program in the `cmd` folder. Here is the table used
in the tester:
```
atable := [][]string{
  {"A", "B"},
  {"A", "C"},
  {"B", "D"},
  {"C", "E"},
  {"D", "F"},
  {"D", "A"},
  {"E", "G"},
}
```

Here is the execution and outputs:
```
$ go run tester1.go
$ ls
tester1.go  tester1_info.csv  tester1_parts.csv  tester1_paths.csv
$ cat tester1_paths.csv
A,/A/C/E/G/#LEAF#
A,/A/B/D/A/#CYCLE#
A,/A/B/D/F/#LEAF#
$ cat tester1_parts.csv
A,G
A,B
A,D
A,F
A,A
A,C
A,E
$ cat tester1_info.csv
INFO,Delimiter value is /
INFO,Number of threads to use is 10
INFO,Include seed in results:true
A,Started at: 2016-11-30 19:41:15.9776685 -0500 EST
A,Total number of threads used:8
A,Maximum depth:4
A,Total number of paths:3
A,Total number of parts:7
A,Elapsed time:253.351µs
```
