package PythonSqlPaser

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
	"testing"
)

// 已经定义了 Token 变量，因此我们将结构体重命名为 ParsedToken
type ParsedToken struct {
	Type  TokenType
	Value string
}

// IsKeyword 检查给定的值是否为关键字
func IsKeyword(value string) TokenType {
	val := strings.ToUpper(value)
	tokenType, isOk := KEYWORDS[val]
	if isOk {
		return *tokenType
	}
	tokenType, isOk = KEYWORDS_COMMON[val]
	if isOk {
		return *tokenType
	}
	tokenType, isOk = KEYWORDS_ORACLE[val]
	if isOk {
		return *tokenType
	}
	tokenType, isOk = KEYWORDS_PLPGSQL[val]
	if isOk {
		return *tokenType
	}
	tokenType, isOk = KEYWORDS_HQL[val]
	if isOk {
		return *tokenType
	}
	tokenType, isOk = KEYWORDS_MSACCESS[val]
	if isOk {
		return *tokenType
	}
	return *Name
}

func GetTokens(text string) ([]ParsedToken, error) {
	var tokens []ParsedToken
	pos := 0

	for pos < len(text) {
		matched := false

		for _, rule := range SQL_REGEX {
			rexMatch, err := regexp2.Compile(rule.Regex, regexp2.None)
			if err != nil {
				return nil, fmt.Errorf("failed to compile regex: %v", err)
			}

			matches, err := rexMatch.FindStringMatchStartingAt(text, pos)
			if err != nil {
				return nil, err
			}

			if matches != nil && matches.Index == pos {
				if rule.Token.String() == "PROCESS_AS_KEYWORD" {
					token := IsKeyword(matches.String())
					tokens = append(tokens, ParsedToken{Type: token, Value: matches.String()})
				} else {
					tokens = append(tokens, ParsedToken{Type: *rule.Token, Value: matches.String()})
				}

				pos = matches.Index + len(matches.String())
				matched = true
				break
			}
		}

		if !matched {
			return nil, fmt.Errorf("error: no match at position %d", pos)
		}
	}

	return tokens, nil
}

// 定义一些测试用的输入字符串
var testInputs = []string{
	`-- 这是一个注释
	SELECT * FROM users WHERE name = 'John Doe';
	这是一些无效的内容
	INSERT INTO orders (user_id, total) VALUES (1, 100);
	/* 另一个注释 */
	CREATE TABLE products (id INT PRIMARY KEY, name VARCHAR(255));
	`,
}

// 测试函数
func TestSQLProcessing(t *testing.T) {
	for i, input := range testInputs {
		// 使用GetTokens函数获取Token列表
		tokens, err := GetTokens(input)
		if err != nil {
			t.Errorf("Test case %d: Error getting tokens: %v", i+1, err)
			continue
		}

		// 使用cleanAndSplitTokens函数清洗和拆分Token列表
		statements, err := cleanAndSplitTokens(tokens)
		if err != nil {
			t.Errorf("Test case %d: Error cleaning and splitting tokens: %v", i+1, err)
			continue
		}

		// 检查结果是否符合预期
		if len(statements) != 3 {
			t.Errorf("Test case %d: Expected 3 SQL statements, got %d", i+1, len(statements))
		} else {
			fmt.Printf("Test case %d: Successfully extracted %d SQL statements.\n", i+1, len(statements))
			for j, statement := range statements {
				fmt.Printf("  Statement %d: %v\n", j+1, statement)
			}
		}
	}
}
