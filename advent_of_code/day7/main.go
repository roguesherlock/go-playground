package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var filename = flag.String("input", "input.txt", "filename for the input")

type DirEntry struct {
	Name    string
	isDir   bool
	Size    int
	parent  *DirEntry
	entries []DirEntry
}

func (d DirEntry) String() string {
	return fmt.Sprintf("Name: %s\n Size: %d\n isDir: %v\n entries: [%v]\n\n", d.Name, d.Size, d.isDir, d.entries)
}

type walkDirFunc func(file DirEntry)

func calculateSize(entry DirEntry) int {
	size := 0

	if !entry.isDir {
		return entry.Size
	}

	for _, file := range entry.entries {
		size += calculateSize(file)
	}
	return size
}

func walkDir(entry DirEntry, fn walkDirFunc) {
	fn(entry)
	if entry.isDir {
		for _, file := range entry.entries {
			walkDir(file, fn)
		}
	}
}

func main() {
	flag.Parse()

	// Read the file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file when we're done
	defer file.Close()

	// commandRegex := regexp.MustCompile(`\$ (cd|ls) ([a-z0-9/]+)`)

	root := DirEntry{
		Name:    "/",
		isDir:   true,
		parent:  nil,
		entries: make([]DirEntry, 0),
	}
	currentDir := &root
	reader := bufio.NewReader(file)
	// we'll use this to determine if we're in an ls command
	inLs := false
scanner:
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		commands := strings.Split(string(line), " ")
		if commands != nil && commands[0] == "$" {
			switch commands[1] {
			case "cd":
				inLs = false
				if commands[2] == ".." {
					currentDir = currentDir.parent
				} else if commands[2] == "/" {
					currentDir = &root
				} else {
					for i, entry := range currentDir.entries {
						if entry.Name == commands[2] {
							// we use i here fetch the entry because the entry inside the for loop is the copy of entry in entries
							// and so if we try to use it's address, we'll be pointing at the address of the variable that holds that particular copy
							// and not the address of the actual entry.
							// i.e. entry here will point to the address of the entry variable in the for loop, which will always point to the last local copy of the entry
							currentDir = &currentDir.entries[i]
							continue scanner
						}
					}
				}
			case "ls":
				inLs = true
				continue
			}
		}

		// since there are only two commands, and cd doesn't return anything,
		// we can assume that if we're not in an ls command, we can ignore the line
		// otherwise we'll add the entries to the currentDir
		if !inLs {
			continue
		}
		result := strings.Split(string(line), " ")
		if strings.Contains(result[0], "dir") {
			currentDir.entries = append(currentDir.entries, DirEntry{
				Name:    result[1],
				isDir:   true,
				parent:  currentDir,
				entries: make([]DirEntry, 0),
			})
		} else {
			// we're in a file
			size, _ := strconv.Atoi(result[0])
			currentDir.entries = append(currentDir.entries, DirEntry{
				Name:   result[1],
				isDir:  false,
				Size:   size,
				parent: currentDir,
			})
		}
	}
	sum := 0
	walkDir(root, func(dir DirEntry) {
		if dir.isDir {
			size := calculateSize(dir)
			if size < 100000 {
				sum += size
			}
		}
	})
	// log.Println("Root", root.entries[0].entries)
	log.Println("Sum", sum)

	const (
		TotalAvailableDiskSpace = 70000000
		RequiredSpaceForUpdate  = 30000000
	)
	TotalRemainingSpace := TotalAvailableDiskSpace - calculateSize(root)

	var sizes []int
	walkDir(root, func(dir DirEntry) {
		if dir.isDir {
			sizes = append(sizes, calculateSize(dir))
		}
	})

	// sort the sizes
	sort.Ints(sizes)
	for _, size := range sizes {
		if (TotalRemainingSpace + size) >= RequiredSpaceForUpdate {
			log.Println("Size", size)
			break
		}

	}

}
