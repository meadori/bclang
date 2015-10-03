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

func (p *Parser) parseSingleDecl() *ast.VarDecl {
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

func (p *Parser) parseDecl() ast.Decl {
	// constdef := < manifest | global >
	//          $( <name> < '=' | ':' > <constant>
	//             [';' <name> <'=' | ':'> <constant>]* $)

	// We know we have a MANIFEST or GLOBAL.
	haveGlobal := p.tok.Kind == token.GLOBAL
	p.match(p.tok.Kind)

	// Build up the list of declarations.
	p.match(token.SECTBRA)
	var decls []*ast.VarDecl
	decls = append(decls, p.parseSingleDecl())
	for p.tok.Kind == token.SEMICOLON {
		p.match(token.SEMICOLON)
		decls = append(decls, p.parseSingleDecl())
	}
	p.match(token.SECTKET)

	// Build the declaration node.
	if haveGlobal {
		return &ast.GlobalDecl{decls}
	} else {
		return &ast.ConstantDecl{decls}
	}
}

func (p *Parser) parseExpr() ast.Expr {
	lit := p.tok.Lit
	p.match(token.NUMBER)
	constant, _ := strconv.Atoi(lit)

	return &ast.ConstExpr{constant}
}

func (p *Parser) parseExprList() *ast.ExprList {
	var exprlist []ast.Expr
	exprlist = append(exprlist, p.parseExpr())

	for p.tok.Kind == token.COMMA {
		p.match(token.COMMA)
		exprlist = append(exprlist, p.parseExpr())
	}

	return &ast.ExprList{exprlist}
}

func (p *Parser) parseVarDef(name string) ast.Def {
	var namelist []*ast.Name
	namelist = append(namelist, &ast.Name{name})

	for p.tok.Kind == token.COMMA {
		p.match(token.COMMA)
		namelist = append(namelist, &ast.Name{p.tok.Lit})
		p.match(token.NAME)
	}

	p.match(token.EQ)

	if p.tok.Kind == token.VEC {
		p.match(token.VEC)
		return &ast.VecDef{name, p.parseExpr()}
	} else {
		exprlist := p.parseExprList()
		if len(namelist) != len(exprlist.Exprs) {
			p.error("assignment count mismatch")
		}
		return &ast.SimpleDef{&ast.NameList{namelist}, exprlist}
	}
}

func (p *Parser) parseSingleDef() (def ast.Def) {
	name := p.tok.Lit
	p.match(token.NAME)

	switch p.tok.Kind {
	case token.COMMA:
		def = p.parseVarDef(name)
	case token.EQ:
		def = p.parseVarDef(name)
	}

	return
}

func (p *Parser) parseSimulDef() (def ast.Def) {
	def = p.parseSingleDef()

	for p.tok.Kind == token.AND {
		p.match(token.AND)
		rhsDef := p.parseSingleDef()
		def = &ast.AndDef{def, rhsDef}
	}

	return
}

func (p *Parser) parseDef() ast.Def {
	p.match(token.LET)
	return p.parseSimulDef()
}

func (p *Parser) Parse() *ast.Program {
	var decls []ast.Decl
	var defs []ast.Def
	for {
		switch p.tok.Kind {
		case token.MANIFEST, token.GLOBAL:
			decls = append(decls, p.parseDecl())
		case token.LET:
			defs = append(defs, p.parseDef())
		case token.EOF:
			goto done
		default:
			p.error(fmt.Sprintf("expected definition found '%s'.", p.tok))
		}
	}
done:
	return &ast.Program{decls, defs}
}

func (p *Parser) Init(src []byte) {
	p.scan.Init(src)
	p.tok = p.scan.Next()
}
