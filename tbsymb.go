package main

import (
	"fmt"
)

//Tipo Skind------------------------------------------------------------------------------
type Skind int

const (
	Snone Skind = iota
	Skey
	Smacro
	Svar
	Sname
	Stype
)

var debugTsymb bool
var Skindname = [...]string{
	"none", "key", "macro", "var", "name", "type",
}

//Tipo Symb-------------------------------------------------------------------------------
type Symb struct {
	name string
	kind Skind
	id   int

	vali int64
	valf float64

	tnd *Symb
	typ *Type

	file string
	line int
}

var (
	tnone     = &Type{kind: Tnone, name: "Tnone"}
	tuint64   = &Type{kind: Tint64, name: "Tuint64", univ: true}
	tint64    = &Type{kind: Tint64, name: "Tint64", args: []*Type{tuint64}}
	tufloat64 = &Type{kind: Tfloat64, name: "Tfuloat64", univ: true}
	tfloat64  = &Type{kind: Tfloat64, name: "Tfloat64", args: []*Type{tufloat64}}

	/* builtins */
	tmiiii = &Type{
		kind: Tmacro, name: "Tmiiii", univ: true,
		args: []*Type{tint64, tint64, tint64, tint64},
	}
	tmiifi = &Type{
		kind: Tmacro, name: "Tmiifi", univ: true,
		args: []*Type{tint64, tint64, tfloat64, tint64},
	}
)
var tdefs []*Type = []*Type{
	tnone, tuint64, tint64, tufloat64, tfloat64,
	tmiiii, tmiifi,
}

var keywords = [...]Symb{
	Symb{name: "macro", id: MACRO},
	Symb{name: "loop", id: LOOP},
}

var types = [...]Symb{
	Symb{name: "int", id: TYPEID, typ: tint64},
	Symb{name: "float", id: TYPEID, typ: tfloat64},
}

var macros = [...]Symb{
	Symb{name: "circle", id: ID, typ: tmiiii},
	Symb{name: "rect", id: ID, typ: tmiifi},
}

type Env struct {
	tag  string
	hash map[string]*Symb
}

var envs []*Env

//Inicializar los builtins----------------------------------------------------------------
func init() {
	envs = []*Env{}
	pushEnv("bltin")
	bltin := envs[0]

	for i := range keywords {
		s := keywords[i]
		bltin.hash[s.name] = &s
		s.kind = Skey
		s.file = "builtin"
	}

	for i := range types {
		s := types[i]
		bltin.hash[s.name] = &s
		s.kind = Stype
		s.file = "builtin"
	}

	for i := range macros {
		s := macros[i]
		bltin.hash[s.name] = &s
		s.kind = Smacro
		s.file = "builtin"
	}
	pushEnv("toplvl")
}

//PUSH------------------------------------------------------------------------------------
func pushEnv(tag string) {
	env := &Env{tag: tag, hash: map[string]*Symb{}}
	envs = append(envs, env)
	if debugTsymb {
		fmt.Printf("pushenv %s\n", tag)
	}
}

//POP-------------------------------------------------------------------------------------
func popEnv() {
	if len(envs) == 1 {
		panic("bug: attempt to pop builtin env")
	}
	env := envs[len(envs)-1]
	envs = envs[:len(envs)-1]
	if debugTsymb {
		fmt.Printf("popenv %s\n", env.tag)
	}
}

func getBuiltin(n string) *Symb {
	if s, ok := envs[0].hash[n]; ok {
		return s
	}
	return nil
}

func getSymb(n string) *Symb {
	for i := len(envs) - 1; i >= 0; i-- {
		if s, ok := envs[i].hash[n]; ok {
			return s
		}
	}
	return nil
}

//Buscar si un símbolo existe en un ámbito paramos en las macros ya que no son variables
//Es para comprobar si existe, y recorro más de un ambito por los loops,
//Esta función no es usada a la hora de definir nuevas var, ya que no valdría ya que ve más haya de un ámbito
//No miro por debako ya que no hay variables globales
func checkVarExist(n string) bool {
	for i := len(envs) - 1; i >= 2; i-- {
		s, ok := envs[i].hash[n]
		if ok && s.kind == Svar {
			return true
		}
	}
	return false
}

//Para definicion de variables, comprobar que existe un símbolo en el mismo ámbito
func checkVarExistSameBlock(n string) bool {
	i := len(envs) - 1
	if _, ok := envs[i].hash[n]; ok {
		return true
	}
	return false
}

//BUSCAR SI LA MACRO ESTA DEFINIDA
func getSymbKindMacro(n string) *Symb {
	for i := len(envs) - 1; i >= 0; i-- {
		s, ok := envs[i].hash[n]
		//fmt.Println(envs[i].hash["line"])
		if ok && s.kind == Smacro {
			return s
		}
	}
	return nil
}

func defSymb(s *Symb) {
	envs[len(envs)-1].hash[s.name] = s
	if debugTsymb {
		fmt.Printf("symb %s defined\n", s.name)
	}
	envs[len(envs)-1].hash[s.name] = s
	if s.typ == nil {
		s.typ = tnone
	}
}
