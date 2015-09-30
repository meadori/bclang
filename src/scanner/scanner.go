// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This code was heavily inspired by the Go programming language's
// scanner code:
//
//   * https://github.com/golang/go/blob/master/src/go/scanner/scanner.go

package scanner

import "github.com/meadori/bcpl-go/src/token"

// Scanner states.
const (
	normal      = iota // The normal state.
	maybeinsert        // A token might be inserted automatically.
	maybesemi          // A semicolon might be inserted.
)

type Scanner struct {
	src      []byte       // The source code.
	ch       rune         // The current character.
	chOffset int          // The current character offset.
	offset   int          // The next character offset.
	savedTok *token.Token // A saved token from an earlier scan.
	state    int          // In semicolon insertion state.
}

func (s *Scanner) next() {
	if s.offset < len(s.src) {
		s.chOffset = s.offset
		s.ch = rune(s.src[s.offset])
		s.offset += 1
	} else {
		s.chOffset = s.offset
		s.offset = len(s.src)
		s.ch = -1
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' && s.state != maybeinsert || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (s *Scanner) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isSemiStart(tok *token.Token) bool {
	// 2.1.2 Hardware Conventions and Preprocessor Rules
	// (d)

	switch tok.Kind {
	case token.TEST, token.FOR, token.IF, token.UNLESS, token.UNTIL,
		token.WHILE, token.GOTO, token.RESULTIS, token.CASE,
		token.DEFAULT, token.BREAK, token.RETURN, token.FINISH,
		token.SECTBRA, token.RBRA, token.VALOF, token.LV, token.RV,
		token.NAME:
		return true
	}
	return false
}

func isDoStart(tok *token.Token) bool {
	// 2.1.2 Hardware Conventions and Preprocessor Rules
	// (e)

	switch tok.Kind {
	case token.TEST, token.FOR, token.IF, token.UNLESS, token.UNTIL,
		token.WHILE, token.GOTO, token.RESULTIS, token.CASE,
		token.DEFAULT, token.BREAK, token.RETURN, token.FINISH:
		return true
	}
	return false
}

func isCommandEnd(tok *token.Token) bool {
	// 2.1.2 Hardware Conventions and Preprocessor Rules
	// (d)
	// (e)

	switch tok.Kind {
	case token.BREAK, token.RETURN, token.FINISH, token.REPEAT,
		token.SKET, token.RKET, token.SECTKET, token.NAME,
		token.STRINGCONST, token.NUMBER, token.TRUE, token.FALSE:
		return true
	}
	return false
}

func (s *Scanner) scanComment() string {
	// 2.1.2 Hardware Conventions and Preprocessor Rules
	// (b)

	start := s.chOffset - 2

	for s.ch != '\n' {
		s.next()
	}

	return string(s.src[start:s.offset])
}

func (s *Scanner) scanName() *token.Token {
	// 2.1.2 Hardware Conventions and Preprocessor Rules
	// (a)

	// (1) A name is either a single small letter or a sequence of letters
	// and digits starting with a capital letter. The character immediately
	// following a name may not be a letter or a digit.
	start := s.chOffset
	for s.isLetter(s.ch) || s.isDigit(s.ch) {
		s.next()
	}
	str := s.src[start:s.chOffset]

	// (2) A sequence of two or more small letters which is not part of a NAME,
	// SECTBRA, SECTKET or STRINGCONST is a reserved system word and may be used
	// to represent a canonical symbol.
	kind := token.NAME
	literal := string(str)
	if len(str) > 1 {
		kind = token.LookupName(literal)
	}
	return token.NewToken(kind, literal)
}

func (s *Scanner) scanNumber() *token.Token {
	start := s.chOffset
	for s.isDigit(s.ch) {
		s.next()
	}
	return token.NewToken(token.NUMBER, string(s.src[start:s.chOffset]))
}

func (s *Scanner) scanStringConst() *token.Token {
	start := s.offset - 1
	s.next()
	for s.ch != '"' {
		s.next()
	}
	s.next()
	return token.NewToken(token.STRINGCONST, string(s.src[start:s.chOffset]))
}

func (s *Scanner) scanOperator(ch rune) *token.Token {
	kind := token.ILLEGAL
	lit := string(ch)
	s.next()

	switch ch {
	case '/':
		if s.ch == '/' {
			s.next()
			kind, lit = token.COMMENT, s.scanComment()
		} else {
			kind = token.DIV
		}
	case '+':
		kind, lit = token.PLUS, "+"
	case '-':
		if s.ch == '>' {
			s.next()
			kind, lit = token.COND, "->"
		} else {
			kind = token.MINUS
		}
	case '=':
		kind = token.EQ
	case '!':
		if s.ch == '=' {
			s.next()
			kind, lit = token.NE, "!="
		} else {
			kind = token.NOT
		}
	case '<':
		switch s.ch {
		case '=':
			s.next()
			kind, lit = token.LE, "<="
		case '<':
			s.next()
			kind, lit = token.LSHIFT, "<<"
		default:
			kind = token.LS
		}
	case '>':
		switch s.ch {
		case '=':
			s.next()
			kind, lit = token.GE, ">="
		case '>':
			s.next()
			kind, lit = token.RSHIFT, ">>"
		default:
			kind = token.GR
		}
	case '&':
		kind = token.LOGAND
	case '|':
		kind = token.LOGOR
	case ',':
		kind = token.COMMA
	case ':':
		if s.ch == '=' {
			kind, lit = token.ASS, ":="
		} else {
			kind = token.COLON
		}
	case '$':
		switch s.ch {
		case '(':
			s.next()
			kind, lit = token.SECTBRA, "$("
		case ')':
			s.next()
			kind, lit = token.SECTKET, "$)"
		}
	case '(':
		kind = token.RBRA
	case ')':
		kind = token.RKET
	case '[':
		kind = token.SBRA
	case ']':
		kind = token.SKET
	case ';':
		kind = token.SEMICOLON
	case '*':
		kind = token.STAR
	case -1:
		kind, lit = token.EOF, ""
	}

	return token.NewToken(kind, lit)
}

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.savedTok = nil
	s.state = normal
	s.next()
}

func (s *Scanner) Next() (tok *token.Token) {
next:
	if s.savedTok != nil {
		tok = s.savedTok
		s.savedTok = nil
	} else {
		s.skipWhitespace()

		switch ch := s.ch; {
		case s.isLetter(ch):
			tok = s.scanName()
		case s.isDigit(ch):
			tok = s.scanNumber()
		case ch == '"':
			tok = s.scanStringConst()
		case ch == '\n':
			s.state = maybesemi
			goto next
		default:
			tok = s.scanOperator(ch)
		}

		switch s.state {
		case maybeinsert:
			if isDoStart(tok) {
				s.savedTok = tok
				tok = token.NewToken(token.DO, "do")
				s.state = normal
			} else if !isCommandEnd(tok) {
				s.state = normal
			}
		case maybesemi:
			if isSemiStart(tok) {
				s.savedTok = tok
				tok = token.NewToken(token.SEMICOLON, ";")
			}
			s.state = normal
		case normal:
			if isCommandEnd(tok) {
				s.state = maybeinsert
			}
		}
	}

	return
}
