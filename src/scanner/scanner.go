// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This code was heavily inspired by the Go programming language's
// scanner code:
//
//   * https://github.com/golang/go/blob/master/src/go/scanner/scanner.go

package scanner

type Scanner struct {
	src      []byte // The source code.
	ch       rune   // The current character.
	chOffset int    // The current character offset.
	offset   int    // The next character offset.
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
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (s *Scanner) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
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

func (s *Scanner) scanName() *Token {
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
	kind := NAME
	literal := string(str)
	if len(str) > 1 {
		kind = LookupName(literal)
	}
	return NewToken(kind, literal)
}

func (s *Scanner) scanNumber() *Token {
	start := s.chOffset
	for s.isDigit(s.ch) {
		s.next()
	}
	return NewToken(NUMBER, string(s.src[start:s.chOffset]))
}

func (s *Scanner) scanStringConst() *Token {
	start := s.offset - 1
	s.next()
	for s.ch != '"' {
		s.next()
	}
	return NewToken(STRINGCONST, string(s.src[start:s.offset]))
}

func (s *Scanner) scanOperator(ch rune) *Token {
	kind := ILLEGAL
	lit := string(ch)
	s.next()

	switch ch {
	case '/':
		if s.ch == '/' {
			s.next()
			kind, lit = COMMENT, s.scanComment()
		} else {
			kind = DIV
		}
	case '+':
		kind, lit = PLUS, "+"
	case '-':
		if s.ch == '>' {
			s.next()
			kind, lit = COND, "->"
		} else {
			kind = MINUS
		}
	case '=':
		kind = EQ
	case '!':
		if s.ch == '=' {
			s.next()
			kind, lit = NE, "!="
		} else {
			kind = NOT
		}
	case '<':
		switch s.ch {
		case '=':
			s.next()
			kind, lit = LE, "<="
		case '<':
			s.next()
			kind, lit = LSHIFT, "<<"
		default:
			kind = LS
		}
	case '>':
		switch s.ch {
		case '=':
			s.next()
			kind, lit = GE, ">="
		case '>':
			s.next()
			kind, lit = RSHIFT, ">>"
		default:
			kind = GR
		}
	case '&':
		kind = LOGAND
	case '|':
		kind = LOGOR
	case ',':
		kind = COMMA
	case ':':
		if s.ch == '=' {
			kind, lit = ASS, ":="
		} else {
			kind = COLON
		}
	case '$':
		switch s.ch {
		case '(':
			s.next()
			kind, lit = SECTBRA, "$("
		case ')':
			s.next()
			kind, lit = SECTKET, "$)"
		}
	case '(':
		kind = RBRA
	case ')':
		kind = RKET
	case '[':
		kind = SBRA
	case ']':
		kind = SKET
	case ';':
		kind = SEMICOLON
	case '*':
		kind = STAR
	case -1:
		kind, lit = EOF, ""
	}

	return NewToken(kind, lit)
}

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.offset = 0
	s.next()
}

func (s *Scanner) Next() (tok *Token) {
	s.skipWhitespace()

	switch ch := s.ch; {
	case s.isLetter(ch):
		tok = s.scanName()
	case s.isDigit(ch):
		tok = s.scanNumber()
	case ch == '"':
		tok = s.scanStringConst()
	default:
		tok = s.scanOperator(ch)
	}

	return
}
