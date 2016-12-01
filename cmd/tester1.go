package main

import (
	"log"
	"os"

	"github.com/mandolyte/TableRecursion"
)

func main() {
	atable := [][]string{
		{"A", "B"},
		{"A", "C"},
		{"B", "D"},
		{"C", "E"},
		{"D", "F"},
		{"D", "A"},
		{"E", "G"},
	}
	pathf, patherr := os.Create("tester1_paths.csv")
	if patherr != nil {
		log.Fatalf("Error os.Create():\n%v\n", patherr)
	}
	defer pathf.Close()

	partf, parterr := os.Create("tester1_parts.csv")
	if parterr != nil {
		log.Fatalf("Error os.Create():\n%v\n", parterr)
	}
	defer partf.Close()

	infof, infoerr := os.Create("tester1_info.csv")
	if infoerr != nil {
		log.Fatalf("Error os.Create():\n%v\n", infoerr)
	}
	defer infof.Close()

	tr := TableRecursion.NewTR(atable, "/", 10,
		infof, pathf, partf, true)
	tr.Execute("A")
}
