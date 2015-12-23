package main

import (
	"fmt"
)

//ENCONTRA MACROS
func (t *AST) findMacro(name string) *Macro {
	for i := range t.macro {
		if t.macro[i].id.name == name {
			return t.macro[i]
		}
	}
	return nil
}

func getParam(st *Stmt) string {
	var ni int64
	var nf float64
	var s string
	str := ""

	for i := range st.callstmt.param {
		ni, nf, s = evalExpresion(st.callstmt.param[i])
		if s == "int64" {
			str += fmt.Sprintf("\t%v", ni)
		} else {
			str += fmt.Sprintf("\t%v", nf)
		}
	}
	return str
}

//TRAS LA LLAMADA ASIGNAR LOS PARÁMETROS
func assigValParam(m *Macro, st *Stmt) {
	var ni int64
	var nf float64

	for i := range st.callstmt.param {
		ni, nf, _ = evalExpresion(st.callstmt.param[i])
		if m.dparam[i].id.typ.kind == Tint64 {
			m.dparam[i].id.vali = ni
		} else {
			m.dparam[i].id.valf = nf
		}
	}
}

//ASIGNACIONES
func valueAssign(st *Stmt) {
	ni, nf, _ := evalExpresion(st.assigmentstmt.val)
	if st.assigmentstmt.id.typ.kind == Tint64 {
		st.assigmentstmt.id.vali = ni
	} else {
		st.assigmentstmt.id.valf = nf
	}
}

//EXPANDIR CALL
func (t *AST) expandirCall(st *Stmt) {
	m := t.findMacro(st.callstmt.id.name)
	if m != nil {
		assigValParam(m, st)
		fmt.Printf("#%s\n", m.id.name)
		for i := range m.body.stmt {
			if m.body.stmt[i].op == "LOOP" {
				fmt.Println("#loop")
				t.expandirloop(m.body.stmt[i])
			}
			if m.body.stmt[i].op == "CALL" {
				t.expandirCall(m.body.stmt[i])
			}
			if m.body.stmt[i].op == "ASSIGN" {
				valueAssign(m.body.stmt[i])
			}
		}
		return
	}
	s := getParam(st)
	fmt.Printf("%s%s\n", st.callstmt.id.name, s)
}

//DEVUELVE EL RESULTADO DE LAS EXPRESIONES
func evalExpresion(nd *Nd) (int64, float64, string) {
	var ni, ni2 int64
	var nf, nf2 float64
	var s string

	switch nd.op {
	case ID:
		if nd.typ.kind == Tint64 {
			return nd.symb.vali, 0.0, "int64"
		} else {
			return 0, nd.symb.valf, "float64"
		}

	case INT:
		return nd.vali, 0.0, "int64"

	case FLOAT:
		return 0, nd.valf, "float64"

	case '+', '-', '*', '/':
		if len(nd.args) == 1 {
			ni, nf, s = evalExpresion(nd.args[0])
			if nd.op == '-' {
				if s == "int64" {
					return -ni, 0.0, s
				}
				return 0, -nf, s
			}
			if s == "int64" {
				return ni, 0.0, s
			}
			return 0, nf, s
		} else {
			ni, nf, s = evalExpresion(nd.args[0])
			ni2, nf2, _ = evalExpresion(nd.args[1])
			if nd.op == '-' {
				if s == "int64" {
					return ni - ni2, 0.0, s
				}
				return 0, nf - nf2, s
			}
			if nd.op == '+' {
				if s == "int64" {
					return ni + ni2, 0.0, s
				}
				return 0, nf + nf2, s
			}
			if nd.op == '*' {
				if s == "int64" {
					return ni * ni2, 0.0, s
				}
				return 0, nf * nf2, s
			}

			if nd.op == '/' {
				if s == "int64" {
					return ni / ni2, 0.0, s
				}
				return 0, nf / nf2, s
			}
		}
	}

	return 0, 0.0, ""
}

//EXPANDIR LOOP
func (t *AST) expandirloop(st *Stmt) {
	inic, _, _ := evalExpresion(st.loopstmt.ini)
	fin, _, _ := evalExpresion(st.loopstmt.fin)

	for st.loopstmt.id.vali = inic; st.loopstmt.id.vali <= fin; st.loopstmt.id.vali++ {
		for j := range st.loopstmt.body.stmt {
			if st.loopstmt.body.stmt[j].op == "LOOP" {
				fmt.Println("#loop")
				t.expandirloop(st.loopstmt.body.stmt[j])
			}
			if st.loopstmt.body.stmt[j].op == "CALL" {
				t.expandirCall(st.loopstmt.body.stmt[j])
			}
			if st.loopstmt.body.stmt[j].op == "ASSIGN" {
				valueAssign(st.loopstmt.body.stmt[j])
			}
		}
	}

}

//EXPANDIR MACROS
func expandirMacros(t *AST) {
	m := t.findMacro("main")
	if m == nil {
		return
	}
	fmt.Printf("#%s\n", m.id.name)
	for i := range m.body.stmt {
		if m.body.stmt[i].op == "LOOP" {
			fmt.Println("#loop")
			t.expandirloop(m.body.stmt[i])
		}
		if m.body.stmt[i].op == "CALL" {
			t.expandirCall(m.body.stmt[i])
		}
		if m.body.stmt[i].op == "ASSIGN" {
			valueAssign(m.body.stmt[i])
		}
	}

}
