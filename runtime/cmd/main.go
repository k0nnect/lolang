package main

import (
	"log"
	"os"
	"runtime/internal/vm"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: " + os.Args[0] + " <filename>")
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("unable to read file")
	}

	err = vm.ExecuteProgram(file)
	if err != nil {
		log.Fatalln(err)
	}
}
