package regexp_syntax

import (
	"fmt"
	"log"
	"regexp/syntax"
	"strings"
	"unicode"
)

// dump prints a string representation of the regexp showing
// the structure explicitly.
func rewriteRegexp(re *syntax.Regexp) string {
	var b strings.Builder
	doRewriteRegexp(&b, re)
	return b.String()
}

var opValues = []string{
	syntax.OpNoMatch:        "no",
	syntax.OpEmptyMatch:     "emp",
	syntax.OpLiteral:        "lit",
	syntax.OpCharClass:      "cc",
	syntax.OpAnyCharNotNL:   ".",
	syntax.OpAnyChar:        ".",
	syntax.OpBeginLine:      "^",
	syntax.OpEndLine:        "$",
	syntax.OpBeginText:      "^",
	syntax.OpEndText:        "$",
	syntax.OpWordBoundary:   "wb",
	syntax.OpNoWordBoundary: "nwb",
	syntax.OpCapture:        "cap",
	syntax.OpStar:           "*",
	syntax.OpPlus:           "+",
	syntax.OpQuest:          "?",
	syntax.OpRepeat:         "*",
	syntax.OpConcat:         "<nil>",
	syntax.OpAlternate:      "alt",
}

// dumpRegexp writes an encoding of the syntax tree for the regexp re to b.
// It is used during testing to distinguish between parses that might print
// the same using re's String method.
func doRewriteRegexp(b *strings.Builder, re *syntax.Regexp) {
	if int(re.Op) >= len(opValues) || opValues[re.Op] == "" {
		fmt.Fprintf(b, "op%d", re.Op)
	} else {
		switch re.Op {
		default:
			if n := opValues[re.Op]; n != "<nil>" {
				b.WriteString(opValues[re.Op])
			}
		case syntax.OpStar, syntax.OpPlus, syntax.OpQuest, syntax.OpRepeat:
			if re.Flags&syntax.NonGreedy != 0 {
				b.WriteByte('n')
			}
			b.WriteString(opValues[re.Op])
		case syntax.OpLiteral:
			if re.Flags&syntax.FoldCase != 0 {
				for _, r := range re.Rune {
					if unicode.SimpleFold(r) != r {
						b.WriteString("fold")
						break
					}
				}
			}
		}
	}

	switch re.Op {
	case syntax.OpEndText:
		if re.Flags&syntax.WasDollar == 0 {
			b.WriteString(`\z`)
		}
	case syntax.OpLiteral:
		for _, r := range re.Rune {
			if string(r) == "." {
				b.WriteRune('\\')
			}
			b.WriteRune(r)
		}
	case syntax.OpConcat, syntax.OpAlternate:
		for _, sub := range re.Sub {
			doRewriteRegexp(b, sub)
		}
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest:
		doRewriteRegexp(b, re.Sub[0])
	case syntax.OpRepeat:
		fmt.Fprintf(b, "%d,%d ", re.Min, re.Max)
		doRewriteRegexp(b, re.Sub[0])
	case syntax.OpCapture:
		if re.Name != "" {
			b.WriteString(re.Name)
			b.WriteByte(':')
		}
		doRewriteRegexp(b, re.Sub[0])
	case syntax.OpCharClass:
		sep := ""
		for i := 0; i < len(re.Rune); i += 2 {
			b.WriteString(sep)
			sep = " "
			lo, hi := re.Rune[i], re.Rune[i+1]
			if lo == hi {
				fmt.Fprintf(b, "%#x", lo)
			} else {
				fmt.Fprintf(b, "%#x-%#x", lo, hi)
			}
		}
	}
}

func Correct(val string) string {
	fmt.Println(val)

	r, err := syntax.Parse(val, syntax.Perl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	// s := r.String()
	//
	// if len(s) > 2 && s[:2] == `\A` {
	//
	// s = "^" + s[2:]
	// }
	//
	// fmt.Println(r.String())
	//
	// dollarSignOperator := "(?-m:$)"
	// fmt.Println(s[len(s)-len(dollarSignOperator):])
	// if len(s) > len(dollarSignOperator) && s[len(s)-len(dollarSignOperator):] == dollarSignOperator {
	// s = s[:len(s)-len(dollarSignOperator)] + "?"
	// }
	//
	// dotOperator := "(?-s:.)"

	return rewriteRegexp(r)
}
