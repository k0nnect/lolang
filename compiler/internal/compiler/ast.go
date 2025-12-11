package compiler

type Program struct {
	Functions []*Function `@@*`
}

type Function struct {
	Lo      string   `"lo"`
	RetType string   `[ @Type ]`
	Name    string   `@Ident`
	Params  []*Param `"(" ( @@ ( "," @@ )* )? ")"`
	Body    *Block   `@@`
}

type Param struct {
	Type string `@Type`
	Name string `@Ident`
}

type Block struct {
	Statements []*Statement `"{" @@* "}"`
}

type Statement struct {
	For *ForLoop `  @@
                         |`
	Var *VarDecl `@@ ";"
                         |`
	Return *ReturnStmt `@@ ";"
                         |`
	Expr *Expr `@@ ";"`
}

type ForLoop struct {
	For  string   `"for"`
	Init *ForInit `"(" @@ ";"`
	Cond *Expr    `@@ ";"`
	Post *Expr    `@@ ")"`
	Body *Block   `@@`
}

type ForInit struct {
	Var *VarDecl `  @@
                   |`
	Expr *Expr `@@`
}

type VarDecl struct {
	Type string `@Type`
	Name string `@Ident`
	Init *Expr  `[ "=" @@ ]`
}

type ReturnStmt struct {
	Return string `"return"`
	Expr   *Expr  `[ @@ ]`
}

type Expr struct {
	Assign *AssignExpr `  @@
                          |`
	Simple *SimpleExpr `@@`
}

type AssignExpr struct {
	Left  string      `@Ident`
	Op    string      `"="`
	Right *SimpleExpr `@@`
}

type SimpleExpr struct {
	Left  *AddExpr `@@`
	Op    *string  `[ @( "==" | "!=" | "<" | ">" | "<=" | ">=" )`
	Right *AddExpr `  @@ ]`
}

type AddExpr struct {
	Left *MulExpr `@@`
	Rest []struct {
		Op   string   `@( "+" | "-" )`
		Expr *MulExpr `@@`
	} `@@*`
}

type MulExpr struct {
	Left *Primary `@@`
	Rest []struct {
		Op   string   `@( "*" | "/" )`
		Expr *Primary `@@`
	} `@@*`
}

type Primary struct {
	Number *int `  @Number
                        |`
	String *string `@String
                        |`
	Ident *IdentExpr `@@
                        |`
	Sub *Expr `"(" @@ ")"`
}

type IdentExpr struct {
	Name string    `@Ident`
	Call *CallExpr `[ @@ ]`
}

type CallExpr struct {
	Args []*Expr `"(" ( @@ ( "," @@ )* )? ")"`
}
