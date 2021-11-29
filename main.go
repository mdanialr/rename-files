package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	var (
		tDir, prefix, regX, extension string
		isRename, noEnd, ok           bool
	)
	flag.StringVar(&tDir, "dir", "", "define where the files to be renamed is located.")
	flag.StringVar(&prefix, "prefix", "", "define the prefix of the renamed file.")
	flag.StringVar(&regX, "re", "", "define custom regex pattern.")
	flag.StringVar(&extension, "ext", "", "filter by extension.")
	flag.BoolVar(&isRename, "r", false, "rename the files.")
	flag.BoolVar(&noEnd, "no-end", false, "omit 'END' from last file's suffix.")
	flag.Parse()
	if tDir == "" {
		log.Fatalln("target dir is empty")
	}
	if !strings.HasSuffix(tDir, "/") {
		tDir += "/"
	}

	files, err := os.ReadDir(tDir)
	if err != nil {
		log.Fatalf("failed to read dir %v: %v\n", tDir, err)
	}

	for i, fl := range files {
		if fl.IsDir() {
			continue
		}
		// make sure filter and the file's extension no empty
		if extension != "" && path.Ext(fl.Name()) != "" {
			// if not match with the filtered extention then
			// just skip it
			if path.Ext(fl.Name()) != extension {
				continue
			}
		}

		pattern := evalRegex(regX)
		res := pattern.Find([]byte(fl.Name()))
		re := trimNumber(res, regX)
		newN := prefix + re

		// if this is the last file in dir, append 'END'
		// suffix by default.
		if i == len(files)-1 && !noEnd {
			newN += " END"
		}

		// append file extension
		if ext := path.Ext(fl.Name()); ext != "" {
			newN += ext
		}

		if isRename {
			if err := os.Rename(tDir+fl.Name(), tDir+newN); err != nil {
				log.Fatalf("failed to renaming file %v: %v\n", fl.Name(), err)
			}
			ok = true
			continue
		}

		fmt.Println("")
		fmt.Println("Old:", fl.Name())
		fmt.Println("New:", newN)
	}

	if ok {
		fmt.Println("all files renamed successfully")
	}
}

// evalRegex evaluate input string and return custom regex
// with the input as pattern otherwise use default pattern.
func evalRegex(v string) *regexp.Regexp {
	if v == "" {
		return regexp.MustCompile(`\d\d\d?\d?`)
	}
	return regexp.MustCompile(v)
}

// trimNumber trim founded number from blank space and any
// non number char.
func trimNumber(b []byte, regX string) string {
	re := strings.Trim(string(b), " ")
	pattern := evalRegex(regX)
	res := pattern.Find([]byte(re))

	return string(res)
}
