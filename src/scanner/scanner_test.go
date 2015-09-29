// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package scanner

import (
	"testing"
)

// Helper test functions.

func assertTokensEqual(t *testing.T, tok, etok *Token) {
	if tok.kind != etok.kind {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
	if tok.literal != etok.literal {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
}

func assertTokensEqualSource(t *testing.T, toks []*Token, str string) {
	var s Scanner
	s.Init([]byte(str))
	for _, etok := range toks {
		tok := s.Next()
		assertTokensEqual(t, tok, etok)
	}

}

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
		assertTokensEqual(t, tok, etok)
	}
}

var test_fact_str = `get "libhdr"
let START() = valof $(
        for I = 1 to 5 do
                writef("%N! = %I4*N", I, FACT(I))
        resultis 0
$)

and FACT(N) = N = 0 -> 1, N * FACT(N - 1)`

var test_fact_tokens = []*Token{
	NewToken(GET, "get"),
	NewToken(STRINGCONST, "\"libhdr\""),
	NewToken(LET, "let"),
	NewToken(NAME, "START"),
	NewToken(RBRA, "("),
	NewToken(RKET, ")"),
	NewToken(EQ, "="),
	NewToken(VALOF, "valof"),
	NewToken(SECTBRA, "$("),
	NewToken(FOR, "for"),
	NewToken(NAME, "I"),
	NewToken(EQ, "="),
	NewToken(NUMBER, "1"),
	NewToken(TO, "to"),
	NewToken(NUMBER, "5"),
	NewToken(DO, "do"),
	NewToken(NAME, "writef"),
	NewToken(RBRA, "("),
	NewToken(STRINGCONST, "\"%N! = %I4*N\""),
	NewToken(COMMA, ","),
	NewToken(NAME, "I"),
	NewToken(COMMA, ","),
	NewToken(NAME, "FACT"),
	NewToken(RBRA, "("),
	NewToken(NAME, "I"),
	NewToken(RKET, ")"),
	NewToken(RKET, ")"),
	NewToken(SEMICOLON, ";"),
	NewToken(RESULTIS, "resultis"),
	NewToken(NUMBER, "0"),
	NewToken(SECTKET, "$)"),
	NewToken(AND, "and"),
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
	NewToken(EOF, ""),
}

func TestMultiple(t *testing.T) {
	assertTokensEqualSource(t, test_fact_tokens, test_fact_str)
}

var test_hello_str = `get "libhdr"
let start() = valof
$( writes("Hello, World!*n")
   resultis 0
$)`

var test_hello_tokens = []*Token{
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
	assertTokensEqualSource(t, test_hello_tokens, test_hello_str)
}

var test_semi_str = `global $(
        COUNT: 200
        ALL: 201
$)`

var test_semi_tokens = []*Token{
	NewToken(GLOBAL, "global"),
	NewToken(SECTBRA, "$("),
	NewToken(NAME, "COUNT"),
	NewToken(COLON, ":"),
	NewToken(NUMBER, "200"),
	NewToken(SEMICOLON, ";"),
	NewToken(NAME, "ALL"),
	NewToken(COLON, ":"),
	NewToken(NUMBER, "201"),
	NewToken(SECTKET, "$)"),
	NewToken(EOF, ""),
}

func TestSemiInsertion(t *testing.T) {
	assertTokensEqualSource(t, test_semi_tokens, test_semi_str)
}

var test_do_str = `
let START() = valof $(
        for I = 1 to 5 for J = 1 to 5 do
                writef("%N * %N = %N", I, J, I * J)
        resultis 0
$)`

var test_do_tokens = []*Token{
	NewToken(LET, "let"),
	NewToken(NAME, "START"),
	NewToken(RBRA, "("),
	NewToken(RKET, ")"),
	NewToken(EQ, "="),
	NewToken(VALOF, "valof"),
	NewToken(SECTBRA, "$("),
	NewToken(FOR, "for"),
	NewToken(NAME, "I"),
	NewToken(EQ, "="),
	NewToken(NUMBER, "1"),
	NewToken(TO, "to"),
	NewToken(NUMBER, "5"),
	NewToken(DO, "do"), // This DO is inserted.
	NewToken(FOR, "for"),
	NewToken(NAME, "J"),
	NewToken(EQ, "="),
	NewToken(NUMBER, "1"),
	NewToken(TO, "to"),
	NewToken(NUMBER, "5"),
	NewToken(DO, "do"),
	NewToken(NAME, "writef"),
	NewToken(RBRA, "("),
	NewToken(STRINGCONST, "\"%N * %N = %N\""),
	NewToken(COMMA, ","),
	NewToken(NAME, "I"),
	NewToken(COMMA, ","),
	NewToken(NAME, "J"),
	NewToken(COMMA, ","),
	NewToken(NAME, "I"),
	NewToken(STAR, "*"),
	NewToken(NAME, "J"),
	NewToken(RKET, ")"),
	NewToken(SEMICOLON, ";"),
	NewToken(RESULTIS, "resultis"),
	NewToken(NUMBER, "0"),
	NewToken(SECTKET, "$)"),
	NewToken(EOF, ""),
}

func TestDoInsertion(t *testing.T) {
	assertTokensEqualSource(t, test_do_tokens, test_do_str)
}
