// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package scanner

import (
	"github.com/meadori/bcpl-go/src/token"
	"testing"
)

// Helper test functions.

func assertTokensEqual(t *testing.T, tok, etok *token.Token) {
	if tok.Kind != etok.Kind {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
	if tok.Lit != etok.Lit {
		t.Errorf("bad token: got '%s', expected '%s'", tok, etok)
	}
}

func assertTokensEqualSource(t *testing.T, toks []*token.Token, str string) {
	var s Scanner
	s.Init([]byte(str))
	for _, etok := range toks {
		tok := s.Next()
		assertTokensEqual(t, tok, etok)
	}

}

var test_single_token = [...]*token.Token{
	token.NewToken(token.EOF, ""),
	token.NewToken(token.COMMENT, "// this is a BCPL a comment!\n"),

	// Names.
	token.NewToken(token.NAME, "Global"),
	token.NewToken(token.NAME, "Let"),
	token.NewToken(token.NAME, "a"),
	token.NewToken(token.NAME, "Z"),

	// Numbers.
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.NUMBER, "98765"),

	// String constants.
	token.NewToken(token.STRINGCONST, "\"foo bar baz\n\""),

	// Operators.
	token.NewToken(token.VALOF, "valof"),
	token.NewToken(token.LV, "lv"),
	token.NewToken(token.RV, "rv"),
	token.NewToken(token.DIV, "/"),
	token.NewToken(token.REM, "rem"),
	token.NewToken(token.PLUS, "+"),
	token.NewToken(token.MINUS, "-"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NE, "!="),
	token.NewToken(token.LS, "<"),
	token.NewToken(token.GR, ">"),
	token.NewToken(token.LE, "<="),
	token.NewToken(token.GE, ">="),
	token.NewToken(token.NOT, "!"),
	token.NewToken(token.LSHIFT, "<<"),
	token.NewToken(token.RSHIFT, ">>"),
	token.NewToken(token.LOGAND, "&"),
	token.NewToken(token.LOGOR, "|"),
	token.NewToken(token.EQV, "eqv"),
	token.NewToken(token.NEQV, "neqv"),
	token.NewToken(token.COND, "->"),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.ASS, ":="),
	token.NewToken(token.SECTBRA, "$("),
	token.NewToken(token.SECTKET, "$)"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.SBRA, "["),
	token.NewToken(token.SKET, "]"),
	token.NewToken(token.COLON, ":"),
	token.NewToken(token.SEMICOLON, ";"),
	token.NewToken(token.STAR, "*"),

	// Keywords.
	token.NewToken(token.TRUE, "true"),
	token.NewToken(token.FALSE, "false"),
	token.NewToken(token.AND, "and"),
	token.NewToken(token.GOTO, "goto"),
	token.NewToken(token.RESULTIS, "resultis"),
	token.NewToken(token.TEST, "test"),
	token.NewToken(token.FOR, "for"),
	token.NewToken(token.IF, "if"),
	token.NewToken(token.UNLESS, "unless"),
	token.NewToken(token.WHILE, "while"),
	token.NewToken(token.UNTIL, "until"),
	token.NewToken(token.REPEAT, "repeat"),
	token.NewToken(token.REPEATWHILE, "repeatwhile"),
	token.NewToken(token.REPEATUNTIL, "repeatuntil"),
	token.NewToken(token.BREAK, "break"),
	token.NewToken(token.RETURN, "return"),
	token.NewToken(token.FINISH, "finish"),
	token.NewToken(token.SWITCHON, "switchon"),
	token.NewToken(token.CASE, "case"),
	token.NewToken(token.DEFAULT, "default"),
	token.NewToken(token.LET, "let"),
	token.NewToken(token.MANIFEST, "manifest"),
	token.NewToken(token.GLOBAL, "global"),
	token.NewToken(token.BE, "be"),
	token.NewToken(token.INTO, "into"),
	token.NewToken(token.TO, "to"),
	token.NewToken(token.DO, "do"),
	token.NewToken(token.OR, "or"),
	token.NewToken(token.VEC, "vec"),
}

