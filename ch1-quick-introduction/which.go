package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get arguments
	arguments := os.Args

	// Check if argument list is empty
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument :(")
	}

	// Get the file
	file := arguments[1]

	// Get all of the OS paths that have the executables
	path := os.Getenv("PATH")
	// Filepath library that splits the different executable paths
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit { // pathSplit can be separated out with the value and the index
		// For each direcatory, see if we can find the binary file
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode() // File modes are unix permission modes
			if mode.IsRegular() {   // Tests if it's a regular file and not a pointer or a directory
				if mode&0111 != 0 { // Check if the file permissions are executable --x--x--x
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
