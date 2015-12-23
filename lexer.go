package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

type yyLex interface {
	Lex(lval *yySymType) int
	Error(e string)
}

type lexema struct {
	in    *bufio.Reader
	saved int
	value []rune

	file string
	line int
}

var file string
var line int
var debugLex bool
var nerrors int

func NewLex(text *bufio.Reader, doc string) *lexema {
	return &lexema{in: text, line: 1, file: doc}
}

//SKIP BLANKS
func (l *lexema) skipBlanks() error {
	for {
		c, _, err := l.in.ReadRune()
		if err != nil {
			return err
		}
		if c == '#' {
			for c != '\n' {
				if c, _, err = l.in.ReadRune(); err != nil {
					return err
				}
			}
		}
		if c != '\t' && c != '\n' && c != ' ' {
			l.in.UnreadRune()
			return nil
		}
		if c == '\n' {
			l.line++
		}
	}
}

//GOT RUNE--------------------------------------------------------------------------------
func (l *lexema) gotRune(r rune) {
	l.value = append(l.value, r)
}

//CONVERSOR DE RUNA A INTEGER-------------------------------------------------------------
func runeToInt(r rune) int64 {
	var i64 int64

	i64, _ = strconv.ParseInt(string(r), 16, 0)
	return i64
}

//FUNCION PARA IMPRIMIR LOS TOKENS
func tokToStr(i int, lval *yySymType) string {
	switch i {
	case ID:
		return fmt.Sprintf("Id->[%s<%s>]", Skindname[lval.symb.kind], lval.symb.name)
	case TYPEID:
		return fmt.Sprintf("Typeid->[%s<%s>]", Skindname[lval.symb.kind], lval.symb.name)
	case INT:
		return fmt.Sprintf("Int<%d>", lval.vali)
	case FLOAT:
		return fmt.Sprintf("Float<%v>", lval.valf)
	case MACRO:
		return "Macro"
	case LOOP:
		return "Loop"
	case 0:
		return "eof"
	case ')':
		return ")"
	case '(':
		return "("
	case '{':
		return "{"
	case '}':
		return "}"
	case ',':
		return ","
	case '.':
		return "."
	case ':':
		return ":"
	case '+':
		return "+"
	case '-':
		return "-"
	case ';':
		return ";"
	case '*':
		return "*"
	case '/':
		return "/"
	case '=':
		return "="
	default:
		return fmt.Sprintf("BADTOK:<%d>", i)
	}
}

func Errorf(s string, v ...interface{}) {
	fmt.Printf("%s:%d: ", file, line)
	fmt.Printf(s, v...)
	fmt.Printf("\n")
	nerrors++

	if nerrors > 5 {
		fmt.Printf("too many errors\n")
		os.Exit(1)
	}
}

func (l *lexema) Error(s string) {
	Errorf("near '%s': %s", string(l.value), s)
}

//FUNCTION Lex----------------------------------------------------------------------------
func (l *lexema) Lex(lval *yySymType) int {
	if l.saved != 0 {
		t := l.saved
		l.saved = 0
		return t
	}

	t := l.lex(lval)
	if debugLex {
		fmt.Printf("tok %s\n", tokToStr(t, lval))
	}
	return t
}

// FUNCTION lex --------------------------------------------------------------------------
func (l *lexema) lex(lval *yySymType) int {
	l.value = l.value[:0]

	if err := l.skipBlanks(); err != nil {
		if err == io.EOF {
			return 0
		}
	}

	file = l.file
	line = l.line

	r, _, err := l.in.ReadRune()
	if err != nil {
		return 0
	}

	l.in.UnreadRune()
	if unicode.IsLetter(r) {
		return l.scanWord(lval)
	}
	if unicode.Is(unicode.Greek, r) {
		return l.scanWordGreek(lval)
	}
	if unicode.IsDigit(r) {
		return l.scanNumb(false, lval)
	}
	return l.scanPunct(lval)
}

//SCAN PUNCT-------------------------------------------------------------------------------
func (l *lexema) scanPunct(lval *yySymType) int {
	r, _, err := l.in.ReadRune()
	if err != nil {
		return 0
	}
	switch r {
	case '(', ')', '{', '}', ',', ':', ';', '+', '-', '*', '/', '=':
		l.gotRune(r)
		return int(r)

	case '.': // '.' -> .64 or .
		l.gotRune(r)
		r, _, err := l.in.ReadRune()
		if err != nil {
			if err == io.EOF {
				return '.'
			}
			return ' '
		}
		if unicode.IsNumber(r) {
			l.in.UnreadRune()
			return l.scanDecimal(lval)
		}
		l.in.UnreadRune()
		return '.'
	}
	l.gotRune(r)
	Errorf("Bad character '%c' in input", r)
	return ' '
}

