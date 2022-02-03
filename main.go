package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type textInfo struct {
	filename, content                              string
	characters, words, lines                       []string
	wordCount, lineCount, sizeInBytes, sizeInChars int64
	charCount                                      map[string]int64
}

func NewTextInfo(filename string) (tf textInfo, err error) {
	content, err := getFileContent(filename)

	tf = textInfo{
		filename:    filename,
		content:     content,
		words:       getWords(content),
		wordCount:   getWordCount(content),
		lines:       getLines(content),
		lineCount:   getLineCount(content),
		characters:  getCharacters(content, true),
		charCount:   getCharacterCount(content, true),
		sizeInBytes: getSizeInBytes(content),
		sizeInChars: getSizeInChars(content),
	}

	return
}

func getSizeInBytes(str string) int64 {
	return int64(len([]byte(str)))
}

func getSizeInChars(str string) int64 {
	return int64(len([]rune(str)))
}

func getLines(str string) (lines []string) {
	return strings.Split(str, "\n")
}

func getLineCount(str string) int64 {
	return int64(len(getLines(str)))
}

func getCharacters(str string, ignoreLineBreak bool) (chars []string) {
	if ignoreLineBreak {
		str = strings.ReplaceAll(str, "\n", " ")
	}

	return strings.Split(str, "")
}

func getCharacterCount(str string, ignoreCasing bool) (r map[string]int64) {
	r = make(map[string]int64)

	for _, v := range getCharacters(str, true) {
		if ignoreCasing {
			v = strings.ToLower(v)
		}
		r[v] += 1
	}

	return
}

func getWords(str string) (list []string) {
	return strings.Fields(str)
}

func getWordCount(str string) int64 {
	return int64(len(getWords(str)))
}

func getFileName() (filename string, err error) {
	args := os.Args

	// Checks if only one argument was passed
	if len(args) < 2 {
		return "", errors.New("no filename was passed")
	} else if len(args) > 2 {
		return "", errors.New("too many args")
	}

	filename = args[1]

	return
}

func getFileContent(filename string) (content string, err error) {
	contentBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	content = string(contentBytes)

	return
}
func main() {
	filename, err := getFileName()

	if err != nil {
		log.Fatal(err)
	}

	info, err := NewTextInfo(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
