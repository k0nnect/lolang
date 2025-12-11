package main

import (
	"compiler/internal/compiler"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: " + os.Args[0] + " <filename> <output>")
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("unable to read file")
	}

	data, err := compiler.Compile(string(file))
	if err != nil {
		log.Fatalln(err)
	}

	_ = os.WriteFile(os.Args[2], data, 0644)
}
