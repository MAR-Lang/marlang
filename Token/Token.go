package Token

type (
	Token struct {
		Type int
		Val  string
	}
)

const (
	EOF     = iota // 0
	ILLEGAL        // 23..324
	IDENT          // a123, name ...
	FLOAT          // 1.432, 554.32553 ...
	INT            // 2832, 593, 20, 3 ...

	ASS  // =
	ADD  // +
	SUB  // -
	MUL  // *
	DIV  // /
	MOD  // %
	POW  // ^
	ADDa // +=
	SUBa // -=
	MULa // *=
	DIVa // /=
	MODa // %=
	POWa // ^=

	EQU  // ==
	LAR  // >
	LES  // <
	NOT  // !
	LARe // >=
	LESe // <=
	NOTe // !=
	AND  // &&
	OR   // ||

	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	COLON     // :
	SEMICOLON // ;
	COMMA     // ,

	THEN // ->
	IF   // ?
)
