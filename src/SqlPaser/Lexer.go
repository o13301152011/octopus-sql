package SqlPaser

import (
	"strings"
	"unicode"
)

type Lexer struct {
	input  string  // 输入字符串
	pos    int     // 当前位置
	tokens []Token // 已解析的令牌列表
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace() // 跳过空白

	if l.pos >= len(l.input) {
		return Token{EOF, ""}
	}

	switch ch := l.input[l.pos]; {
	case unicode.IsLetter(rune(ch)):
		token := l.lexKeywordOrIdentifier()
		if l.input[l.pos] == '(' {
			token.Type = FUNCTION
		}
		return token
	case unicode.IsDigit(rune(ch)):
		return l.lexNumber()
	case ch == ',':
		l.pos++
		return Token{COMMA, ","}
	case ch == ';':
		l.pos++
		return Token{SEMICOLON, ";"}
	case ch == '(':
		l.pos++
		return Token{LEFT_PAREN, "("}
	case ch == ')':
		l.pos++
		return Token{RIGHT_PAREN, ")"}
	case ch == '=': // 处理字符串值
		l.pos++
		return Token{EQUALS, "="}
	case ch == '<':
		l.pos++
		if l.input[l.pos] == '=' {
			l.pos++
			return Token{LESS_EQUALS, "<="}
		} else if l.input[l.pos] == '>' {
			l.pos++
			return Token{NOT_EQUALS, "<>"}
		}
		return Token{LESS_THAN, "<"}
	case ch == '>':
		l.pos++
		if l.input[l.pos] == '=' {
			l.pos++
			return Token{GREATER_EQUALS, ">="}
		}
		return Token{GREATER_THAN, ">"}
	case ch == '!':
		l.pos++
		if l.input[l.pos] == '=' {
			l.pos++
			return Token{NOT_EQUALS, "!="}
		}
	case ch == '\'': // 处理字符串值
		return l.lexString()
	default:
		l.pos++
		return Token{EOF, string(ch)}
	}
	return Token{EOF, ""}
}

// 解析关键字或标识符
func (l *Lexer) lexKeywordOrIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) && (unicode.IsLetter(rune(l.input[l.pos])) || unicode.IsDigit(rune(l.input[l.pos])) || l.input[l.pos] == '_' || l.input[l.pos] == '.') {
		l.pos++
	}
	word := strings.ToUpper(l.input[start:l.pos])

	// 检查单词关键字
	if keyword, ok := keywords[word]; ok {
		return Token{keyword, word}
	}
	return Token{IDENTIFIER, l.input[start:l.pos]}
}

func (l *Lexer) lexString() Token {
	pos := l.pos + 1 // 跳过开头的单引号
	for {
		if pos >= len(l.input) || l.input[pos] == '\'' {
			break
		}
		pos++
	}
	val := l.input[l.pos+1 : pos] // 获取字符串值，不包括单引号
	l.pos = pos + 1               // 更新位置
	return Token{STRING, val}
}

// 解析数字
func (l *Lexer) lexNumber() Token {
	start := l.pos
	for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
		l.pos++
	}
	return Token{NUMBER, l.input[start:l.pos]}
}

// 跳过空白
func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}
}

// 向前预览字符
func (l *Lexer) peek(dist int) byte {
	pos := l.pos + dist
	if pos >= len(l.input) {
		return 0
	}
	return l.input[pos]
}

/*
func main() {
	lexer := NewLexer("SELECT column1, column2 FROM table WHERE column1=value GROUP BY column2 ORDER BY column1 LIMIT 10;")
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		fmt.Println(token)
	}
}
*/
