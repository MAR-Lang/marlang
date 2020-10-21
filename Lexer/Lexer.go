package Lexer

import (
	. "../Token"
)

type Lexer struct {
	input string
	rPos  int
	ch    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) ReadCh() byte {
	if l.rPos < len(l.input) {
		l.ch = l.input[l.rPos]
		l.rPos++
	} else {
		l.ch = 0
	}
	return l.ch
}

func (l *Lexer) Lex() []*Token {
	var tokens []*Token
	for {
		switch l.ReadCh() {
		case ' ', '\n', '\t', 'v', '\r':
			continue
		case '+':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: ADDa})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: ADD})
			}
		case '-':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: SUBa})
			} else if l.ch == '>' {
				tokens = append(tokens, &Token{Type: THEN})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: SUB})
			}
		case '*':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: MULa})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: MUL})
			}
		case '/':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: DIVa})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: DIV})
			}
		case '%':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: MODa})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: MOD})
			}
		case '^':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: POWa})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: POW})
			}
		case '=':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: EQU})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: ASS})
			}
		case '>':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: LARe})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: LAR})
			}
		case '<':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: LESe})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: LES})
			}
		case '!':
			if l.ReadCh() == '=' {
				tokens = append(tokens, &Token{Type: NOTe})
			} else {
				l.rPos--
				tokens = append(tokens, &Token{Type: NOT})
			}
		case '&':
			if l.ReadCh() == '&' {
				tokens = append(tokens, &Token{Type: AND})
			} else {
				l.rPos--
				//tokens = append(tokens, &Token{Type: BitAND})
			}
		case '|':
			if l.ReadCh() == '|' {
				tokens = append(tokens, &Token{Type: OR})
			} else {
				l.rPos--
				//tokens = append(tokens, &Token{Type: BitOR})
			}
		case '(':
			tokens = append(tokens, &Token{Type: LPAREN})
		case ')':
			tokens = append(tokens, &Token{Type: RPAREN})
		case '{':
			tokens = append(tokens, &Token{Type: LBRACE})
		case '}':
			tokens = append(tokens, &Token{Type: RBRACE})
		case ':':
			tokens = append(tokens, &Token{Type: COLON})
		case '?':
			tokens = append(tokens, &Token{Type: IF})
		case 0:
			goto end
		default:
			if IsLetter(l.ch) {
				tokens = append(tokens, l.ReadIdent())
			} else if IsDigit(l.ch) {
				tokens = append(tokens, l.ReadNum())
			}
		}
	}
end:
	return tokens
}

func (l *Lexer) ReadNum() *Token {
	t := new(Token)
	retPos := l.rPos
	t.Type = INT
loop:
	for IsDigit(l.ch) {
		t.Val += string(l.ch)
		retPos = l.rPos
		l.ReadCh()
	}
	if l.ch == '.' {
		if IsDigit(l.ReadCh()) && t.Type != FLOAT {
			t.Type = FLOAT
			t.Val += "."
			goto loop
		} else {
			t.Type = ILLEGAL
		}
	}
	l.rPos = retPos
	return t
}

func (l *Lexer) ReadIdent() *Token {
	t := new(Token)
	retPos := l.rPos
	t.Type = IDENT
	for IsLetter(l.ch) || IsDigit(l.ch) {
		t.Val += string(l.ch)
		retPos = l.rPos
		l.ReadCh()
	}
	l.rPos = retPos
	return t
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsLetter(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch == '_'
}
