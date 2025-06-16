package main

type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' || l.ch == 0xd || l.ch == 0xa {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?'
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) NextWord() (word Word) {
	l.skipWhitespace()

	switch l.ch {
	case '-':
		p := l.peekChar()
		if isDigit(p) {
			word.Type = INT
			l.readChar()
			word.Literal = "-" + l.readNumber()
			return word
		} else {
			w := WordName(string(l.ch))
			wt := GetWordType(w)
			word = NewWord(w, wt, string(w))
		}
	case ';', '.', '+', '*', '/', '%', '<', '>', '=':
		w := WordName(string(l.ch))
		wt := GetWordType(w)
		word = NewWord(w, wt, string(w))
	case ':':
		w := WordName("define") // WordName(string(l.ch))
		wt := DEFINE            // GetWordType(w, nil)
		word = NewWord(w, wt, string(w))
	case 0x00:
		word = NewWord("end_of_file", EOF, "0x00")
	default:
		if isLetter(l.ch) {
			str := l.readString()
			w := WordName(str)
			wt := GetWordType(w)
			word = NewWord(w, wt, str)
		} else if isDigit(l.ch) {
			word.Type = INT
			word.Literal = l.readNumber()
			word.Name = WordName(word.Literal)
			return word
		} else {
			word = NewWord("illegal", ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return word
}
