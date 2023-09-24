package lexer

import "octo/token"

type Lexer struct {
	input        string
	charPosition int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}

	lexer.readChar()

	return lexer
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func isLetter(input byte) bool {
	return 'a' <= input && input >= 'z' || 'A' <= input && input >= 'Z' || input == '_'
}

func isFloat(input byte) bool {
	return ('0' <= input && input <= '9') || input == '.'
}

func isDigit(input byte) bool {
	return '0' <= input && input <= '9'
}

func (lexer *Lexer) readIdent() string {
	position := lexer.charPosition

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.charPosition]
}

func (lexer *Lexer) readNumberOrDot() string {
	position := lexer.charPosition

	for isDigit(lexer.ch) || lexer.ch == '.' {
		lexer.readChar()
	}

	return lexer.input[position:lexer.charPosition]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.charPosition

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.charPosition]
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.charPosition = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var nextToken token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '!':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			nextToken = token.Token{Type: token.N_EQ, Literal: "!="}
		} else {
			nextToken = token.Token{Type: token.BANG, Literal: string(lexer.ch)}
		}
	case '=':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			nextToken = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			nextToken = token.Token{Type: token.ASSIGN, Literal: string(lexer.ch)}
		}
	case '+':
		nextToken = token.Token{Type: token.PLUS, Literal: string(lexer.ch)}
	case '-':
		nextToken = token.Token{Type: token.MINUS, Literal: string(lexer.ch)}
	case '*':
		nextToken = token.Token{Type: token.STAR, Literal: string(lexer.ch)}
	case '/':
		nextToken = token.Token{Type: token.SLASH, Literal: string(lexer.ch)}
	case ',':
		nextToken = token.Token{Type: token.COMMA, Literal: string(lexer.ch)}
	case ';':
		nextToken = token.Token{Type: token.SEMICOLON, Literal: string(lexer.ch)}
	case '(':
		nextToken = token.Token{Type: token.L_PAREN, Literal: string(lexer.ch)}
	case ')':
		nextToken = token.Token{Type: token.R_PAREN, Literal: string(lexer.ch)}
	case '{':
		nextToken = token.Token{Type: token.L_BRACE, Literal: string(lexer.ch)}
	case '}':
		nextToken = token.Token{Type: token.R_BRACE, Literal: string(lexer.ch)}
	case '<':
		nextToken = token.Token{Type: token.L_THAN, Literal: string(lexer.ch)}
	case '>':
		nextToken = token.Token{Type: token.G_THAN, Literal: string(lexer.ch)}
	case 0:
		nextToken = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isDigit(lexer.ch) {
			nextToken.Literal = lexer.readNumber()
			nextToken.Type = token.INT

			return nextToken
        } else if isFloat(lexer.ch) {
			nextToken.Literal = lexer.readNumberOrDot()
			nextToken.Type = token.FLOAT

			return nextToken
		} else if isLetter(lexer.ch) {
			nextToken.Literal = lexer.readIdent()
			nextToken.Type = token.LookupIdent(nextToken.Literal)

			return nextToken
		} else {
			nextToken = token.Token{Type: token.ILLEGAL, Literal: string(lexer.ch)}
		}
	}

	lexer.readChar()

	return nextToken
}
