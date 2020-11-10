package main

import (
	"codebase_info/codebaseLib"
	"flag"
	"log"
)

func Usage() {
	log.Fatalln("Usage: -path path_to_codebase")
}

func main() {

	pathPtr := flag.String("path", "", "Path of codebase")
	flag.Parse()

	if *pathPtr == "" {
		Usage()
	}

	summarizer := codebaseLib.NewSummarizer(*pathPtr)
	summarizer.Run()
}
