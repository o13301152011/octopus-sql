package SqlPaser

import (
	"testing"
)

func (p *Parser) parseSelectStatement() *SelectStatement {
	stmt := &SelectStatement{}

	// 解析是否存在 DISTINCT 关键字。
	if p.match(DISTINCT) {
		stmt.Distinct = true
	}

	// 解析列。
	stmt.Columns = p.parseColumns()

	// 如果存在，解析FROM子句。
	if p.match(FROM) {
		stmt.From = p.parseFromClause()
	}

	// 解析所有的JOIN子句
	for p.match(INNER, LEFT, RIGHT, FULL) {
		joinType := p.currentToken().Type
		p.advance() // 跳过 JOIN 类型
		p.expect(JOIN)
		stmt.From.Joins = append(stmt.From.Joins, p.parseJoinClause(joinType))
	}

	// 解析 WHERE 子句（如果存在）。
	if p.match(WHERE) {
		stmt.Where = p.parseWhereClause()
	}

	// 解析 GROUP BY 子句（如果存在）。
	if p.match(GROUP_BY) {
		stmt.GroupBy = p.parseGroupByClause()
	}

	// 解析 HAVING 子句（如果存在）。
	if p.match(HAVING) {
		stmt.Having = p.parseHavingClause()
	}

	// 解析 ORDER BY 子句（如果存在）。
	if p.match(ORDER_BY) {
		stmt.OrderBy = p.parseOrderByClause()
	}

	// 解析 LIMIT 子句（如果存在）。
	if p.match(LIMIT) {
		stmt.Limit = p.parseLimitClause()
	}

	// TODO: 解析其他子句，如 ALIAS等。

	return stmt
}

func TestParseSelectStatement(t *testing.T) {
	input := `
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
		if token.Type == EOF {
			break
		}
		tokens = append(tokens, token)
	}
	parser := NewParser(tokens)
	stmt := parser.Parse()
	// 打印整个AST
	printAST(stmt, 0)
	// 断言检查
	selectStmt, ok := stmt.(*SelectStatement)
	if !ok {
		t.Fatalf("Expected statement to be of type *SelectStatement, got %T", stmt)
	}

	if selectStmt.Distinct {
		t.Errorf("Expected Distinct to be false")
	}

	if len(selectStmt.Columns) != 4 {
		t.Errorf("Expected 4 columns, got %d", len(selectStmt.Columns))
	}

	// ... 这里继续添加更多的断言来验证其他部分，如FROM子句、JOIN子句、WHERE子句等 ...

	// 例如，验证FROM子句
	if selectStmt.From.TableName != "users" {
		t.Errorf("Expected FROM table to be 'users', got %s", selectStmt.From.TableName)
	}

	if selectStmt.From.Alias != "u" {
		t.Errorf("Expected FROM table alias to be 'u', got %s", selectStmt.From.Alias)
	}

	// ... 更多断言 ...
}
