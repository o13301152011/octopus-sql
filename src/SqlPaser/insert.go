package SqlPaser

import (
	"fmt"
	"testing"
)

func (p *Parser) parseInsertStatement() *InsertStatement {
	stmt := &InsertStatement{}

	// 检查并消费 INSERT 关键字
	p.expect(INSERT)

	// 检查并消费 INTO 关键字
	p.expect(INTO)

	// 获取并消费表名
	tableNameToken := p.expect(IDENTIFIER)
	if tableNameToken.Type != IDENTIFIER {
		// 抛出错误或进行其他错误处理
		panic(fmt.Sprintf("Expected IDENTIFIER for table name, got %s", tableNameToken.Type))
	}
	stmt.TableName = tableNameToken.Value

	// 检查是否有列名列表
	if p.match(LEFT_PAREN) {
		for {
			// 解析列名
			columnToken := p.expect(IDENTIFIER)
			stmt.Columns = append(stmt.Columns, columnToken.Value)

			// 如果下一个令牌是逗号，则继续解析其他列名，否则退出循环
			if !p.match(COMMA) {
				break
			}
		}
		p.expect(RIGHT_PAREN)
	}
	// 检查是否有列名列表
	if p.match(VALUES) {

		// 解析 VALUES 后面的数据
		for {
			p.expect(LEFT_PAREN)
			var valuesRow []ASTNode
			for {
				valueToken := p.peek()
				if valueToken.Type == STRING || valueToken.Type == NUMBER {
					valuesRow = append(valuesRow, &LiteralValue{Type: valueToken.Type, Value: valueToken.Value})
					p.advance()
				} else {
					// TODO: 这里可以处理更复杂的值类型，如函数调用、子查询等。
					// 当前，我们只处理基本的字符串和数字。
					// 抛出错误或其他处理。
					panic("Unexpected token type in VALUES")
				}

				if !p.match(COMMA) {
					break
				}
			}
			p.expect(RIGHT_PAREN)
			stmt.Values = append(stmt.Values, valuesRow)

			if !p.match(COMMA) {
				break
			}
		}
	}
	// 当遇到 SELECT 关键字时
	if p.match(SELECT) {
		stmt.SelectStatement = p.parseSelectStatement() // 利用已有的 SELECT 解析逻辑
	}
	// 最后，期望一个分号或其他合适的终结符
	// p.expect(SEMICOLON)

	return stmt
}

func TestParseInsertStatement(t *testing.T) {
	input := `
INSERT INTO target_table (target_col1, target_col2, target_col3, target_col4)
SELECT 
    u.id AS user_id,  
    u.name, 
    o.order_id, 
    SUM(p.price) AS total_price
FROM 
    users u
LEFT JOIN 
    orders o ON u.id = o.user_id
INNER JOIN 
    order_details od ON o.order_id = od.order_id
LEFT JOIN 
    products p ON od.product_id = p.id
WHERE 
    u.active = 1 AND (o.order_date BETWEEN '2021-01-01' AND '2021-12-31')
GROUP BY 
    u.id, o.order_id
HAVING 
    total_price > 100
ORDER BY 
    total_price DESC, u.name ASC
LIMIT 
    10 OFFSET 5;
`

	lexer := NewLexer(input)
	tokens := []Token{}
	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)
		if token.Type == EOF {
			break
		}
	}
	parser := NewParser(tokens)
	stmt := parser.Parse()
	printAST(stmt, 0)
	// 转换为 InsertStatement 类型
	insertStmt, ok := stmt.(*InsertStatement)
	if !ok {
		t.Fatalf("Expected InsertStatement, got %T", stmt)
	}

	// 验证目标表名
	if insertStmt.TableName != "target_table" {
		t.Errorf("Expected table name 'target_table', got %q", insertStmt.TableName)
	}

	// 验证目标列名
	expectedColumns := []string{"target_col1", "target_col2", "target_col3", "target_col4"}
	if len(insertStmt.Columns) != len(expectedColumns) {
		t.Fatalf("Expected %d columns, got %d", len(expectedColumns), len(insertStmt.Columns))
	}
	for i, col := range expectedColumns {
		if insertStmt.Columns[i] != col {
			t.Errorf("Expected column %q, got %q", col, insertStmt.Columns[i])
		}
	}

	// TODO: 在此处继续验证SELECT子句的各个部分

	// 为了完整性，您需要确保您的InsertStatement结构能够处理SELECT子句。
}
