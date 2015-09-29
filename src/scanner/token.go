// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This code was heavily inspired by the Go programming language's
// scanner code:
//
//   * https://github.com/golang/go/blob/master/src/go/token/token.go

package scanner

type TokenKind int

type Token struct {
	kind    TokenKind
	literal string
}

// 2.1.1 BCPL Canonical Symbols
const (
	ILLEGAL TokenKind = iota

	// End of file
	EOF

	// Comment
	COMMENT

	// Get
	GET

	// Literals
	NAME
	NUMBER
	STRINGCONST

	// Operators
	ASS
	COLON
	COMMA
	COND
	DIV
	EQ
	GE
	GR
	LE
	LOGAND
	LOGOR
	LS
	LSHIFT
	MINUS
	NE
	NOT
	PLUS
	RBRA
	RKET
	RSHIFT
	SBRA
	SECTBRA
	SECTKET
	SEMICOLON
	SKET
	STAR

	// Reserved "system" words.
	reserved_begin
	AND
	BE
	BREAK
	CASE
	DEFAULT
	DO
	EQV
	FALSE
	FINISH
	FOR
	GLOBAL
	GOTO
	IF
	INTO
	LET
	LV
	MANIFEST
	NEQV
	OR
	REM
	REPEAT
	REPEATUNTIL
	REPEATWHILE
	RESULTIS
	RETURN
	RV
	SWITCHON
	TEST
	TO
	TRUE
	UNLESS
	UNTIL
	VALOF
	VEC
	WHILE
	reserved_end
)

var restoks = [...]string{
	AND:         "and",
	BE:          "be",
	BREAK:       "break",
	CASE:        "case",
	DEFAULT:     "default",
	DO:          "do",
	EQV:         "eqv",
	FALSE:       "false",
	FINISH:      "finish",
	FOR:         "for",
	GLOBAL:      "global",
	GOTO:        "goto",
	IF:          "if",
	INTO:        "into",
	LET:         "let",
	LV:          "lv",
	MANIFEST:    "manifest",
	NEQV:        "neqv",
	OR:          "or",
	REM:         "rem",
	REPEAT:      "repeat",
	REPEATUNTIL: "repeatuntil",
	REPEATWHILE: "repeatwhile",
	RESULTIS:    "resultis",
	RETURN:      "return",
	RV:          "rv",
	SWITCHON:    "switchon",
	TEST:        "test",
	TO:          "to",
	TRUE:        "true",
	UNLESS:      "unless",
	UNTIL:       "until",
	VALOF:       "valof",
	VEC:         "vec",
	WHILE:       "while",
}

// Map from reserved system words to token kind.
var reswords map[string]TokenKind

func init() {
	reswords = make(map[string]TokenKind)
	for i := reserved_begin + 1; i < reserved_end; i++ {
		reswords[restoks[i]] = i
	}
}

// Lookup the given string and determine if it is a name or
// a reserved system word.
func LookupName(str string) TokenKind {
	if tok, is_reserved := reswords[str]; is_reserved {
		return tok
	}
	return NAME
}

// Create a new token.
func NewToken(kind TokenKind, literal string) *Token {
	t := new(Token)
	t.kind = kind
	t.literal = literal
	return t
}

// Return the string representation of the token.
func (tok Token) String() string {
	return tok.literal
}
