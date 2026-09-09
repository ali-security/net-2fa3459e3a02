// Command create_mod packages a module source tree into a Go module zip using
// golang.org/x/mod/zip.CreateFromDir, which applies the module proxy's exact
// inclusion rules (excluding .git/, vendor/ and nested modules).
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
