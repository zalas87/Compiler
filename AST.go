package main

import (
	"fmt"
	"strings"
)

var debugAST bool

//IMPRESIÓN DEL ÁRBOL---------------------------------------------------------------------
func (t *AST) StringAST() string {
	s := "AST\n"
	lvl := 0
	s += t.StringMacro(lvl)
	return s
}

func (t *AST) StringMacro(lvl int) string {
	s := ""
	lvl++
	tab := strings.Repeat("\t", lvl)
	for i := range t.macro {
		if debugType {
			s += fmt.Sprintf("%s%s : '%s' (%s)\n", tab, Skindname[t.macro[i].id.kind], t.macro[i].id.name, t.macro[i].id.typ.name)
		} else {
			s += fmt.Sprintf("%s%s : '%s'\n", tab, Skindname[t.macro[i].id.kind], t.macro[i].id.name)
		}
		s += t.macro[i].StringDparam(lvl) + t.macro[i].body.StringBody(lvl)
	}
	return s
}

func (m *Macro) StringDparam(lvl int) string {
	s := ""
	tab := strings.Repeat("\t", lvl)

	lvl++
	tab = strings.Repeat("\t", lvl)
	for i := range m.dparam {
		if debugType {
			s += fmt.Sprintf("%s%s: '%s' (%s)\n", tab, Skindname[m.dparam[i].tipo.kind], m.dparam[i].tipo.name, m.dparam[i].tipo.typ.name)
			s += fmt.Sprintf("%s%s: '%s'<Tipo: %s(%s)>\n", tab, Skindname[m.dparam[i].id.kind], m.dparam[i].id.name, m.dparam[i].id.tnd.name, m.dparam[i].id.typ.name)
		} else {
			s += fmt.Sprintf("%s%s: '%s'\n", tab, Skindname[m.dparam[i].tipo.kind], m.dparam[i].tipo.name)
			s += fmt.Sprintf("%s%s: '%s'<Tipo: %s>\n", tab, Skindname[m.dparam[i].id.kind], m.dparam[i].id.name, m.dparam[i].id.tnd.name)
		}
	}
	return s
}

func (m *Body) StringBody(lvl int) string {
	tab := strings.Repeat("\t", lvl)
	s := tab + "{\n" + m.StringVars(lvl) + m.StringStmt(lvl) + tab + "}\n"
	return s
}

func (bd *Body) StringVars(lvl int) string {
	s := ""
	lvl++
	tab := strings.Repeat("\t", lvl)
	for j := range bd.vars {
		if debugType {
			s += fmt.Sprintf("%s%s: '%s' (%s)\n", tab, Skindname[bd.vars[j].tipo.kind], bd.vars[j].tipo.name, bd.vars[j].tipo.typ.name)
			s += fmt.Sprintf("%s%s: '%s'<Tipo: %s(%s)>\n", tab, Skindname[bd.vars[j].id.kind], bd.vars[j].id.name, bd.vars[j].id.tnd.name, bd.vars[j].id.typ.name)
		} else {
			s += fmt.Sprintf("%s%s: '%s'\n", tab, Skindname[bd.vars[j].tipo.kind], bd.vars[j].tipo.name)
			s += fmt.Sprintf("%s%s: '%s'<Tipo: %s>\n", tab, Skindname[bd.vars[j].id.kind], bd.vars[j].id.name, bd.vars[j].id.tnd.name)
		}
	}
	return s
}

func (bd *Body) StringStmt(lvl int) string {
	s := ""

	for i := range bd.stmt {
		switch bd.stmt[i].op {
		case "CALL":
			s += bd.stmt[i].StringCallStmt(lvl)
		case "ASSIGN":
			s += bd.stmt[i].StringAssigmentStmt(lvl)
		case "LOOP":
			s += bd.stmt[i].StringLoopStmt(lvl)
		}
	}
	return s
}

func (bl *Stmt) StringCallStmt(lvl int) string {
	lvl++
	tab := strings.Repeat("\t", lvl)
	s := fmt.Sprintf("%sCALL -> %s : '%s'\n", tab, Skindname[bl.callstmt.id.kind], bl.callstmt.id.name)
	s += bl.callstmt.StringParam(lvl)
	return s
}

func (cl *CallStmt) StringParam(lvl int) string {
	tab := strings.Repeat("\t", lvl)
	s := ""
	lvl++
	tab = strings.Repeat("\t", lvl)

	for i := range cl.param {
		s += tab + cl.param[i].StringNd() + "\n"
	}
	return s
}

