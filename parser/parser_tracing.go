package parser

import (
	"fmt"
	"strings"
)

var traceL int = 0

const traceIndentPH string = "\t"

func indentLevel() string {
	return strings.Repeat(traceIndentPH, traceL-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentLevel(), fs)
}

func incIdent() { traceL = traceL + 1 }
func decIdent() { traceL = traceL - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
