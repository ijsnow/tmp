package main

import (
	"fmt"
	"regexp/syntax"
)

/****

FoldCase      Flags = 1 << iota // case-insensitive match
Literal                         // treat pattern as literal string
ClassNL                         // allow character classes like [^a-z] and [[:space:]] to match newline
DotNL                           // allow . to match newline
OneLine                         // treat ^ and $ as only matching at beginning and end of text
NonGreedy                       // make repetition operators default to non-greedy
PerlX                           // allow Perl extensions
UnicodeGroups                   // allow \p{Han}, \P{Han} for Unicode group and negation
WasDollar                       // regexp OpEndText was $, not \z
Simple                          // regexp contains no counted repetition

MatchNL = ClassNL | DotNL

Perl        = ClassNL | OneLine | PerlX | UnicodeGroups // as close to Perl as possible
POSIX Flags = 0                                         // POSIX syntax

*/

func main() {
	flags := syntax.ClassNL | syntax.UnicodeGroups

	re, _ := syntax.Parse("^f.o$", syntax.Perl)
	fmt.Println(re.String())

	re, err := syntax.Parse("(|)^$.[*+?]{5,10},", flags)
	re.Simplify()
	fmt.Println(re, err)
}