func (bl *Stmt) StringLoopStmt(lvl int) string {
	s := ""
	lvl++
	tab := strings.Repeat("\t", lvl)
	sloop := tab + "loop"

	if debugType {
		s += fmt.Sprintf("%s [%s: '%s'<Tipo: %s>(%s)] : ", sloop, Skindname[bl.loopstmt.id.kind], bl.loopstmt.id.name, bl.loopstmt.id.tnd.name, bl.loopstmt.id.typ.name)
	} else {
		s += fmt.Sprintf("%s [%s: '%s'<Tipo: %s>] : ", sloop, Skindname[bl.loopstmt.id.kind], bl.loopstmt.id.name, bl.loopstmt.id.tnd.name)
	}
	s += bl.loopstmt.ini.StringNd() + ", "
	s += bl.loopstmt.fin.StringNd() + "\n"
	s += bl.loopstmt.body.StringBody(lvl)
	return s
}

func (bl *Stmt) StringAssigmentStmt(lvl int) string {
	s := ""
	lvl++
	tab := strings.Repeat("\t", lvl)
	if debugType {
		s += fmt.Sprintf("%s%s: '%s'(%s) ", tab, Skindname[bl.assigmentstmt.id.kind], bl.assigmentstmt.id.name, bl.assigmentstmt.id.typ.name)
	} else {
		s += fmt.Sprintf("%s%s: '%s' ", tab, Skindname[bl.assigmentstmt.id.kind], bl.assigmentstmt.id.name)
	}
	s += "= " + bl.assigmentstmt.val.StringNd() + "\n"
	return s
}

func (nd *Nd) StringNd() string {
	s := ""
	switch nd.op {
	case ID:
		if debugType {
			s = fmt.Sprintf("[%s: '%s' (%s)]", Skindname[nd.symb.kind], nd.symb.name, nd.symb.typ.name)
		} else {
			s = fmt.Sprintf("[%s: '%s']", Skindname[nd.symb.kind], nd.symb.name)
		}
	case '+', '-', '*', '/':
		if len(nd.args) == 1 {
			s = fmt.Sprintf("%s", tokToStr(nd.op, nil)) + nd.args[0].StringNd()
		} else {
			s = nd.args[0].StringNd() + fmt.Sprintf(" %s ", tokToStr(nd.op, nil)) + nd.args[1].StringNd()
		}
	case INT:
		if debugType {
			s = fmt.Sprintf("%v (%s)", nd.vali, nd.typ.name)
		} else {
			s = fmt.Sprintf("%v", nd.vali)
		}
	case FLOAT:
		if debugType {
			s = fmt.Sprintf("%v (%s)", nd.valf, nd.typ.name)
		} else {
			s = fmt.Sprintf("%v", nd.valf)
		}
	}
	return s
}

//Estructura AST--------------------------------------------------------------------------
type AST struct {
	macro []*Macro
}

type Macro struct {
	id     *Symb
	dparam []*Dparam
	body   Body
}

type Dparam struct {
	tipo *Symb
	id   *Symb
}

type Body struct {
	vars []*Vars
	stmt []*Stmt
}

type Vars struct {
	tipo *Symb
	id   *Symb
}

type Stmt struct {
	op            string
	loopstmt      LoopStmt
	callstmt      CallStmt
	assigmentstmt AssigmentStmt
}

type LoopStmt struct {
	id   *Symb
	ini  *Nd
	fin  *Nd
	body Body
}

type CallStmt struct {
	id    *Symb
	param []*Nd
}

type AssigmentStmt struct {
	id  *Symb
	val *Nd
}

type Nd struct {
	op   int
	symb *Symb
	vali int64
	valf float64
	args []*Nd

	typ  *Type
	line int
	file string
}

//----------------------------------------------------------------------------------------

//NewNd para expresiones
func newNd(op int, symb *Symb, args ...*Nd) *Nd {
	nd := &Nd{op: op, symb: symb, args: args, typ: tnone}
	if symb != nil {
		nd.file = symb.file
		nd.line = symb.line
	} else if len(args) > 1 {
		nd.file = args[0].file
		nd.line = args[0].line
	} else {
		nd.file = file
		nd.line = line
	}
	return nd
}

//NewExpr comprobar el tipo de las expresiones
func newExpr(op int, symb *Symb, args ...*Nd) *Nd {
	nd := newNd(op, symb, args...)
	setexprtype(nd)
	return nd
}

//DEFINIR LOS SIMBOLOS MACRO
func defMacro(symb *Symb) {
	s := getSymbKindMacro(symb.name)
	if s != nil {
		Errorf(" '%s' redeclared in this block\n\t previous declaration at %s:%d", symb.name, s.file, s.line)
	}
	symb.kind = Smacro
	symb.id = ID
	defSymb(symb)
}

//CREO SÍMBOLO TIPO VARIABLE
func defVar(symb *Symb, t *Symb) {
	symb.kind = Svar
	symb.id = ID
	symb.tnd = t
	symb.typ = t.typ
	if checkVarExistSameBlock(symb.name) {
		s := getSymb(symb.name)
		Errorf(" '%s' redeclared in this block\n\t previous declaration at %s:%d", symb.name, s.file, s.line)
	} else {
		defSymb(symb)
	}

}
