
//line parser.y:2
package main
import __yyfmt__ "fmt"
//line parser.y:2
		  
import (
  "bufio"
  "flag"
  "fmt"
  "os"
)
  

//line parser.y:13
type yySymType struct{
	yys int
  symb *Symb
  valf float64
  vali int64
  ast *AST
  macros []*Macro
  macro *Macro
  dparams []*Dparam
  dparam *Dparam
  body Body
  variables []*Vars
  variable *Vars
  stmts []*Stmt
  stmt *Stmt
  ass AssigmentStmt
  calls CallStmt
  loops LoopStmt
  nd *Nd
  list []*Nd
}

const MACRO = 57346
const ID = 57347
const TYPEID = 57348
const FLOAT = 57349
const INT = 57350
const LOOP = 57351

var yyToknames = []string{
	"MACRO",
	"ID",
	"TYPEID",
	"FLOAT",
	"INT",
	"LOOP",
	" (",
	" )",
	" {",
	" }",
	" :",
	" ,",
	" ;",
	" =",
	" +",
	" -",
	" *",
	" /",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:369


func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) != 1 {
    fmt.Println("Invalid command line -> go run *.go [documento]")
    os.Exit(1)
  }

  file, err := os.Open(args[0]) //base
  if err != nil {
    fmt.Printf("error opening file= %s", err)
    os.Exit(1)
  }
  
  l := NewLex(bufio.NewReader(file), args[0])
  //debugLex = true
  //debugTsymb = true
  //debugAST = true
  //Solo se ve con el AST activo
  //debugType = true
  yyParse(l)
  
  err = file.Close()
  if err != nil {
    fmt.Printf("error closing file= %s", err)
    os.Exit(1)
  }
  os.Exit(nerrors)
}
//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 10,
	11, 11,
	-2, 0,
	-1, 13,
	13, 23,
	-2, 0,
	-1, 23,
	13, 22,
	-2, 0,
}

const yyNprod = 50
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 86

