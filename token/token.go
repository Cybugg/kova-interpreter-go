package token

type tokenType string

type token struct {
	type tokenType
	Value string
}