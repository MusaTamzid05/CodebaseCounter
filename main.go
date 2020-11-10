package main

import (
	"codebase_info/codebaseLib"
)

func main() {

	summarizer := codebaseLib.NewSummarizer("/home/musa/git_clones/requests")
	summarizer.Run()
}
