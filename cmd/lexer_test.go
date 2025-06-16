package main

import "testing"

func TestNextWord(t *testing.T) {
	input := `5 -10 + 1 . 1 - -1 * -6 / dup drop 2 swap over spin . . . 97 emit`

	output := []struct {
		expectedType    WordType
		expectedLiteral string
	}{
		{INT, "5"},
		{INT, "-10"},
		{ADD, "+"},
		{INT, "1"},
		{POP, "."},
		{INT, "1"},
		{SUBTRACT, "-"},
		{INT, "-1"},
		{MULTIPLY, "*"},
		{INT, "-6"},
		{DIVIDE, "/"},
		{DUP, "dup"},
		{DROP, "drop"},
		{INT, "2"},
		{SWAP, "swap"},
		{OVER, "over"},
		{SPIN, "spin"},
		{POP, "."},
		{POP, "."},
		{POP, "."},
		{INT, "97"},
		{EMIT, "emit"},
	}

	l := NewLexer(input)
	for i, tt := range output {
		tok := l.NextWord()
		t.Run("single", func(t *testing.T) {
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		})
	}
}

func TestWordTable(t *testing.T) {
	type expected struct {
		expectedType    WordType
		expectedLiteral string
	}
	type test struct {
		test_name string
		input     string
		output    []expected
	}
	tests := []test{
		{
			test_name: "true false",
			input:     `true false invert`,
			output: []expected{
				{expectedType: TRUE, expectedLiteral: "true"},
				{expectedType: FALSE, expectedLiteral: "false"},
				{expectedType: INVERT, expectedLiteral: "invert"},
			},
		},
		{
			test_name: "mod",
			input:     `5 5 mod`,
			output: []expected{
				{expectedType: INT, expectedLiteral: "5"},
				{expectedType: INT, expectedLiteral: "5"},
				{expectedType: MOD, expectedLiteral: "mod"},
			},
		},
		{
			test_name: "%",
			input:     `5 5 %`,
			output: []expected{
				{expectedType: INT, expectedLiteral: "5"},
				{expectedType: INT, expectedLiteral: "5"},
				{expectedType: MOD, expectedLiteral: "%"},
			},
		},
		{
			test_name: "dup a number",
			input:     `420 dup`,
			output: []expected{
				{expectedType: INT, expectedLiteral: "420"},
				{expectedType: DUP, expectedLiteral: "dup"},
			},
		},
		{
			test_name: "cr cr cr",
			input:     `cr cr cr`,
			output: []expected{
				{CR, "cr"},
				{CR, "cr"},
				{CR, "cr"},
			},
		},
		{
			test_name: "LT and GT",
			input:     `1 2 < -2 > -1 =`,
			output: []expected{
				{INT, "1"},
				{INT, "2"},
				{LT, "<"},
				{INT, "-2"},
				{GT, ">"},
				{INT, "-1"},
				{EQ, "="},
			},
		},
		{
			test_name: "and",
			input:     `10 12 and`,
			output: []expected{
				{INT, "10"},
				{INT, "12"},
				{AND, "and"},
			},
		},
		{
			test_name: "test or with two numbers",
			input:     `10 12 or`,
			output: []expected{
				{INT, "10"},
				{INT, "12"},
				{OR, "or"},
			},
		},
		{
			test_name: "invert: bitwise not",
			input:     `1 invert`,
			output: []expected{
				{INT, "1"},
				{INVERT, "invert"},
			},
		},
	}
	for i, tc := range tests {
		l := NewLexer(tc.input)
		for _, o := range tc.output {
			tok := l.NextWord()
			t.Run(tc.test_name, func(t *testing.T) {
				if tok.Type != o.expectedType {
					t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d", i, o.expectedType, tok.Type)
				}
			})
			t.Run(tc.test_name, func(t *testing.T) {
				if tok.Literal != o.expectedLiteral {
					t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, o.expectedLiteral, tok.Literal)
				}
			})
		}
	}
}
