package lexer

import (
	"fmt"
	"octo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
    let ten = 10;
    
    let add = fn(x, y) {
      x + y;
    };
    
    let result = add(five, ten);
    !-/*5;
    5 < 10 > 5;
    
    if (5 < 10) {
        return true;
    } else {
        return false;
    }
    10 == 10;
    10 != 9;
    `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.L_PAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.L_PAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.L_THAN, "<"},
		{token.INT, "10"},
		{token.G_THAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.L_PAREN, "("},
		{token.INT, "5"},
		{token.L_THAN, "<"},
		{token.INT, "10"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.ELSE, "else"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.R_BRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.N_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, testToken := range tests {
		tok := l.NextToken()

		fmt.Printf("test[%d] - tokentype got=%q expected=%q\n", i, tok.Type, testToken.expectedType)
		fmt.Printf("test[%d] - tokenLiteral got=%q expected=%q\n", i, tok.Literal, testToken.expectedLiteral)

		if tok.Type != testToken.expectedType {
			t.Fatalf("test[%d] - tokentype got=%q expected=%q\n", i, tok.Type, testToken.expectedType)
		}

		if tok.Literal != testToken.expectedLiteral {
			t.Fatalf("test[%d] - tokenLiteral got=%q expected=%q\n", i, tok.Literal, testToken.expectedLiteral)
		}
	}
}
