package compiler

import "log"

func Compile() {
	src := `
lo main() {
    println("hello");
    return;
}

lo int bruh() {
	return 5;
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
