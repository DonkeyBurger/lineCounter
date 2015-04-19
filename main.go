package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var verbose *int
var extension *string
var cFileCounter *countCFile

func main() {
	extension = flag.String("ext", "go", "-ext <extension>")
	verbose = flag.Int("verbose", 0, "-verbose <verbose level>")
	flag.Parse()
	sourceRoot := flag.Args()

	if len(sourceRoot) == 0 {
		sourceRoot = append(sourceRoot, ".")
	}
	cFileCounter = NewCountCFile()
	for _, val := range sourceRoot {
		line := 0

		fi, err := os.Stat(val)
		if err != nil {
			log.Fatal(err)
			return
		}
		if fi.IsDir() {
			line = countDir(&val)
		} else {
			line = getFileCounter().countFile(&val)
		}
		fmt.Printf("%s: %d\n", val, line)
	}
}

func getFileCounter() (ci CountAFile) {
	return cFileCounter
}

type CountAFile interface {
	countFile(fileName *string) (line int)
}

func countDir(rootPath *string) (line int) {
	file, err := os.Open(*rootPath)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer file.Close()

	fis, err := file.Readdir(0)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	for _, fi := range fis {
		filePath := path.Join(*rootPath, fi.Name())
		if fi.IsDir() {
			line += countDir(&filePath)
		} else {
			line += getFileCounter().countFile(&filePath)
		}
	}
	if *verbose > 0 {
		fmt.Printf("%s: %d\n", *rootPath, line)
	}
	return line
}
