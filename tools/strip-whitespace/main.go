package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var Extensions = map[string]bool{
	".txt": true,

	".java":  true,
	".py":    true,
	".rb":    true,
	".js":    true,
	".m":     true,
	".proto": true,

	".j2":         true,
	".ini":        true,
	".conf":       true,
	".properties": true,

	".css":  true,
	".html": true,

	".md": true,

	".json": true,
	".xml":  true,
	".yml":  true,

	".sql": true,
}

var (
	Lines []string = make([]string, 0, 10000)
	DryF  bool
)

func processFile(path string, info os.FileInfo) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	Lines = Lines[:0]
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		Lines = append(Lines, strings.TrimRightFunc(scanner.Text(), unicode.IsSpace))
	}
	if err = scanner.Err(); err != nil {
		log.Print(err)
		// log.Fatal(err)
		return
	}

	f.Close()

	if DryF {
		return
	}

	f, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range Lines {
		_, err = f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Close()
}

func main() {
	flag.BoolVar(&DryF, "dry", false, "dry run")
	flag.Parse()

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if path == ".git" || path == ".env" || path == "_env" {
				return filepath.SkipDir
			}
			return nil
		}

		if Extensions[filepath.Ext(path)] {
			processFile(path, info)
			log.Print(path)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
