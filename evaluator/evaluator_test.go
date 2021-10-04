package evaluator

import (
	"testing"

	"github.com/jarviliam/inti/lexer"
	"github.com/jarviliam/inti/object"
	"github.com/jarviliam/inti/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{input: "5", expected: 5},
		{input: "10", expected: 10},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testIntegerObject(t, evaluated, tC.expected)
	}
}
func TestEvalBooleanExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "true", expected: true},
		{input: "false", expected: false},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testBooleanObject(t, evaluated, tC.expected)
	}
}

func TestBangOp(t *testing.T) {
	tests := []struct {
		in  string
		exp bool
	}{
		{"!true", false},
		{"!false", true},
	}
	for _, tc := range tests {
		evaluated := testEval(tc.in)
		testBooleanObject(t, evaluated, tc.exp)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expectd int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not integer. got:%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expectd {
		t.Errorf("object has wrong value. got:%d want:%d", result.Value, expectd)
		return false
	}
	return true
}
func testBooleanObject(t *testing.T, obj object.Object, expectd bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not boolean. got:%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expectd {
		t.Errorf("object has wrong value. got:%t want:%t", result.Value, expectd)
		return false
	}
	return true
}
