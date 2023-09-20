package SqlPaser

import (
	"strings"
	"unicode"
)

type Token struct {
	Type  TokenType // 令牌类型
	Value string    // 令牌值
}

type TokenType int

const (
	EOF TokenType = iota
	// 关键字
	SELECT
	INSERT
	UPDATE
	DELETE
	FROM
	WHERE
	JOIN
	ON
	GROUP_BY
	HAVING
	ORDER_BY
	LIMIT
	AND
	OR
	NOT
	AS
	DISTINCT
	// 操作符
	EQUALS
	NOT_EQUALS
	LESS_THAN
	GREATER_THAN
	LESS_EQUALS
	GREATER_EQUALS
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	// 字面量
	STRING
	NUMBER
	// 特殊字符
	COMMA
	LEFT_PAREN
	RIGHT_PAREN
	DOT
	SEMICOLON
	// 其他
	IDENTIFIER
	COMMENT
	FUNCTION
	// 新增关键字
	INNER
	LEFT
	RIGHT
	OUTER
	FULL
	UNION
	INTERSECT
	EXCEPT
	IS
	NULL
	LIKE
	BETWEEN
	IN
	EXISTS
	CASE
	WHEN
	THEN
	ELSE
	END
	BEGIN
	COMMIT
	ROLLBACK
	CREATE
	ALTER
	DROP
	INDEX
	TABLE
	DATABASE
	VIEW
	TRIGGER
	// 方言特性
	// MySQL
	REPLACE
	// PostgreSQL
	RETURNING
	// SQL Server
	TOP
	// Oracle
	ROWNUM
	PARTITION
	// SQLite
	VACUUM
	// ...
	SET
	STAR
	VALUES
)

// 关键字映射
var keywords = map[string]TokenType{
	"SELECT":   SELECT,
	"INSERT":   INSERT,
	"UPDATE":   UPDATE,
	"DELETE":   DELETE,
	"FROM":     FROM,
	"WHERE":    WHERE,
	"JOIN":     JOIN,
	"ON":       ON,
	"GROUP":    GROUP_BY,
	"HAVING":   HAVING,
	"ORDER":    ORDER_BY,
	"LIMIT":    LIMIT,
	"AND":      AND,
	"OR":       OR,
	"NOT":      NOT,
	"AS":       AS,
	"DISTINCT": DISTINCT,
	",":        COMMA,
	"=":        EQUALS,
	"!=":       NOT_EQUALS,
	"<>":       NOT_EQUALS,
	"<":        LESS_THAN,
	">":        GREATER_THAN,
	"<=":       LESS_EQUALS,
	">=":       GREATER_EQUALS,
	"+":        PLUS,
	"-":        MINUS,
	"*":        MULTIPLY,
	"/":        DIVIDE,
	"(":        LEFT_PAREN,
	")":        RIGHT_PAREN,
	".":        DOT,
	";":        SEMICOLON,
	// ... 更多关键字
	"INNER":     INNER,
	"LEFT":      LEFT,
	"RIGHT":     RIGHT,
	"OUTER":     OUTER,
	"FULL":      FULL,
	"UNION":     UNION,
	"INTERSECT": INTERSECT,
	"EXCEPT":    EXCEPT,
	"IS":        IS,
	"NULL":      NULL,
	"LIKE":      LIKE,
	"BETWEEN":   BETWEEN,
	"IN":        IN,
	"EXISTS":    EXISTS,
	"CASE":      CASE,
	"WHEN":      WHEN,
	"THEN":      THEN,
	"ELSE":      ELSE,
	"END":       END,
	"BEGIN":     BEGIN,
	"COMMIT":    COMMIT,
	"ROLLBACK":  ROLLBACK,
	"CREATE":    CREATE,
	"ALTER":     ALTER,
	"DROP":      DROP,
	"INDEX":     INDEX,
	"TABLE":     TABLE,
	"DATABASE":  DATABASE,
	"VIEW":      VIEW,
	"TRIGGER":   TRIGGER,
	// 方言特性
	"REPLACE":   REPLACE,   // MySQL
	"RETURNING": RETURNING, // PostgreSQL
	"TOP":       TOP,       // SQL Server
	"ROWNUM":    ROWNUM,    // Oracle
	"PARTITION": PARTITION, // Oracle
	"VACUUM":    VACUUM,    // SQLite
	// ...
	"SET":    SET,
	"VALUES": VALUES,
}

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
		return l.lexKeywordOrIdentifier()
	case unicode.IsDigit(rune(ch)):
		return l.lexNumber()
	case ch == ',':
		l.pos++
		return Token{COMMA, ","}
	// ... 其他字符处理
	default:
		l.pos++
		return Token{EOF, string(ch)}
	}
}

// 解析关键字或标识符
func (l *Lexer) lexKeywordOrIdentifier() Token {
	start := l.pos
	for l.pos < len(l.input) && (unicode.IsLetter(rune(l.input[l.pos])) || unicode.IsDigit(rune(l.input[l.pos]))) {
		l.pos++
	}
	word := strings.ToUpper(l.input[start:l.pos])
	if keyword, ok := keywords[word]; ok {
		return Token{keyword, word}
	}
	return Token{IDENTIFIER, l.input[start:l.pos]}
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
