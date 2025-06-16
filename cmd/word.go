package main

type WordType int
type WordName string
type WordLiteral string

type Word struct {
	Name    WordName
	Type    WordType
	Literal string
}

// A UDW is a User Defined Word
type UDW []Word

// A Definition is the definition of the function.
// It can be predefined or user-defined
// Predefined ones are defined by the interpreter.
type Definition struct {
	preDefined func(*Interpreter) error
	Words      UDW
}

const (
	// Boolean Operations
	TRUE  WordType = -1
	FALSE WordType = 0
	EQ    WordType = iota
	NOTEQ
	LT
	GT
	AND
	OR
	INVERT // 8

	// Stack
	INT
	POP
	DUP
	DROP
	SWAP
	OVER
	SPIN
	EMIT
	CR // 17

	// Math Operations
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	MOD // 22

	// Conditionals
	IF
	ELSE
	THEN // 25

	// UDF
	DEFINE // colon ':'

	// extra
	SEMICOLON //
	NEWLINE
	EOF
	VAR
	CONST
	ILLEGAL //
)

var PreDefinedWords = map[string]WordType{
	"+":      ADD,
	"*":      MULTIPLY,
	"-":      SUBTRACT,
	"/":      DIVIDE,
	".":      POP,
	"%":      MOD,
	"mod":    MOD,
	"dup":    DUP,
	"drop":   DROP,
	"swap":   SWAP,
	"over":   OVER,
	"spin":   SPIN,
	"emit":   EMIT,
	"cr":     CR,
	"true":   TRUE,
	"false":  FALSE,
	"=":      EQ,
	"<":      LT,
	">":      GT,
	"!=":     NOTEQ,
	"and":    AND,
	"or":     OR,
	"invert": INVERT,
	":":      DEFINE,
	";":      SEMICOLON,
	"var":    VAR,
	"const":  CONST,
	"if":     IF,
	"else":   ELSE,
	"then":   THEN,
}

func NewWord(n WordName, wt WordType, literal string) Word {
	return Word{Name: n, Type: wt, Literal: literal}
}

func LookupWordType(n WordName, udw_table map[Word]UDW) WordType {
	if wT, ok := PreDefinedWords[string(n)]; ok {
		return wT
	} else if _, ok := udw_table[Word{Name: n, Type: DEFINE, Literal: string(n)}]; ok {
		return DEFINE
	}
	return ILLEGAL
}

func GetWordType(n WordName) WordType {
	if wT, ok := PreDefinedWords[string(n)]; ok {
		return wT
	}
	return ILLEGAL
}
