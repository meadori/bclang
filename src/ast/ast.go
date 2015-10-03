// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ast

// ----------------------------------------------------------------------------
// 4.0 Primary expressions

type Expr interface {
}

// ----------------------------------------------------------------------------
// 4.1 Names

// A name is a sequence of characters used
// to declare variables and define functions.
type Name struct {
	Val string
}

// A list of names.
type NameList struct {
	Names []*Name
}

// ----------------------------------------------------------------------------
// 5.0 Compound Expressions

type ExprList struct {
	Exprs []Expr
}

type ConstExpr struct {
	Contant int
}

// ----------------------------------------------------------------------------
// 7.0 Definitions

type Def interface {
	def()
}

// A global or constant declaration.
type Decl interface {
	VarDecls() []*VarDecl
}

// A single declaration from a constant or global
// declaration.
type VarDecl struct {
	Name     string
	Constant int
}

// ----------------------------------------------------------------------------
// 7.3 Global Declarations

type GlobalDecl struct {
	Items []*VarDecl
}

func (g *GlobalDecl) VarDecls() []*VarDecl {
	return g.Items
}

// ----------------------------------------------------------------------------
// 7.4 Manifest Declarations

type ConstantDecl struct {
	Items []*VarDecl
}

func (c *ConstantDecl) VarDecls() []*VarDecl {
	return c.Items
}

// ----------------------------------------------------------------------------
// 7.5 Simple Definitions

type AndDef struct {
	Lhs Def
	Rhs Def
}

func (*AndDef) def() {}

type SimpleDef struct {
	Names *NameList
	Exprs *ExprList
}

func (*SimpleDef) def() {}

type VecDef struct {
	Name string
	Expr Expr
}

func (*VecDef) def() {}

// A top-level module that is a collection of all
// declarations and definitions in the program segment.
type Program struct {
	Decls []Decl
	Defs  []Def
}
