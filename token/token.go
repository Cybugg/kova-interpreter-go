package token

type tokenType string

type token struct {
	type tokenType
	Value string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" 
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)