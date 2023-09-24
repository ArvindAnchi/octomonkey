package ast

import (
	"octo/token"
	"testing"
)

func TestString(t *testing.T){
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "myVar"},
                    Value: "myVal",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT},
                    Value: "anotherVal",
                },
            },
        },
    }

    if program.String() != "let myVal = anotherVal;" {
        t.Errorf("program.String() wrong. got=%q", program.String())
    }
}

