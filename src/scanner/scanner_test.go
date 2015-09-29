// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package scanner

import (
	"testing"
)

type test struct {
	tok TokenKind
	str string
}

var test_single_token = [...]*Token{
	NewToken(EOF, ""),
	NewToken(COMMENT, "// this is a BCPL a comment!\n"),

	// Names.
	NewToken(NAME, "Global"),
	NewToken(NAME, "Let"),
	NewToken(NAME, "a"),
	NewToken(NAME, "Z"),

	// Numbers.
	NewToken(NUMBER, "1"),
	NewToken(NUMBER, "98765"),

	// String constants.
	NewToken(STRINGCONST, "\"foo bar baz\n\""),

	// Operators.
	NewToken(VALOF, "valof"),
	NewToken(LV, "lv"),
	NewToken(RV, "rv"),
	NewToken(DIV, "/"),
	NewToken(REM, "rem"),
	NewToken(PLUS, "+"),
	NewToken(MINUS, "-"),
	NewToken(EQ, "="),
	NewToken(NE, "!="),
	NewToken(LS, "<"),
	NewToken(GR, ">"),
	NewToken(LE, "<="),
	NewToken(GE, ">="),
	NewToken(NOT, "!"),
	NewToken(LSHIFT, "<<"),
	NewToken(RSHIFT, ">>"),
	NewToken(LOGAND, "&"),
	NewToken(LOGOR, "|"),
	NewToken(EQV, "eqv"),
	NewToken(NEQV, "neqv"),
	NewToken(COND, "->"),
	NewToken(COMMA, ","),
	NewToken(ASS, ":="),
	NewToken(SECTBRA, "$("),
	NewToken(SECTKET, "$)"),
	NewToken(RBRA, "("),
	NewToken(RKET, ")"),
	NewToken(SBRA, "["),
	NewToken(SKET, "]"),
	NewToken(COLON, ":"),
	NewToken(SEMICOLON, ";"),
	NewToken(STAR, "*"),

	// Keywords.
	NewToken(TRUE, "true"),
	NewToken(FALSE, "false"),
	NewToken(AND, "and"),
	NewToken(GOTO, "goto"),
	NewToken(RESULTIS, "resultis"),
	NewToken(TEST, "test"),
	NewToken(FOR, "for"),
	NewToken(IF, "if"),
	NewToken(UNLESS, "unless"),
	NewToken(WHILE, "while"),
	NewToken(UNTIL, "until"),
	NewToken(REPEAT, "repeat"),
	NewToken(REPEATWHILE, "repeatwhile"),
	NewToken(REPEATUNTIL, "repeatuntil"),
	NewToken(BREAK, "break"),
	NewToken(RETURN, "return"),
	NewToken(FINISH, "finish"),
	NewToken(SWITCHON, "switchon"),
	NewToken(CASE, "case"),
	NewToken(DEFAULT, "default"),
	NewToken(LET, "let"),
	NewToken(MANIFEST, "manifest"),
	NewToken(GLOBAL, "global"),
	NewToken(BE, "be"),
	NewToken(INTO, "into"),
	NewToken(TO, "to"),
	NewToken(DO, "do"),
	NewToken(OR, "or"),
	NewToken(VEC, "vec"),
}

func TestSingleToken(t *testing.T) {
	var s Scanner

	for _, etok := range test_single_token {
		s.Init([]byte(etok.literal))

		tok := s.Next()
		if tok.kind != etok.kind {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
		if tok.literal != etok.literal {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
	}
}

var test_fact_str = "FACT(N) = N = 0 -> 1, N * FACT(N - 1)"
var test_fact_tokens = [...]*Token{
	NewToken(NAME, "FACT"),
	NewToken(RBRA, "("),
	NewToken(NAME, "N"),
	NewToken(RKET, ")"),
	NewToken(EQ, "="),
	NewToken(NAME, "N"),
	NewToken(EQ, "="),
	NewToken(NUMBER, "0"),
	NewToken(COND, "->"),
	NewToken(NUMBER, "1"),
	NewToken(COMMA, ","),
	NewToken(NAME, "N"),
	NewToken(STAR, "*"),
	NewToken(NAME, "FACT"),
	NewToken(RBRA, "("),
	NewToken(NAME, "N"),
	NewToken(MINUS, "-"),
	NewToken(NUMBER, "1"),
	NewToken(RKET, ")"),
}

func TestMultiple(t *testing.T) {
	var s Scanner
	s.Init([]byte(test_fact_str))

	for _, etok := range test_fact_tokens {
		tok := s.Next()
		if tok.kind != etok.kind {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
		if tok.literal != etok.literal {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
	}
}

var test_hello_str = `get "libhdr"
let start() = valof
$( writes("Hello, World!*n")
   resultis 0
$)`

var test_hello_tokens = [...]*Token{
	NewToken(GET, "get"),
	NewToken(STRINGCONST, "\"libhdr\""),
	NewToken(LET, "let"),
	NewToken(NAME, "start"),
	NewToken(RBRA, "("),
	NewToken(RKET, ")"),
	NewToken(EQ, "="),
	NewToken(VALOF, "valof"),
	NewToken(SECTBRA, "$("),
	NewToken(NAME, "writes"),
	NewToken(RBRA, "("),
	NewToken(STRINGCONST, "\"Hello, World!*n\""),
	NewToken(RKET, ")"),
	NewToken(SEMICOLON, ";"),
	NewToken(RESULTIS, "resultis"),
	NewToken(NUMBER, "0"),
	NewToken(SECTKET, "$)"),
	NewToken(EOF, ""),
}

func TestHello(t *testing.T) {
	var s Scanner
	s.Init([]byte(test_hello_str))

	for _, etok := range test_hello_tokens {
		tok := s.Next()
		if tok.kind != etok.kind {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
		if tok.literal != etok.literal {
			t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
		}
	}
}