func TestSingleToken(t *testing.T) {
	var s Scanner

	for _, etok := range test_single_token {
		s.Init([]byte(etok.Lit))
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

var test_fact_tokens = []*token.Token{
	token.NewToken(token.GET, "get"),
	token.NewToken(token.STRINGCONST, "\"libhdr\""),
	token.NewToken(token.LET, "let"),
	token.NewToken(token.NAME, "START"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.VALOF, "valof"),
	token.NewToken(token.SECTBRA, "$("),
	token.NewToken(token.FOR, "for"),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.TO, "to"),
	token.NewToken(token.NUMBER, "5"),
	token.NewToken(token.DO, "do"),
	token.NewToken(token.NAME, "writef"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.STRINGCONST, "\"%N! = %I4*N\""),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "FACT"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.SEMICOLON, ";"),
	token.NewToken(token.RESULTIS, "resultis"),
	token.NewToken(token.NUMBER, "0"),
	token.NewToken(token.SECTKET, "$)"),
	token.NewToken(token.AND, "and"),
	token.NewToken(token.NAME, "FACT"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.NAME, "N"),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NAME, "N"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NUMBER, "0"),
	token.NewToken(token.COND, "->"),
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "N"),
	token.NewToken(token.STAR, "*"),
	token.NewToken(token.NAME, "FACT"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.NAME, "N"),
	token.NewToken(token.MINUS, "-"),
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.EOF, ""),
}

func TestMultiple(t *testing.T) {
	assertTokensEqualSource(t, test_fact_tokens, test_fact_str)
}

var test_hello_str = `get "libhdr"
let start() = valof
$( writes("Hello, World!*n")
   resultis 0
$)`

var test_hello_tokens = []*token.Token{
	token.NewToken(token.GET, "get"),
	token.NewToken(token.STRINGCONST, "\"libhdr\""),
	token.NewToken(token.LET, "let"),
	token.NewToken(token.NAME, "start"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.VALOF, "valof"),
	token.NewToken(token.SECTBRA, "$("),
	token.NewToken(token.NAME, "writes"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.STRINGCONST, "\"Hello, World!*n\""),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.SEMICOLON, ";"),
	token.NewToken(token.RESULTIS, "resultis"),
	token.NewToken(token.NUMBER, "0"),
	token.NewToken(token.SECTKET, "$)"),
	token.NewToken(token.EOF, ""),
}

func TestHello(t *testing.T) {
	assertTokensEqualSource(t, test_hello_tokens, test_hello_str)
}

var test_semi_str = `global $(
        COUNT: 200
        ALL: 201
$)`

var test_semi_tokens = []*token.Token{
	token.NewToken(token.GLOBAL, "global"),
	token.NewToken(token.SECTBRA, "$("),
	token.NewToken(token.NAME, "COUNT"),
	token.NewToken(token.COLON, ":"),
	token.NewToken(token.NUMBER, "200"),
	token.NewToken(token.SEMICOLON, ";"),
	token.NewToken(token.NAME, "ALL"),
	token.NewToken(token.COLON, ":"),
	token.NewToken(token.NUMBER, "201"),
	token.NewToken(token.SECTKET, "$)"),
	token.NewToken(token.EOF, ""),
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

var test_do_tokens = []*token.Token{
	token.NewToken(token.LET, "let"),
	token.NewToken(token.NAME, "START"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.VALOF, "valof"),
	token.NewToken(token.SECTBRA, "$("),
	token.NewToken(token.FOR, "for"),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.TO, "to"),
	token.NewToken(token.NUMBER, "5"),
	token.NewToken(token.DO, "do"), // This DO is inserted.
	token.NewToken(token.FOR, "for"),
	token.NewToken(token.NAME, "J"),
	token.NewToken(token.EQ, "="),
	token.NewToken(token.NUMBER, "1"),
	token.NewToken(token.TO, "to"),
	token.NewToken(token.NUMBER, "5"),
	token.NewToken(token.DO, "do"),
	token.NewToken(token.NAME, "writef"),
	token.NewToken(token.RBRA, "("),
	token.NewToken(token.STRINGCONST, "\"%N * %N = %N\""),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "J"),
	token.NewToken(token.COMMA, ","),
	token.NewToken(token.NAME, "I"),
	token.NewToken(token.STAR, "*"),
	token.NewToken(token.NAME, "J"),
	token.NewToken(token.RKET, ")"),
	token.NewToken(token.SEMICOLON, ";"),
	token.NewToken(token.RESULTIS, "resultis"),
	token.NewToken(token.NUMBER, "0"),
	token.NewToken(token.SECTKET, "$)"),
	token.NewToken(token.EOF, ""),
}

func TestDoInsertion(t *testing.T) {
	assertTokensEqualSource(t, test_do_tokens, test_do_str)
}
