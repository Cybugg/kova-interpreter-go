package lexer

type Lexer struct {
inputstring
positionint
readPosition int
ch
byte 
}

func New(input string) *Lexer {
l := &Lexer{input: input}
return l
}