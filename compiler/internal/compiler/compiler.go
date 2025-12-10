package compiler

import "log"

func Compile() {
	src := `
lo main(string name) {
    println("hello");
    return;
}
`

	prog, err := ParseString(src)
	if err != nil {
		log.Fatal("parse error:", err)
	}

	if err := CheckProgram(prog); err != nil {
		log.Fatal("type error:", err)
	}
}
