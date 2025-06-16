package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Interpreter struct {
	l           *Lexer
	currentWord Word
	peekWord    Word
	state       Stack // State of the program is a Stack
	// returnStack []Word
	isCompiling bool
	errors      []string
	udw_table   map[WordName]UDW
}

func NewInterpreter(l *Lexer) *Interpreter {
	intrprtr := &Interpreter{l: l}
	intrprtr.pushNextWord()
	intrprtr.pushNextWord()
	return intrprtr
}

func (i *Interpreter) pushNextWord() {
	i.currentWord = i.peekWord
	i.peekWord = i.l.NextWord()
	i.state.stack = append(i.state.stack, i.currentWord)
}

func (i *Interpreter) ParseProgram() []Word {
	program := i.state.stack

	for i.currentWord.Type != EOF {
		i.Eval(i.state.stack)
		i.pushNextWord()
	}

	return program
}

// DefineWord stores the user's definition of the
// UDW on the udw_table field of the interpreter.
func (i *Interpreter) DefineWord() {
	i.l.readChar() // skip ':', which means 'DEFINE' word
	i.l.skipWhitespace()
	udw_name := i.l.readString()

	var definition []Word
	for i.l.ch != 0x00 {
		word := i.l.NextWord()
		if word.Type == SEMICOLON && word.Literal == ";" {
			break
		}
		definition = append(definition, word)
	}
	w := Word{Name: WordName(udw_name)}
	i.udw_table[w.Name] = definition
}

func (i *Interpreter) Eval(words []Word) error {
	var s Stack // dont need this s

	for _, w := range words {
		switch w.Type {
		case TRUE:
			w.Name = "true"
			w.Literal = "true"
			i.state.Push(w)
		case FALSE:
			w.Name = "false"
			w.Literal = "false"
			i.state.Push(w)
		case AND:
			first := i.state.Pop()
			second := i.state.Pop()
			a, _ := strconv.Atoi(first.Literal)
			b, _ := strconv.Atoi(second.Literal)
			result := a & b
			i.state.Push(result)
		case OR:
			i.state.Push(i.state.Pop() | i.state.Pop())
		case INVERT:
			i.state.Push(^i.state.Pop())
		case EQ:
			if i.state.Pop() == i.state.Pop() {
				i.state.Push(int(TRUE))
			} else {
				i.state.Push(int(FALSE))
			}
		case LT:
			v1 := i.state.Pop()
			v2 := i.state.Pop()
			if v2 < v1 {
				i.state.Push(int(TRUE))
			} else {
				i.state.Push(int(FALSE))
			}
		case GT:
			v1 := i.state.Pop()
			v2 := i.state.Pop()
			if v2 > v1 {
				i.state.Push(int(TRUE))
			} else {
				i.state.Push(int(FALSE))
			}
		case ADD:
			i.state.Push(i.state.Pop() + i.state.Pop())
		case SUBTRACT:
			i.state.Push(i.state.Pop() - i.state.Pop())
		case MULTIPLY:
			i.state.Push(i.state.Pop() * i.state.Pop())
		case DIVIDE:
			i.state.Push(i.state.Pop() / i.state.Pop())
		case MOD:
			f := i.state.Pop()
			sec := i.state.Pop()
			i.state.Push(sec % f)
		case POP:
			top := i.state.Pop()
			fmt.Println(top)
		case DUP:
			top := i.state.Top()
			i.state.Push(top)
		case DROP:
			i.state.Pop()
		case SWAP:
			first := i.state.Pop()
			second := i.state.Pop()
			i.state.Push(first)
			i.state.Push(second)
		case OVER:
			sec := i.state.Second()
			i.state.Push(sec)
		case SPIN:
			n1, n2, n3 := i.state.Pop(), i.state.Pop(), i.state.Pop()
			i.state.Push(n2)
			i.state.Push(n3)
			i.state.Push(n1)
		case EMIT:
			w := i.state.Pop()
			fmt.Println(string(rune(w.Literal)))
		case CR:
			fmt.Println()
		case EOF:
			fmt.Println()
		case ILLEGAL:
			fmt.Printf("%x of type %d is illegal.\n", w.Literal, w.Type)
		case INT:
			v, e := strconv.Atoi(w.Literal)
			if e != nil {
				log.Fatal(e)
			}
			s.Push(v)
		default:
			log.Fatalf("reached default %s (%T) has type %v\n", w.Literal, w.Literal, w.Type)
			return errors.New("reached eval default") //  s.stack, errors.New("reached default.")
		}
	}

	return nil
}
