package main

import "fmt"
import "flag"
import "os"
import "strings"
import "io/ioutil"
import "path/filepath"

func main() {
	var fVersion = flag.Bool("version", false, "Print version and exit successfully")
	var fAll = flag.Bool("all", false, "Print all instances of the found program")
	var fPrefix = flag.Bool("prefix", false, "find by a prefix match")

	flag.Parse()
	if *fVersion {
		fmt.Print("Which for Windows, Version 0.0.1")
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Expect something to search for")
		os.Exit(-1)
	}

	var seen = make(map[string]int)

	var path = os.Getenv("PATH")
	var searchSlice = strings.Split(path, ";")
	searchSlice = append([]string{"."}, searchSlice...)

	for _, pathpart := range searchSlice {

		files, err := ioutil.ReadDir(pathpart)
		if err != nil {
			continue
		}

		for _, f := range files {
			if isMatch(f, flag.Args()[0], *fPrefix) &&
				isExecutable(f) {
				var fullpath = filepath.Join(pathpart, f.Name())
				_, ok := seen[fullpath]
				if !ok {
					seen[fullpath] = 1
					fmt.Println(fmtPath(fullpath))
					if !*fAll {
						os.Exit(0)
					}
				}
			}
		}
	}

	if !*fAll {
		fmt.Println("Which: cannot find " + flag.Args()[0])
		fmt.Println("searched path: " + path)
	}
}

func isExecutable(f os.FileInfo) bool {
	return filepath.Ext(f.Name()) == ".exe" ||
		filepath.Ext(f.Name()) == ".bat"
}

func isMatch(f os.FileInfo, arg string, matchPrefix bool) bool {
	if matchPrefix {
		return strings.HasPrefix(f.Name(), arg)
	} else {
		var fName = strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		return fName == arg || f.Name() == arg
	}
}

func fmtPath(f string) string {
	f = filepath.Clean(f)
	if strings.Contains(f, " ") {
		f = "\"" + f + "\""
	}

	return f
}
