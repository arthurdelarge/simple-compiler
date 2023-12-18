package main

import "github.com/arthurdelarge/simple-compiler/pkg/analyzer/syntactic"

func main() {
	file := "sinput.txt"

	parser := syntactic.NewParser(file)
	defer parser.Close()

	parser.Parse()
}
