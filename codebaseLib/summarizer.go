package codebaseLib

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Summarizer struct {
	path           string
	fileExtMap     map[string]int
	lineInfo       map[string]int
	totalCodeCount int
}

func NewSummarizer(path string) *Summarizer {
	summarizer := Summarizer{path: path}
	summarizer.fileExtMap = make(map[string]int)
	summarizer.lineInfo = make(map[string]int)
	summarizer.totalCodeCount = 0
	return &summarizer
}

func (s *Summarizer) Run() {

	paths, err := GetFilePath(s.path)

	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range paths {
		data := strings.Split(path, "/")
		s.updateExts(data[len(data)-1])

		err = s.loadLineInfo(path)

		if err != nil {
			log.Println(err)
		}
	}

	s.showCodeLineInfo()
	s.showLanguage()
}

func (s *Summarizer) updateExts(name string) {

	if !strings.Contains(name, ".") {
		return
	}

	data := strings.Split(name, ".")
	ext := data[len(data)-1]

	if _, ok := s.fileExtMap[ext]; ok {
		s.fileExtMap[ext] = s.fileExtMap[ext] + 1
		return
	}

	s.fileExtMap[ext] = 0
}

func (s *Summarizer) showLanguage() {

	languages := map[string]string{"py": "python", "cpp": "c++"}

	for ext, count := range s.fileExtMap {
		for languageExt, languageName := range languages {
			if languageExt == ext {
				log.Printf("%s = %d files.\n", languageName, count)
			}
		}
	}

}

func (s *Summarizer) loadLineInfo(path string) error {

	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line == "" {
			continue
		}

		count = count + 1

	}

	s.lineInfo[path] = count
	s.totalCodeCount = s.totalCodeCount + count

	return nil

}

func (s *Summarizer) showCodeLineInfo() {

	for path, count := range s.lineInfo {
		log.Printf("%s = %d\n", path, count)
	}
	log.Printf("Total count : %d\n", s.totalCodeCount)
}
