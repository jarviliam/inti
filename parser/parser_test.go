package parser

import (
	"fmt"
	"testing"

	"github.com/jarviliam/inti/ast"
	"github.com/jarviliam/inti/lexer"
)

func TestLetStatements(t *testing.T) {
	in := `let x = 5; let y = 5; let foo = 989898;`

	l := lexer.New(in)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)

	if program == nil {
		t.Fatalf("ParseProgram")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("statements dont equal; want : 3; got : %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, ts := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, ts.expectedIdent) {
			return
		}
	}

}

func TestReturnStatements(t *testing.T) {
	in := `
    return 5;
    return 10;
    return 911;
  `

	l := lexer.New(in)
	p := New(l)

	prog := p.ParseProgram()
	checkParserError(t, p)

	if len(prog.Statements) != 3 {
		t.Fatalf("statements dont equal; want : 3; got : %d", len(prog.Statements))
	}

	for _, st := range prog.Statements {

		returnSt, ok := st.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Expected returnStatement. Got %T", st)
			continue
		}
		if returnSt.TokenLiteral() != "return" {
			t.Errorf("returnStmt token literal; got : %q", returnSt.TokenLiteral())
		}
	}

}

func TestIdentifierExpression(t *testing.T) {
	in := "foobar;"
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("statements dont equal; want : 1; got : %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program Statement is not an ast.ExpressionStatement; got: %T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.value not %s; got %s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("tokenliteral not %s, got %s", "foobar", ident.TokenLiteral())
	}

}

func TestLiteralExpressions(t *testing.T) {
	in := "5;"
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("statements dont equal; want : 1; got : %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program Statement is not an ast.ExpressionStatement; got: %T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("ident.value not %d; got %d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("tokenliteral not %s, got %s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTest := []struct {
		in   string
		op   string
		intv int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTest {
		l := lexer.New(tt.in)
		p := New(l)
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("statements dont equal; want : 1; got : %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program Statement is not an ast.ExpressionStatement; got: %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.op {
			t.Fatalf("exp.Op is not %s; got %s", tt.op, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.intv) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("literal is not let. got:%q", s.TokenLiteral())
		return false
	}

	letSt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letSt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. got %s", name, letSt.Name.Value)
		return false
	}

	if letSt.Name.TokenLiteral() != name {
		t.Errorf("token lit not %s; got %s", name, letSt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []struct {
		in    string
		left  int64
		op    string
		right int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range tests {
		l := lexer.New(tt.in)
		p := New(l)
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("statements dont equal; want : 1; got : %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program Statement is not an ast.ExpressionStatement; got: %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp not *ast.InfixExpression. got=%T", stmt.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, tt.left) {
			return
		}
		if exp.Operator != tt.op {
			t.Fatalf("exp.Op is not %s; got %s", tt.op, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.right) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserError(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected %q; got %q", tt.expected, actual)
		}
	}
}

func checkParserError(t *testing.T, p *Parser) {
	err := p.Errors()
	if len(err) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(err))

	for _, m := range err {
		t.Errorf("error : %q", m)
	}
	t.FailNow()
}
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not integerLit got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d; got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d; got %s", value, integ.TokenLiteral())
		return false
	}
	return true
}
