package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Must have an argument")
	}

	fileGlobs, err := filepath.Glob(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, filename := range fileGlobs {
		metadata, err := getMetaDataFromFile(filename)
		if err != nil {
			continue
		}

		trackNo, _ := metadata.Track()
		name := fmt.Sprintf("%02d - %s - %s - %s",
			trackNo,
			removeControlCharacters(metadata.Artist()),
			removeControlCharacters(metadata.Album()),
			removeControlCharacters(metadata.Title()))

		fileStat := getPathStat(filename)

		os.Rename(filename, fileStat.Basedir+"/"+name+fileStat.Extension)
	}
}

func getMetaDataFromFile(filename string) (tag.Metadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tag, err := tag.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func removeControlCharacters(str string) string {
	crtChars := []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"}
	for _, crtChar := range crtChars {
		str = strings.ReplaceAll(str, crtChar, "")
	}
	return str
}

type pathStat struct {
	Basedir   string
	Extension string
}

func getPathStat(filename string) pathStat {
	return pathStat{
		Basedir:   filepath.Dir(filename),
		Extension: filepath.Ext(filename),
	}
}
