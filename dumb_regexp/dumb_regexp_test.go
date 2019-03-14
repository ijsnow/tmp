package dumb_regexp

import "testing"

func TestFix(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"$foo[", `\$foo\[`},
		// {"foo(", `foo\(`},
		// {"foo[", `foo\[`},
		// {"*foo", `\*foo`},
		{"$foo", `\$foo`},
		{`foo\s=\s$bar`, `foo\s=\s\$bar`},
		// {"foo)", `foo\)`},
		// {"foo]", `foo\]`},

		// Valid regexps
		{"^f.o$", "^f.o$"},
		{"$foo", `\$foo`},
		// {`foo\(`, `foo\(`},
		// {`foo\[`, `foo\[`},
		// {`\*foo`, `\*foo`},
		{`\$foo`, `\$foo`},
		{`foo$`, `foo$`},
		{`foo\s=\s\$bar`, `foo\s=\s\$bar`},
		// {"[$]", `[$]`}, // Do we support?
	}

	for _, test := range tests {
		got := Fix(test.input)

		if got != test.output {
			t.Errorf("input tranformed in an unexpected way\ngot: %v\nwant: %v\n", got, test.output)
		}
	}
}
