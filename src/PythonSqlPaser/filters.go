package PythonSqlPaser

import (
	"fmt"
	"strings"
)

func cleanAndSplitTokens(tokens []ParsedToken) ([][]ParsedToken, error) {
	var statements [][]ParsedToken
	var currentStatement []ParsedToken
	isInSQLStatement := false

	for _, token := range tokens {
		// 移除注释Token
		if token.Type.String() == CommentSingle.String() || token.Type.String() == Multiline.String() {
			continue
		}

		// 检查是否是SQL语句的起始Token
		if token.Type.String() == DML.String() || token.Type.String() == DDL.String() || (token.Type.String() == Keyword.String() && token.Value == "WITH") {
			isInSQLStatement = true
		}

		// 如果当前处于SQL语句中，则处理当前Token
		if isInSQLStatement {
			// 合并连续的文本Token，例如Identifier和String
			if len(currentStatement) > 0 &&
				(token.Type.String() == Name.String() || token.Type.String() == String.String() ||
					token.Type.String() == Single.String() || token.Type.String() == Symbol.String()) &&
				(currentStatement[len(currentStatement)-1].Type.String() == Name.String() ||
					currentStatement[len(currentStatement)-1].Type.String() == String.String() ||
					currentStatement[len(currentStatement)-1].Type.String() == Single.String() ||
					currentStatement[len(currentStatement)-1].Type.String() == Symbol.String()) {
				currentStatement[len(currentStatement)-1].Value += token.Value
			} else {
				currentStatement = append(currentStatement, token)
			}

			// 如果遇到语句结束标记，则将currentStatement添加到statements，并重置currentStatement
			if token.Type.String() == Punctuation.String() && token.Value == ";" {
				statements = append(statements, currentStatement)
				currentStatement = []ParsedToken{}
				isInSQLStatement = false
			}
		}
	}

	// 添加最后一个语句
	if isInSQLStatement && len(currentStatement) > 0 {
		statements = append(statements, currentStatement)
	}

	return statements, nil
}

func testfilters() {
	// 示例：调用 cleanAndSplitTokens 函数处理Token列表
	tokens := []ParsedToken{
		// 这里添加从GetTokens获取到的Token
	}
	statements, err := cleanAndSplitTokens(tokens)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印处理后的二级Token数组
	for i, statement := range statements {
		fmt.Printf("Statement %d:\n", i+1)
		for _, token := range statement {
			fmt.Printf("  Type: %v, Value: %s\n", token.Type, strings.TrimSpace(token.Value))
		}
	}
}
