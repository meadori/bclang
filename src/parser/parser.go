// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"github.com/meadori/bcpl-go/src/ast"
	"github.com/meadori/bcpl-go/src/scanner"
	"github.com/meadori/bcpl-go/src/token"
	"strconv"
)

type Parser struct {
	scan scanner.Scanner // The scanner.
	tok  *token.Token    // The current token produced by the scanner.
}

func (p *Parser) error(msg string) {
	panic("error: " + msg)
}

func (p *Parser) match(kind token.TokenKind) bool {
	if kind == p.tok.Kind {
		p.tok = p.scan.Next()
		return true
	} else {
		p.error(fmt.Sprintf("expected '%s' found '%s'.", kind, p.tok))
		return false
	}
}

func (p *Parser) parseNameList() *ast.NameList {
	// namelist := <name> [',' <name> ]
	if p.tok.Kind == token.NAME {
		var namelist []*ast.Name
		namelist = append(namelist, &ast.Name{p.tok.Lit})
		p.scan.Next()
		for p.match(token.COMMA) {
			namelist = append(namelist, &ast.Name{p.tok.Lit})
		}
		return &ast.NameList{namelist}
	}
	return nil
}

func (p *Parser) parseVarDecl() *ast.VarDecl {
	name, constant := p.tok.Lit, 0
	p.match(token.NAME)

	switch p.tok.Kind {
	case token.EQ, token.COLON:
		p.match(p.tok.Kind)
		lit := p.tok.Lit
		p.match(token.NUMBER)
		constant, _ = strconv.Atoi(lit)
	default:
		p.error("expected '=' or ':'.")
	}

	return &ast.VarDecl{name, constant}
}

func (p *Parser) parseConstantDefinition() ast.Decl {
	// constdef := < manifest | global >
	//          $( <name> < '=' | ':' > <constant>
	//             [';' <name> <'=' | ':'> <constant>]* $)

	// We know we have a MANIFEST or GLOBAL.
	haveGlobal := p.tok.Kind == token.GLOBAL
	p.match(p.tok.Kind)

	// Build up the list of declarations.
	p.match(token.SECTBRA)
	var decls []*ast.VarDecl
	decls = append(decls, p.parseVarDecl())
	for p.tok.Kind == token.SEMICOLON {
		p.match(token.SEMICOLON)
		decls = append(decls, p.parseVarDecl())
	}
	p.match(token.SECTKET)

	// Build the declaration node.
	if haveGlobal {
		return &ast.GlobalDecl{decls}
	} else {
		return &ast.ConstantDecl{decls}
	}
}

func (p *Parser) parseDefinitions() *ast.Module {
	// definition := <D> | <constdef>
	var decls []ast.Decl
	for {
		switch p.tok.Kind {
		case token.MANIFEST, token.GLOBAL:
			decls = append(decls, p.parseConstantDefinition())
		case token.EOF:
			goto done
		default:
			p.error("expected start of definition.")
		}
	}
done:
	return &ast.Module{decls}
}

func (p *Parser) Parse() *ast.Module {
	// 2.2 Canonical Syntax
	return p.parseDefinitions()
}

func (p *Parser) Init(src []byte) {
	p.scan.Init(src)
	p.tok = p.scan.Next()
}
