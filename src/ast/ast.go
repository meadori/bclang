// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ast

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
// 7.0 Definitions

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

// A top-level module that is a collection of all
// declarations and definitions in the program segment.
type Module struct {
	Decls []Decl
}
