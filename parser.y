%{
package main
  
import (
  "bufio"
  "flag"
  "fmt"
  "os"
)
  
%}

%union{
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

%token MACRO
%token <symb> ID
%token <symb> TYPEID
%token <valf> FLOAT
%token <vali> INT
%token LOOP
%token '('
%token ')'
%token '{'
%token '}'
%token ':'
%token ','
%token ';'
%token '='

%left '+','-'
%left '*','/'

%type <ast> file
%type <macros> macros
%type <macro> macr macrid macrhead
%type <dparams> optdparams dparams
%type <dparam> dparam
%type <body> body
%type <variables> optvars vars
%type <variable> var
%type <stmts> optstmts stmts
%type <stmt> stmt
%type <ass> assigmentstmt
%type <calls> callstmt
%type <loops> condicionloop
%type <nd> expr primary num param
%type <list> optparams params


%%
  file 
      : macros
      {
        $$ = new(AST)
        $$.macro = $1
     //   s := envs[1].hash["line"];
    //     fmt.Printf("FIN: %s(%s) -> %s",s.name, s.typ.name, Skindname[s.kind] )
        if debugAST && nerrors == 0 {
          fmt.Println($$.StringAST())
        }
        if nerrors == 0{
          expandirMacros($$)
        }
      }
      | /*empty*/
      { $$ = nil }
      ;
 
   macros 
      : macr
      {
        $$ = []*Macro{$1}
      }
      | macros macr 
      {
        $$ = append($1, $2)
      }
      ;
   
   macr
      : macrhead body
      {
        $$.body = $2
        popEnv()
      }
      ;
      
  macrhead
      : macrid '(' optdparams ')'
      {
        //Guardar el tipo de los parametros
        mkMacro($$.id, $3)
        $$.dparam = $3
   //     fmt.Printf("DEFMACRO:%s(%s) -> %s\n",$$.id.name, $$.id.typ.name, Skindname[$$.id.kind] )
      }
      | macrid '(' error ')'
      {
        Errorf("bad macro header")
      }
      ;
      
  macrid 
      : MACRO ID 
      {
        $$ = new(Macro)
        defMacro($2)
        pushEnv($2.name)
        $$.id = $2   
      } 
      | MACRO error
      {
        pushEnv("macro")
        Errorf("bad macro header")
      }
      ;
  
  optdparams 
        :  dparams
        | /* empty */
        { $$ = nil }
        ;
       
  dparams
        : dparam
        {
          $$ = []*Dparam{$1}
        }
        | dparams ',' dparam
        {
          $$ = append($1, $3)
        }
        ;
     
  dparam 
        : TYPEID ID
        {
          $$ = new(Dparam)
          $$.tipo = $1
       //   s := envs[1].hash["line"]
       //  fmt.Printf("BEFORE: %s(%s) -> %s\n",s.name, s.typ.name, Skindname[s.kind] )
          defVar($2, $1)
          $$.id = $2 
       //   fmt.Printf("DEF VAR:%s(%s) -> %s\n",$$.id.name, $$.id.typ.name, Skindname[$$.id.kind] )
       //    s = envs[1].hash["line"]
       //  fmt.Printf("AFTER: %s(%s) -> %s\n",s.name, s.typ.name, Skindname[s.kind] )
        }
        ;
  
  body
        : '{' optvars optstmts '}'
        {
          $$.vars = $2
          $$.stmt = $3
        }
        ;
        
  optvars
        : vars
        |/*empty*/
        { $$ = nil }
        ;
        
  vars
      : var
      {
        $$ = []*Vars{$1}
      }
      | vars var
      {
        $$ = append($1, $2)
      }
      ;
      
  var
      : TYPEID ID ';'
      {
        $$ = new(Vars)
        $$.tipo = $1
        defVar($2, $1)
        $$.id = $2
      }
      | TYPEID error ';'
      {
        Errorf("bad variable declaration")
      }
      ;
     
      
  optstmts
        : stmts
        |/*empty*/
        { $$ = nil }
        ;

  stmts
      : stmt
      {
        $$ = []*Stmt{$1}
      }
      | stmts stmt
      {
        $$ = append($1, $2)
      }
      ;
      
  stmt
      : tokenloop condicionloop body 
      {
        $$ = new(Stmt)
        $$.loopstmt = $2
        $$.loopstmt.body = $3
        $$.op = "LOOP"
        popEnv()
      }
      | callstmt ';'
      { 
        $$ = new(Stmt)
        $$.callstmt = $1
        $$.op = "CALL"
      }
      | assigmentstmt ';'
      { 
        $$ = new(Stmt)
        $$.assigmentstmt = $1
        $$.op = "ASSIGN"
      }
      | error ';'
      {
        Errorf("bad statement")
      }
      ;
  
  tokenloop
      : LOOP 
      {
        pushEnv("loop")
      }
      ;
      
  condicionloop   
      : ID ':' expr ',' expr
      { 
        err := checkLoopstmt($3, $5)
        if err != nil{  
          Errorf("Loop: %s", err)
        }
        //Defino la variable un simbolo de tipo integer
        s := getBuiltin("int")
        defVar($1, s)
        $$.id = $1
        $$.ini = $3
        $$.fin = $5
        //Error de que ini sea mayor con fin
        ini,_,_ := evalExpresion($$.ini)
        fin,_,_ := evalExpresion($$.fin)
        if ini > fin{
          Errorf("Loop: error")
        }
      }
      ;
      
  assigmentstmt
        : ID '=' expr 
        {
          err := checkAssigmentstmt($1, $3)
          if err != nil{  
            Errorf("Assigment: %s", err)
          }
          $$.id = $1
          $$.val = $3
          
        }
        ;
        
  callstmt      
        : ID '(' optparams ')' 
        {
          err := checkCallstmt($1, $3)
          if err != nil{  
            Errorf("Call: %s", err)
          }
          $$.id = $1
          $$.param = $3
        }
        ;
      
  optparams
          : params
          | /*empty*/
          { $$ = nil}
          ;
          
  params
        : param
        {
          $$ = []*Nd{$1}
        }
        | params ',' param
        {
          $$ = append($1, $3)
        }
        ;
        
  param
        : expr
        ;
        
  expr
      : primary
      | expr '+' expr
      {$$ = newExpr('+', nil, $1, $3)}
      | expr '-' expr
      {$$ = newExpr('-', nil, $1, $3)}
      | expr '*' expr
      {$$ = newExpr('*', nil, $1, $3)}
      | expr '/' expr
      {
        $$ = newExpr('/', nil, $1, $3)
        ni,nf,s := evalExpresion($3)
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
      ;
      
   primary
        : '-' primary
        { $$ = newExpr('-', nil, $2 ) }
        | '+' primary
        { $$ = newExpr('+', nil, $2 ) }
        | ID
        { $$ = newExpr(ID, $1) }
        | num 
        ;
               
   num
      : FLOAT
      { $$ = &Nd{op: FLOAT, valf: $1, typ: tfloat64} }
      | INT
      { $$ = &Nd{op: INT, vali: $1, typ: tint64} }
      ;

%%

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