package main

import (
	"testing"
)

func TestParseNextWord(t *testing.T) {
	// item := "1 1 +"

}

// func TestEvalTable(t *testing.T) {
// 	type expected struct {
// 		expectedName    WordName
// 		expectedType    WordType
// 		expectedLiteral string
// 		expectedStk     []int
// 	}
// 	type test struct {
// 		test_name string
// 		input     string
// 		output    []expected
// 	}
// 	tests := []test{
// 		{
// 			test_name: "EQ",
// 			input:     `8 8 = 4 =`,
// 			output: []expected{
// 				{"8", INT, "8", []int{8}},
// 				{"8", INT, "8", []int{8, 8}},
// 				{"=", EQ, "=", []int{-1}},
// 				{"4", INT, "4", []int{-1, 4}},
// 				{"=", EQ, "=", []int{0}},
// 			},
// 		},
// 	}

// 	for _, tc := range tests {
// 		l := NewLexer(tc.input)
// 		interpreter := NewInterpreter(l)

// 		for range tc.output {
// 			interpreter.ParseNextWord()

// 			interpreter.Eval(interpreter.dataStack)

// 			t.Run(tc.test_name, func(t *testing.T) {
// 				// if w.Type != o.expectedType {
// 				// 	t.Fatalf("tests[%d, %d] - tokentype wrong. expected=%v, got=%v", i, n, o.expectedType, w.Type)
// 				// }
// 			})
// 			t.Run(tc.test_name, func(t *testing.T) {
// 				// if w.Literal != o.expectedLiteral {
// 				// 	t.Fatalf("tests[%d, %d] - literal wrong. expected=%q, got=%q", i, n, o.expectedLiteral, w.Literal)
// 				// }
// 			})
// 			t.Run(tc.test_name, func(t *testing.T) {
// 				// if !slices.Equal(o.expectedStk, got) {
// 				// 	t.Fatalf("wrong evaluation. expected=%v, got=%v", o.expectedStk, got)
// 				// }
// 			})
// 		}
// 	}
// }
