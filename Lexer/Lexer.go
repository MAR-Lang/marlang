package Lexer

import (
	"fmt"
	. "../Token"
)

type Lexer struct {
	input string
	position  int
	char  byte
}

func New(input string) *Lexer {
	this := &Lexer{input: input}
	return this
}

func (this *Lexer) Read() byte {
	if this.position < len(this.input) {
		this.char = this.input[this.position]
		this.position++
	} else {
		this.char = 0
	}
	return this.char
}

func (this *Lexer) Lex() []*Token {
	var tokens []*Token
	for {
		switch this.Read() {
		case ' ', '\n', '\t', '\v', '\r':
			continue
		case '+':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: ADDa})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: ADD})
			}
		case '-':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: SUBa})
			} else if this.char == '>' {
				tokens = append(tokens, &Token{Type: THEN})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: SUB})
			}
		case '*':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: MULa})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: MUL})
			}
		case '/':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: DIVa})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: DIV})
			}
		case '%':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: MODa})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: MOD})
			}
		case '^':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: POWa})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: POW})
			}
		case '=':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: EQU})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: ASS}) // ♂ stick finger in my ASS ♂
			}
		case '>':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: LARe})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: LAR})
			}
		case '<':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: LESe})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: LES})
			}
		case '!':
			if this.Read() == '=' {
				tokens = append(tokens, &Token{Type: NOTe})
			} else {
				this.position--
				tokens = append(tokens, &Token{Type: NOT})
			}
		case '&':
			if this.Read() == '&' {
				tokens = append(tokens, &Token{Type: AND})
			} else {
				this.position--
				//tokens = append(tokens, &Token{Type: BitAND})
			}
		case '|':
			if this.Read() == '|' {
				tokens = append(tokens, &Token{Type: OR})
			} else {
				this.position--
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
		case '\'', '"':
			tokens = append(tokens, this.ReadString(this.char));
		case 0:
			return tokens
		default:
			if IsLetter(this.char) {
				tokens = append(tokens, this.ReadIdentifier())
			} else if IsDigit(this.char) {
				tokens = append(tokens, this.ReadNumber())
			}
		}
	}
}

func (this *Lexer) ReadNumber() *Token {
	token := &Token{Type: INT}
	retPos := this.position
loop:
	for IsDigit(this.char) {
		token.Value += string(this.char)
		retPos = this.position
		this.Read()
	}
	if this.char == '.' {
		if IsDigit(this.Read()) && token.Type != FLOAT {
			token.Type = FLOAT
			token.Value += "."
			goto loop
		} else {
			token.Type = ILLEGAL
		}
	}
	this.position = retPos
	return token
}

func (this *Lexer) ReadIdentifier() *Token {
	token := &Token{Type: IDENT}
	retPos := this.position
	for IsLetter(this.char) || IsDigit(this.char) {
		token.Value += string(this.char)
		retPos = this.position
		this.Read()
	}
	this.position = retPos
	return token
}

func (this *Lexer) ReadString(start_char byte) *Token {
	token := &Token{Type: STRING}
	for this.Read() != start_char {
		if this.char == 0 {
			fmt.Println("Unexcepted EOF")
			break
		}
		token.Value += string(this.char)
	}
	return token
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsLetter(char byte) bool {
	return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char == '_'
}
