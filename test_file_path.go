package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	wiki_files := make([]string, 3)
	files, _ := filepath.Glob("*.txt")
	for _, f := range files {
		//fmt.Println()
		wiki_files = append(wiki_files, f[:(len(f)-4)])
	}
	fmt.Println("The result of wiki_files")
	fmt.Println(wiki_files) // contains a list of all files in the current directory
}
