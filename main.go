package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the flag before parsing
	disableDir := flag.Bool("disableDir", false, "Disable checking for directories")

	// Parse flags first
	flag.Parse()

	// Determine the directory path
	var dirPath string
	var err error

	// Check if a directory path was provided as a non-flag argument
	if flag.NArg() > 0 {
		dirPath = filepath.Clean(flag.Arg(0))
	} else {
		// If no directory specified, use current working directory
		dirPath, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting directory:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Directory Path:", dirPath)
	fmt.Println("Disable Directories:", *disableDir)

	// Read directory contents
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	// Group files by extension
	groupedDir := make(map[string][]os.DirEntry)

	fmt.Println("Files in directory:")
	for _, file := range files {
		// Skip directories if disableDir is true
		if file.IsDir() && *disableDir {
			continue
		}

		// Get file extension
		ext := filepath.Ext(file.Name())
		if ext == "" {
			ext = "No Extension"
		}

		// Get file info
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("Error getting file info for %s: %v\n", file.Name(), err)
			continue
		}

		// Determine file type
		fileType := "File"
		if file.IsDir() {
			fileType = "Directory"
		}

		// Print file details
		fmt.Printf("%s, %s, %s, %d bytes\n", file.Name(), fileType, ext, fileInfo.Size())

		// Group files by extension
		groupedDir[ext] = append(groupedDir[ext], file)
	}

	// Print grouped files
	fmt.Println("\nGrouped Files:")
	for ext, groupedFiles := range groupedDir {
		fmt.Printf("Extension %s:\n", ext)
		for _, file := range groupedFiles {
			fmt.Println("  ", file.Name())
		}
	}
}
