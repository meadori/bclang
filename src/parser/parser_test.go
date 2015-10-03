// Copyright 2015 Meador Inge.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package parser

import (
	"testing"
)

var test_decl_str = `global $(
        COUNT: 200;
        ALL = 13
	fizz: 1234
$)

manifest $(
        N = 20000
        PI: 314;
	X = 1010
$)

global $( FOO = 42 $)
`

type test_decl struct {
	name  string
	value int
}

var test_decls = [][]test_decl{
	{
		test_decl{"COUNT", 200},
		test_decl{"ALL", 13},
		test_decl{"fizz", 1234},
	},
	{
		test_decl{"N", 20000},
		test_decl{"PI", 314},
		test_decl{"X", 1010},
	},
	{
		test_decl{"FOO", 42},
	},
}

func TestDeclarations(t *testing.T) {
	var p Parser
	p.Init([]byte(test_decl_str))
	m := p.Parse()

	if len(test_decls) != len(m.Decls) {
		t.Errorf("Expected %d declarations.", len(test_decls))
	}

	for i, decl_group := range test_decls {
		if len(decl_group) != len(m.Decls[i].VarDecls()) {
			t.Errorf("Expected %d declarations.", len(decl_group))
			continue
		}
		for j, edecl := range decl_group {
			decl := m.Decls[i].VarDecls()[j]
			if decl.Name != edecl.name {
				t.Errorf("Variable name does not equal '%s'.", edecl.name)
			}
			if decl.Constant != edecl.value {
				t.Errorf("Variable value does not equal '%d'.", edecl.value)
			}
		}
	}
}

var test_simple_def_str = `
let	X, Y, Z = 1, 2, 3
and	W, S = 4, 5
and	V = vec 5
`

func TestSimpleDefs(t *testing.T) {
	var p Parser
	p.Init([]byte(test_simple_def_str))
	p.Parse()
}
