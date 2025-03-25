package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("File Organizer CLI tool")

	var dirPath string
	var err error

	if len(os.Args) > 1 {
		dirPath = filepath.Clean(os.Args[1])
	} else {
		dirPath, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting directory", err)
			os.Exit(1)

		}
	}
	fmt.Println("Directory Path:", dirPath)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading inside dir", err)
		os.Exit(1)
	}
	groupedDir := make(map[string][]os.DirEntry)

	fmt.Println("Files in directory:")
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("Error getting file info for %s: %v\\n", file.Name(), err)
			continue
		}
		fileType := "File"
		if file.IsDir() {
			fileType = "Directory"
		}

		fmt.Printf("%s, %s, %s, %d bytes\n", file.Name(), fileType, ext, fileInfo.Size())

		existingFiles := groupedDir[ext]
		updatedFiles := append(existingFiles, file)
		groupedDir[ext] = updatedFiles
	}

	for _, value := range groupedDir {
		fmt.Println(value)
	}

}