//SCAN WORD-------------------------------------------------------------------------------
func (l *lexema) scanWord(lval *yySymType) int {
	for {
		r, _, err := l.in.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return ' '
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			l.in.UnreadRune()
			break
		}
		l.gotRune(r)
	}
	name := string(l.value)
	s := getSymb(name)
	if s == nil {
		s = &Symb{
			name: name,
			kind: Sname,
			id:   ID,
			file: l.file,
			line: l.line,
		}
	}
	lval.symb = s
	return s.id
}

//SCAN GREEK WORD-------------------------------------------------------------------------
func (l *lexema) scanWordGreek(lval *yySymType) int {
	for {
		r, _, err := l.in.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return ' '
		}
		if !unicode.Is(unicode.Greek, r) && !unicode.IsDigit(r) {
			l.in.UnreadRune()
			break
		}
		l.gotRune(r)
	}
	name := string(l.value)
	s := getSymb(name)
	if s == nil {
		s = &Symb{
			name: name,
			kind: Sname,
			id:   ID,
			file: l.file,
			line: l.line,
		}
	}
	lval.symb = s
	return s.id
}

//SCAN NUMB ------------------------------------------------------------------------------
func (l *lexema) scanNumb(isFloat bool, lval *yySymType) int {

	if isFloat {
		return l.scanDecimal(lval)
	}

	return l.scanInt(lval)
}

//SCAN INTEGER----------------------------------------------------------------------------
func (l *lexema) scanInt(lval *yySymType) int {
	lval.vali = 0

	for {
		r, _, err := l.in.ReadRune()
		if err != nil {
			if err == io.EOF {
				return INT
			}
			return ' '
		}
		if !unicode.IsNumber(r) {
			if r == 'x' { //Hexadecimal
				l.gotRune(r)
				return l.scanHexa(lval)
			}
			if r == '.' {
				l.gotRune(r)
				t := l.scanDecimal(lval)
				lval.valf = float64(lval.vali) + lval.valf
				return t
			}
			if unicode.IsLetter(r) { //DETECT ERROR 10a
				l.gotRune(r)
				Errorf("bad integer '%s' in input", string(l.value))
				return ' '
			}

			l.in.UnreadRune()
			return INT
		}
		n := runeToInt(r)
		lval.vali = n + lval.vali*10.0
		l.gotRune(r) //Para salida de error
	}
}

//SCAN DECIMAL ---------------------------------------------------------------------------
func (l *lexema) scanDecimal(lval *yySymType) int {
	d := 10.0

	r, _, err := l.in.ReadRune()
	if err != nil && err != io.EOF {
		return ' '
	}
	if unicode.IsNumber(r) {
		n := runeToInt(r)
		lval.valf = float64(n) / d
		d *= 10.0

		l.gotRune(r)
		for {
			r, _, err := l.in.ReadRune()
			if err != nil {
				if err == io.EOF {
					return FLOAT
				}
				return ' '
			}
			if !unicode.IsNumber(r) {
				if unicode.IsLetter(r) { //DETECT ERROR 0.10a,0.12.,...
					l.gotRune(r)
					Errorf("bad float '%s' in input", string(l.value))
					return ' '
				}
				l.in.UnreadRune()
				return FLOAT
			}

			n := runeToInt(r)
			lval.valf += float64(n) / d
			d *= 10.0

			l.gotRune(r) //Para la salida de error
		}
	}
	l.in.UnreadRune()
	Errorf("bad input: %s", string(l.value))
	return ' '
}

//SCAN HEXADECIMAL BODY-------------------------------------------------------------------
func (l *lexema) scanHexa(lval *yySymType) int {
	var n int64
	var badinput bool

	n = 0
	if lval.vali != 0 { //Si no es 0x, debe dar error!
		badinput = true
	}

	r, _, err := l.in.ReadRune()
	if err != nil && err != io.EOF {
		return 0
	}
	if unicode.IsNumber(r) || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F') {
		lval.vali = runeToInt(r)

		l.gotRune(r)
		for {
			r, _, err := l.in.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				}
				return ' '
			}
			if !unicode.IsNumber(r) && !('a' <= r && r <= 'f') && !('A' <= r && r <= 'F') {
				if unicode.IsLetter(r) { //DETECT ERROR 0x10h, 0x12.,...
					l.gotRune(r)
					Errorf("bad integer '%s' in input", string(l.value))
					return ' '
				}
				l.in.UnreadRune()
				break
			}
			n = runeToInt(r)
			lval.vali = lval.vali*16 + n

			l.gotRune(r)
		}

		if badinput {
			Errorf("bad input '%s'", string(l.value))
			return ' '
		}
		return INT
	}
	l.in.UnreadRune()
	Errorf("bad input: %s", string(l.value))
	return ' '
}
