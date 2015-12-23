package main

import (
	"fmt"
)

type Tkind int

const (
	Tnone Tkind = iota
	Tint64
	Tfloat64
	Tmacro
)

type Type struct {
	kind Tkind
	name string
	id   int
	univ bool

	args []*Type
}

var debugType bool

func mktype(kind Tkind, name string) *Type {
	id := len(tdefs)
	t := &Type{kind: kind, name: name, id: id}
	tdefs = append(tdefs, t)
	return t
}

//DEFINIR TYP PARA MACROS NO PREDEFINIDAS
func mkMacro(s *Symb, args []*Dparam) {
	s.typ = mktype(Tmacro, "anon")
	for i := range args {
		s.typ.args = append(s.typ.args, args[i].id.typ)
	}
}

func tcompat(t1 *Type, t2 *Type) bool {
	if t1 == t2 {
		return true
	}
	return (t1.univ || t2.univ) && t1.kind == t2.kind
}

//COMPROBAR ARGUMENTOS DE LAS OPERACIONES (operaciones)
func checkArgumentos(args []*Nd) error {
	for i := range args {
		if !tcompat(args[0].typ, args[i].typ) {
			return fmt.Errorf("Incompatible types")
		}
	}
	return nil
}

//COMPROBAR ASIGNACIONES
func checkAssigmentstmt(symb *Symb, nd *Nd) error {
	if !checkVarExist(symb.name) {
		//Asigno tnone ya que no esta definido
		symb.typ = tnone
		return fmt.Errorf("var '%s' is not defined", symb.name)
	}
	if !tcompat(symb.typ, nd.typ) {
		return fmt.Errorf("Incompatible types")
	}
	return nil
}

//COMPROBAR EL ELEMENTO INI Y FIN DEL lOOP
func checkLoopstmt(ini *Nd, fin *Nd) error {
	if ini.typ != tint64 || fin.typ != tint64 {
		return fmt.Errorf("Incompatible types")
	}
	return nil
}

//COMPROBAR LOS PARAMETROS DE LAS MACROS EN LLAMADAS A MACROS
func checkCallstmt(symb *Symb, param []*Nd) error {
	s := getSymbKindMacro(symb.name)
	if s == nil {
		return fmt.Errorf("macro: '%s' is not defined", symb.name)
	}
	if len(param) != len(s.typ.args) {
		if len(param) > len(s.typ.args) {
			return fmt.Errorf("too many arguments in call to %s (%d - %d )", s.name, len(param), len(s.typ.args))
		} else {
			return fmt.Errorf("not enough arguments in call to %s", s.name)
		}
	}
	for i := range s.typ.args {
		if param[i].typ != s.typ.args[i] {
			return fmt.Errorf("cannot use %s as type %s in call to  %s", param[i].typ.name, s.typ.args[i].name, s.name)
		}
	}
	return nil
}

func setexprtype(nd *Nd) {
	nd.typ = tnone

	switch nd.op {
	case ID:
		if nd.symb == nil {
			return
		}
		s := getSymb(nd.symb.name)
		if s == nil {
			Errorf("var '%s' is not defined", nd.symb.name)
			return
		}
		//Nodo tiene tipo Tnone, hay que asignarle el del símbolo definido
		nd.symb = s
		nd.typ = s.typ
	case '+', '-', '*', '/':
		nd.typ = nd.args[0].typ
		if err := checkArgumentos(nd.args); err != nil {
			Errorf("'%s': %s", tokToStr(nd.op, nil), err)
			return
		}
	}
}
