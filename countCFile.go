package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type countCFile struct {
	startRe, endRe *regexp.Regexp
}

func NewCountCFile() *countCFile {
	this := &countCFile{}
	this.startRe = regexp.MustCompile("/\\*")
	this.endRe = regexp.MustCompile("\\*/")
	return this
}

//
// return value:
//  allComment:  true: it's an all comment line
//              false: it's include real code
//  transition: -1: start comment
//               0: no change
//               1: end comment
func (this *countCFile) checkLine(line string) (allComment bool, transition int) {
	allComment = true
	transition = 0
	matched := false

	matched, _ = regexp.MatchString("^\\s*$", line)
	if matched {
		return
	}

	matched, _ = regexp.MatchString("^\\s*//.*$", line)
	if matched {
		return
	}

	matched, _ = regexp.MatchString("^\\s*[\\w})\"]\\w*", line)
	if matched {
		allComment = false
	}

	commentStarts := this.startRe.FindStringIndex(line)

	commentEnds := this.endRe.FindStringIndex(line)

	if len(commentStarts) > len(commentEnds) {
		transition = -1
	} else if len(commentStarts) < len(commentEnds) {
		transition = 1
	}
	return
}

func (this *countCFile) countFile(fileName *string) (line int) {
	line = 0
	inComment := false

	if *verbose > 3 {
		fmt.Printf("count: %s\n", *fileName)
	}
	if !strings.HasSuffix(*fileName, *extension) {
		return
	}

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allComment, transition := this.checkLine(scanner.Text())
		if *verbose > 4 {
			fmt.Printf("%v,%v: %s\n", allComment, transition, scanner.Text())
		}
		if inComment {
			if transition == 0 {
				continue
			} else if transition == 1 {
				inComment = false
			}
		} else {
			if transition == -1 {
				inComment = true
			}
		}
		if allComment != true {
			line++
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Println(*fileName, ": ", err)
	}
	if *verbose > 1 {
		fmt.Printf("%s:%d\n", *fileName, line)
	}
	return line
}
