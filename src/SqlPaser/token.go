package SqlPaser

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
	ASC
	DESC
	OFFSET
	INTO
	LEFT_JOIN
	RIGHT_JOIN
	INNER_JOIN
	FULL_JOIN
	TRUE
	FALSE
	BY
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
	"ASC":    ASC,
	"DESC":   DESC,
	"OFFSET": OFFSET,
	"INTO":   INTO,
	"BY":     BY,
}
