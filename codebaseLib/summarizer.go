package codebaseLib

import (
	"log"
	"strings"
)

type Summarizer struct {
	path       string
	fileExtMap map[string]int
}

func NewSummarizer(path string) *Summarizer {
	summarizer := Summarizer{path: path}
	summarizer.fileExtMap = make(map[string]int)
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
	}

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
