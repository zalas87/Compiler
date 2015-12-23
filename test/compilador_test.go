package main

import (
  "bufio"
  "fmt"
  "strings"
  "testing"
)

//---------------------------------------TEST PARSER------------------------------------------
//SCAN FICHERO VACIO
func Test_Scan_1(t *testing.T) {
  r := bufio.NewReader(strings.NewReader(""))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors == 0{
    fmt.Printf("Test_scan_1 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SCAN UNA MACRO
func Test_Scan_2(t *testing.T) {
  r := bufio.NewReader(strings.NewReader("macro line2 (int x, int y) {}"))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors == 0{
    fmt.Printf("Test_scan_2 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR EN UNA MACRO
func Test_Scan_3(t *testing.T) {
  //ya que se acumulan
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro line3 (int x, in y){} "))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_3 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR EN UNA DECLARACIÓN DE VARIABLES 
func Test_Scan_4(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro line4 (int x, int y){int 2; } "))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_4 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}


//ERROR EN LOS STMTS
func Test_Scan_5(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro line5 (int x, int y){int z; z =*5; } "))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_5 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SCAN MACRO COMPLETA DEL EJEMPLO DEL EXÁMEN(CORREGIDO ERROR QUE TIENE EN EL ENUNCIADO EN EL PARÁMETRO)
func Test_Scan_6(t *testing.T) {
  nerrors = 0;
  text:= "macro line6(int x, int y){loop i:0,x{circle(2, 3, y, 5);}}"
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors == 0{
    fmt.Printf("Test_scan_6 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//-----------------------------------------TEST TBSYMB Y ÁMBITOS-------------------------------------
//MACRO DUPLICADA EN UN ÁMBITOS
func Test_Scan_7(t *testing.T) {
  nerrors = 0;
  text:= "macro hola(){} macro hola(int z){}"
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_7 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//VARIABLE DUPLICADA EN UN ÁMBITOS
func Test_Scan_8(t *testing.T) {
  nerrors = 0;
  text:= "macro hola1(int w){int w;} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_8 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//VARIABLE CON EL MISMO ID EN DISTINTOS ÁMBITOS
func Test_Scan_9(t *testing.T) {
  nerrors = 0;
  text:= "macro hola2(){int w; loop i: 0,10{int w;}} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")

  yyParse(l)
  if nerrors == 0{
    fmt.Printf("Test_scan_9 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//--------------------------------------TEST AST-------------------------------------------------
//SI NO HAY ERRORES SE IMPRIME EL AST
func Test_Scan_10(t *testing.T) {
  nerrors = 0;
  text:= "macro buenas(){int w; rect(5,5,.5,0xff);loop i: 0,10{int w;}} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  debugAST = true
  yyParse(l)
  
  if nerrors == 0{
    fmt.Printf("Test_scan_10 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugAST = false
}

//SI HAY ERRORES NO SE IMPRIME EL AST
func Test_Scan_11(t *testing.T) {
  nerrors = 0;
  text:= "macro buenas1 (){int w; rect(5,5 0.5,0xff);loop i: 0,10{int w;}} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  debugAST = true
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_11 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugAST = false
}

//--------------------------------------------TEST TIPOS DE DATOS-------------------------------------------
//ERROR EN LAS EXPRESIONES AL INTENTAR INTEROPERAR CON DISTINTOS TIPOS
func Test_Scan_12(t *testing.T) {
  nerrors = 0;
  text:= "macro hello (){int w; w = 3 + 3.5} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_12 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR SI NO COINCIDEN LOS TIPOS DE LOS PARÁMETRO AL REALIZAR UNA LLAMADA
func Test_Scan_13(t *testing.T) {
  nerrors = 0;
  text:= "macro hello1 (){circle(0.2,2,2,5)} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_13 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}


//ERROR SI EN LA LLAMADA A UNA FUNCIÓN NO COINCIDE EL NÚMERO DE PARÁMETROS
func Test_Scan_14(t *testing.T) {
  nerrors = 0;
  text:= "macro hello2 (){circle(2,2,5)} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_14 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR SI EN LOOPSTMT EL ARRANQUE O FIN NO SON DE TIPO INT
func Test_Scan_15(t *testing.T) {
  nerrors = 0;
  text:= "macro hello3 (){loop i: 5, 3.5{}} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_15 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR SI EN LOOPSTMT SI EL ARRANQUE ES MAYOR AL FIN 
func Test_Scan_16(t *testing.T) {
  nerrors = 0;
  text:= "macro hello4 (){loop i: 5, 3{}} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_16 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//ERROR EN ASIGNACIONES DE DISTINTOS TIPOS
func Test_Scan_17(t *testing.T) {
  nerrors = 0;
  text:= "macro hello5 (int x){x = 3.5;} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_17 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//-----------------------------TEST INTERPRETE (EXTENSIÓN DE MACROS)---------------------------
//ERROR AL DIVIDIR ENTRE 0
func Test_Scan_18(t *testing.T) {
  nerrors = 0;
  text:= "macro hi (int x){x = 7/0;} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_18 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SI NO HAY ERRORES SE REALIZA LA EXTENSIÓN DE LAS MACROS
func Test_Scan_19(t *testing.T) {
  nerrors = 0;
  text:= "macro hi2 (int x){loop i: 0,x{circle(2,i,2,5);}} macro main(){hi2(5);} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors == 0{
    fmt.Printf("Test_scan_19 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SI HAY ERRORES NO SE REALIZA LA EXTENSIÓN DE LAS MACROS
func Test_Scan_20(t *testing.T) {
  nerrors = 0;
  text:= "macro hi3 (int x){loop i: 0,x{circle(2,i,2,5);}}  main(){hi2(5);} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_20 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SI LLAMAS A UNA VARIABLE NO DEFINIDA
func Test_Scan_21(t *testing.T) {
  nerrors = 0;
  text:= "macro hi4 (int x){z = 3} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_21 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//SI LLAMAS A UNA MACRO NO DEFINIDA
func Test_Scan_22(t *testing.T) {
  nerrors = 0;
  text:= "macro hi4 (int x){zeta()} "
  r := bufio.NewReader(strings.NewReader(text))
  l := NewLex(r, "String")
  yyParse(l)
  
  if nerrors != 0{
    fmt.Printf("Test_scan_22 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
}

//-------------------------------------TEST LEXER-------------------------------------
//SCAN UNA MACRO
func Test_Scan_23(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro adios (int x, int y) {float δ;}"))
  l := NewLex(r, "String")

  debugLex = true
  yyParse(l)
  if nerrors == 0{
    fmt.Printf("Test_scan_23 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugLex = false
}

//TOKEN INVALIDO
func Test_Scan_24(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro adios (int x, int y) {float δ; %}"))
  l := NewLex(r, "String")

  debugLex = true
  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_24 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugLex = false
}

//TOKEN INVALIDO
func Test_Scan_25(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro adios (int x, int y) {float δ; δ= .25a}"))
  l := NewLex(r, "String")

  debugLex = true
  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_25 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugLex = false
}

//TOKEN INVALIDO
func Test_Scan_26(t *testing.T) {
  nerrors = 0;
  r := bufio.NewReader(strings.NewReader("macro adios (int x, int y) {int δ; δ= 1x153f}"))
  l := NewLex(r, "String")

  debugLex = true
  yyParse(l)
  if nerrors != 0{
    fmt.Printf("Test_scan_26 Passed\n")
  }else{
    t.Error("Scan did not work as expected.")
  }
  debugLex = false
}