package PythonSqlPaser

var tokens map[string]interface{}

// TokenType represents a token type
type TokenType struct {
	parent *TokenType
	name   string
}

// String returns the string representation of the TokenType
func (t *TokenType) String() string {
	if t.parent == nil {
		return t.name
	}
	return t.parent.String() + "." + t.name
}

// NewTokenType creates a new TokenType with the given parent and name
func NewTokenType(parent *TokenType, name string) *TokenType {
	return &TokenType{parent: parent, name: name}
}

// Define the root Token
var Token = &TokenType{parent: nil, name: "Token"}

// Define special token types
var (
	Text        = NewTokenType(Token, "Text")
	Whitespace  = NewTokenType(Text, "Whitespace")
	Newline     = NewTokenType(Whitespace, "Newline")
	Error       = NewTokenType(Token, "Error")
	Other       = NewTokenType(Token, "Other")
	Keyword     = NewTokenType(Token, "Keyword")
	Name        = NewTokenType(Token, "Name")
	Literal     = NewTokenType(Token, "Literal")
	String      = NewTokenType(Literal, "String")
	Number      = NewTokenType(Literal, "Number")
	Punctuation = NewTokenType(Token, "Punctuation")
	Operator    = NewTokenType(Token, "Operator")
	Comparison  = NewTokenType(Operator, "Comparison")
	Wildcard    = NewTokenType(Token, "Wildcard")
	Comment     = NewTokenType(Token, "Comment")
	Assignment  = NewTokenType(Token, "Assignment")
	Generic     = NewTokenType(Token, "Generic")
	Command     = NewTokenType(Generic, "Command")
	DML         = NewTokenType(Keyword, "DML")
	DDL         = NewTokenType(Keyword, "DDL")
	CTE         = NewTokenType(Keyword, "CTE")
	// 定义额外的 token 类型
	CommentSingle      = NewTokenType(Comment, "Single")
	Multiline          = NewTokenType(Comment, "Multiline")
	SingleHint         = NewTokenType(CommentSingle, "Hint")
	MultilineHint      = NewTokenType(Multiline, "Hint")
	Placeholder        = NewTokenType(Name, "Placeholder")
	Hexadecimal        = NewTokenType(Number, "Hexadecimal")
	Float              = NewTokenType(Number, "Float")
	Integer            = NewTokenType(Number, "Integer")
	Single             = NewTokenType(String, "Single")
	Symbol             = NewTokenType(String, "Symbol")
	TZCast             = NewTokenType(Keyword, "TZCast")
	Builtin            = NewTokenType(Name, "Builtin")
	PROCESS_AS_KEYWORD = NewTokenType(nil, "PROCESS_AS_KEYWORD")
	Order              = NewTokenType(Keyword, "Order")
)
