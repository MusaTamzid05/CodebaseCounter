package codebaseLib

import "log"

type Summarizer struct {
	path string
}

func NewSummarizer(path string) *Summarizer {
	summarizer := Summarizer{path: path}
	return &summarizer
}

func (s *Summarizer) Run() {

	paths, err := GetFilePath(s.path)

	if err != nil {
		log.Fatalln(err)
	}

	for _, path := range paths {
		log.Println(path)
	}
}
