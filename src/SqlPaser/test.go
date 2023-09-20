package SqlPaser

import (
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "SELECT column1, column2 FROM table;",
			expected: []Token{
				{SELECT, "SELECT"},
				{IDENTIFIER, "COLUMN1"},
				{COMMA, ","},
				{IDENTIFIER, "COLUMN2"},
				{FROM, "FROM"},
				{IDENTIFIER, "TABLE"},
				{SEMICOLON, ";"},
				{EOF, ""},
			},
		},
		// ... 更多测试用例
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		tokens := []Token{}
		for {
			token := l.NextToken()
			tokens = append(tokens, token)
			if token.Type == EOF {
				break
			}
		}
		if !reflect.DeepEqual(tokens, tt.expected) {
			t.Errorf("For input %q, expected %v but got %v", tt.input, tt.expected, tokens)
		}
	}
}