var yyAct = []int{

	63, 53, 62, 8, 20, 73, 24, 48, 65, 66,
	67, 68, 65, 66, 67, 68, 67, 68, 56, 46,
	58, 59, 47, 44, 43, 72, 45, 42, 15, 36,
	39, 55, 54, 51, 38, 9, 71, 35, 34, 10,
	21, 49, 28, 31, 50, 30, 52, 18, 16, 29,
	33, 21, 64, 32, 41, 12, 69, 70, 11, 37,
	6, 3, 25, 61, 7, 60, 74, 75, 76, 77,
	57, 40, 26, 27, 79, 78, 23, 22, 14, 13,
	19, 17, 4, 5, 2, 1,
}
var yyPact = []int{

	56, -1000, 56, -1000, 23, 29, 53, -1000, -1000, 42,
	45, -1000, -1000, 40, 42, -1000, 48, 27, 26, 14,
	-1000, 54, 21, 40, -1000, 49, 11, 8, 7, -1000,
	9, -1000, 6, -9, -1000, -1000, 34, -1000, -1000, -1000,
	23, 19, -1000, -1000, -1000, 13, 13, -1000, -1000, -1000,
	-1000, 13, -6, -1000, 13, 13, -1000, -1000, -1000, -1000,
	25, 10, -1000, -6, -10, 13, 13, 13, 13, -1000,
	-1000, -1000, 13, 13, -4, -4, -1000, -1000, -1000, -6,
}
var yyPgo = []int{

	0, 85, 84, 61, 83, 82, 81, 80, 4, 3,
	79, 78, 28, 77, 76, 6, 73, 72, 71, 0,
	1, 70, 2, 65, 63, 62,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 3, 5, 5, 4, 4,
	6, 6, 7, 7, 8, 9, 10, 10, 11, 11,
	12, 12, 13, 13, 14, 14, 15, 15, 15, 15,
	25, 18, 16, 17, 23, 23, 24, 24, 22, 19,
	19, 19, 19, 19, 20, 20, 20, 20, 21, 21,
}
var yyR2 = []int{

	0, 1, 0, 1, 2, 2, 4, 4, 2, 2,
	1, 0, 1, 3, 2, 4, 1, 0, 1, 2,
	3, 3, 1, 0, 1, 2, 3, 2, 2, 2,
	1, 5, 3, 4, 1, 0, 1, 3, 1, 1,
	3, 3, 3, 3, 2, 2, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, -5, -4, 4, -3, -9, 12,
	10, 5, 2, -10, -11, -12, 6, -6, 2, -7,
	-8, 6, -13, -14, -15, -25, -17, -16, 2, 9,
	5, -12, 5, 2, 11, 11, 15, 5, 13, -15,
	-18, 5, 16, 16, 16, 17, 10, 16, 16, -8,
	-9, 14, -19, -20, 19, 18, 5, -21, 7, 8,
	-23, -24, -22, -19, -19, 18, 19, 20, 21, -20,
	-20, 11, 15, 15, -19, -19, -19, -19, -22, -19,
}
var yyDef = []int{

	2, -2, 1, 3, 0, 0, 0, 4, 5, 17,
	-2, 8, 9, -2, 16, 18, 0, 0, 0, 10,
	12, 0, 0, -2, 24, 0, 0, 0, 0, 30,
	0, 19, 0, 0, 6, 7, 0, 14, 15, 25,
	0, 0, 27, 28, 29, 0, 35, 20, 21, 13,
	26, 0, 32, 39, 0, 0, 46, 47, 48, 49,
	0, 34, 36, 38, 0, 0, 0, 0, 0, 44,
	45, 33, 0, 0, 40, 41, 42, 43, 37, 31,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	10, 11, 20, 18, 15, 19, 3, 21, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 14, 16,
	3, 17, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 12, 3, 13,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.y:72
		{
	        yyVAL.ast = new(AST)
	        yyVAL.ast.macro = yyS[yypt-0].macros
	     //   s := envs[1].hash["line"];
    //     fmt.Printf("FIN: %s(%s) -> %s",s.name, s.typ.name, Skindname[s.kind] )
        if debugAST && nerrors == 0 {
	          fmt.Println(yyVAL.ast.StringAST())
	        }
	        if nerrors == 0{
	          expandirMacros(yyVAL.ast)
	        }
	      }
	case 2:
		//line parser.y:85
		{ yyVAL.ast = nil }
	case 3:
		//line parser.y:90
		{
	        yyVAL.macros = []*Macro{yyS[yypt-0].macro}
	      }
	case 4:
		//line parser.y:94
		{
	        yyVAL.macros = append(yyS[yypt-1].macros, yyS[yypt-0].macro)
	      }
	case 5:
		//line parser.y:101
		{
	        yyVAL.macro.body = yyS[yypt-0].body
	        popEnv()
	      }
	case 6:
		//line parser.y:109
		{
	        //Guardar el tipo de los parametros
        mkMacro(yyVAL.macro.id, yyS[yypt-1].dparams)
	        yyVAL.macro.dparam = yyS[yypt-1].dparams
	   //     fmt.Printf("DEFMACRO:%s(%s) -> %s\n",$$.id.name, $$.id.typ.name, Skindname[$$.id.kind] )
      }
	case 7:
		//line parser.y:116
		{
	        Errorf("bad macro header")
	      }
	case 8:
		//line parser.y:123
		{
	        yyVAL.macro = new(Macro)
	        defMacro(yyS[yypt-0].symb)
	        pushEnv(yyS[yypt-0].symb.name)
	        yyVAL.macro.id = yyS[yypt-0].symb   
	      }
	case 9:
		//line parser.y:130
		{
	        pushEnv("macro")
	        Errorf("bad macro header")
	      }
	case 10:
		yyVAL.dparams = yyS[yypt-0].dparams
	case 11:
		//line parser.y:139
		{ yyVAL.dparams = nil }
	case 12:
		//line parser.y:144
		{
	          yyVAL.dparams = []*Dparam{yyS[yypt-0].dparam}
	        }
	case 13:
		//line parser.y:148
		{
	          yyVAL.dparams = append(yyS[yypt-2].dparams, yyS[yypt-0].dparam)
	        }
	case 14:
		//line parser.y:155
		{
	          yyVAL.dparam = new(Dparam)
	          yyVAL.dparam.tipo = yyS[yypt-1].symb
	       //   s := envs[1].hash["line"]
       //  fmt.Printf("BEFORE: %s(%s) -> %s\n",s.name, s.typ.name, Skindname[s.kind] )
          defVar(yyS[yypt-0].symb, yyS[yypt-1].symb)
	          yyVAL.dparam.id = yyS[yypt-0].symb 
	       //   fmt.Printf("DEF VAR:%s(%s) -> %s\n",$$.id.name, $$.id.typ.name, Skindname[$$.id.kind] )
       //    s = envs[1].hash["line"]
       //  fmt.Printf("AFTER: %s(%s) -> %s\n",s.name, s.typ.name, Skindname[s.kind] )
        }
	case 15:
		//line parser.y:170
		{
	          yyVAL.body.vars = yyS[yypt-2].variables
	          yyVAL.body.stmt = yyS[yypt-1].stmts
	        }
	case 16:
		yyVAL.variables = yyS[yypt-0].variables
	case 17:
		//line parser.y:179
		{ yyVAL.variables = nil }
	case 18:
		//line parser.y:184
		{
	        yyVAL.variables = []*Vars{yyS[yypt-0].variable}
	      }
	case 19:
		//line parser.y:188
		{
	        yyVAL.variables = append(yyS[yypt-1].variables, yyS[yypt-0].variable)
	      }
	case 20:
		//line parser.y:195
		{
	        yyVAL.variable = new(Vars)
	        yyVAL.variable.tipo = yyS[yypt-2].symb
	        defVar(yyS[yypt-1].symb, yyS[yypt-2].symb)
	        yyVAL.variable.id = yyS[yypt-1].symb
	      }
	case 21:
		//line parser.y:202
		{
	        Errorf("bad variable declaration")
	      }
	case 22:
		yyVAL.stmts = yyS[yypt-0].stmts
	case 23:
		//line parser.y:211
		{ yyVAL.stmts = nil }
	case 24:
		//line parser.y:216
		{
	        yyVAL.stmts = []*Stmt{yyS[yypt-0].stmt}
	      }
	case 25:
		//line parser.y:220
		{
	        yyVAL.stmts = append(yyS[yypt-1].stmts, yyS[yypt-0].stmt)
	      }
	case 26:
		//line parser.y:227
		{
	        yyVAL.stmt = new(Stmt)
	        yyVAL.stmt.loopstmt = yyS[yypt-1].loops
	        yyVAL.stmt.loopstmt.body = yyS[yypt-0].body
	        yyVAL.stmt.op = "LOOP"
	        popEnv()
	      }
	case 27:
		//line parser.y:235
		{ 
	        yyVAL.stmt = new(Stmt)
	        yyVAL.stmt.callstmt = yyS[yypt-1].calls
	        yyVAL.stmt.op = "CALL"
	      }
	case 28:
		//line parser.y:241
		{ 
	        yyVAL.stmt = new(Stmt)
	        yyVAL.stmt.assigmentstmt = yyS[yypt-1].ass
	        yyVAL.stmt.op = "ASSIGN"
	      }
	case 29:
		//line parser.y:247
		{
	        Errorf("bad statement")
	      }
	case 30:
		//line parser.y:254
		{
	        pushEnv("loop")
	      }
	case 31:
		//line parser.y:261
		{ 
	        err := checkLoopstmt(yyS[yypt-2].nd, yyS[yypt-0].nd)
	        if err != nil{  
	          Errorf("Loop: %s", err)
	        }
	        //Defino la variable un simbolo de tipo integer
        s := getBuiltin("int")
	        defVar(yyS[yypt-4].symb, s)
	        yyVAL.loops.id = yyS[yypt-4].symb
	        yyVAL.loops.ini = yyS[yypt-2].nd
	        yyVAL.loops.fin = yyS[yypt-0].nd
	        //Error de que ini sea mayor con fin
        ini,_,_ := evalExpresion(yyVAL.loops.ini)
	        fin,_,_ := evalExpresion(yyVAL.loops.fin)
	        if ini > fin{
	          Errorf("Loop: error")
	        }
	      }
	case 32:
		//line parser.y:283
		{
	          err := checkAssigmentstmt(yyS[yypt-2].symb, yyS[yypt-0].nd)
	          if err != nil{  
	            Errorf("Assigment: %s", err)
	          }
	          yyVAL.ass.id = yyS[yypt-2].symb
	          yyVAL.ass.val = yyS[yypt-0].nd
	          
	        }
	case 33:
		//line parser.y:296
		{
	          err := checkCallstmt(yyS[yypt-3].symb, yyS[yypt-1].list)
	          if err != nil{  
	            Errorf("Call: %s", err)
	          }
	          yyVAL.calls.id = yyS[yypt-3].symb
	          yyVAL.calls.param = yyS[yypt-1].list
	        }
	case 34:
		yyVAL.list = yyS[yypt-0].list
	case 35:
		//line parser.y:309
		{ yyVAL.list = nil}
	case 36:
		//line parser.y:314
		{
	          yyVAL.list = []*Nd{yyS[yypt-0].nd}
	        }
	case 37:
		//line parser.y:318
		{
	          yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].nd)
	        }
	case 38:
		yyVAL.nd = yyS[yypt-0].nd
	case 39:
		yyVAL.nd = yyS[yypt-0].nd
	case 40:
		//line parser.y:330
		{yyVAL.nd = newExpr('+', nil, yyS[yypt-2].nd, yyS[yypt-0].nd)}
	case 41:
		//line parser.y:332
		{yyVAL.nd = newExpr('-', nil, yyS[yypt-2].nd, yyS[yypt-0].nd)}
	case 42:
		//line parser.y:334
		{yyVAL.nd = newExpr('*', nil, yyS[yypt-2].nd, yyS[yypt-0].nd)}
	case 43:
		//line parser.y:336
		{
	        yyVAL.nd = newExpr('/', nil, yyS[yypt-2].nd, yyS[yypt-0].nd)
	        ni,nf,s := evalExpresion(yyS[yypt-0].nd)
	        if s == "int64"{
	          if ni == 0{
	            Errorf("divide by 0")
	          }
	        }else{
	          if nf < 0.000001{
	            Errorf("divide by 0")
	          }
	        }
	        
	      }
	case 44:
		//line parser.y:354
		{ yyVAL.nd = newExpr('-', nil, yyS[yypt-0].nd ) }
	case 45:
		//line parser.y:356
		{ yyVAL.nd = newExpr('+', nil, yyS[yypt-0].nd ) }
	case 46:
		//line parser.y:358
		{ yyVAL.nd = newExpr(ID, yyS[yypt-0].symb) }
	case 47:
		yyVAL.nd = yyS[yypt-0].nd
	case 48:
		//line parser.y:364
		{ yyVAL.nd = &Nd{op: FLOAT, valf: yyS[yypt-0].valf, typ: tfloat64} }
	case 49:
		//line parser.y:366
		{ yyVAL.nd = &Nd{op: INT, vali: yyS[yypt-0].vali, typ: tint64} }
	}
	goto yystack /* stack new state and value */
}
