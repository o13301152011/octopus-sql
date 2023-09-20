package SqlPaser

// Parser 结构用于解析令牌数组。
type Parser struct {
	tokens   []Token
	current  int
	previous Token
	errors   []string
}

// NewParser 创建并返回一个新的Parser实例。
func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) next() {
	if p.current < len(p.tokens)-1 {
		p.current++
	}
}

func (p *Parser) currentToken() Token {
	if p.current >= len(p.tokens) {
		return Token{EOF, ""}
	}
	return p.tokens[p.current]
}

func (p *Parser) getPreviousToken() Token {
	if p.current-1 < 0 {
		return Token{EOF, ""}
	}
	return p.tokens[p.current-1]
}

// peek 查看当前令牌，但不消费它。
func (p *Parser) peek() Token {
	if p.current >= len(p.tokens) {
		return Token{EOF, ""}
	}
	return p.tokens[p.current]
}

// advance 消费当前令牌并前进到下一个。
func (p *Parser) advance() Token {
	if p.current < len(p.tokens) {
		p.current++
	}
	return p.getPreviousToken()
}

// match 检查当前令牌是否匹配给定的令牌类型之一。
func (p *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if p.peek().Type == t {
			p.advance()
			return true
		}
	}
	return false
}

// expect 检查当前令牌是否匹配给定的令牌类型，并消费它。
func (p *Parser) expect(t TokenType) *Token {
	if p.peek().Type == t {
		token := p.advance()
		return &token
	}
	return nil
}

// Parse 将提供的令牌解析为一个AST。
func (p *Parser) Parse() ASTNode {
	switch {
	case p.match(SELECT):
		return p.parseSelectStatement()
	case p.match(INSERT):
		return p.parseInsertStatement()
	case p.match(UPDATE):
		return p.parseUpdateStatement()
	case p.match(DELETE):
		return p.parseDeleteStatement()
	default:
		// TODO: 处理未识别的令牌类型或语法错误。
		return nil
	}
}

// parseDeleteStatement 解析一个DELETE语句。
func (p *Parser) parseDeleteStatement() *DeleteStatement {
	stmt := &DeleteStatement{}
	stmt.TableName = p.expect(IDENTIFIER).Value

	// 解析WHERE子句（如果存在）。
	if p.match(WHERE) {
		stmt.Where = p.parseWhereClause()
	}

	return stmt
}

// ... 其他已存在的函数和逻辑。

// parseSelectStatement 解析一个SELECT语句。
func (p *Parser) parseSelectStatement() *SelectStatement {
	stmt := &SelectStatement{}
	// 解析是否存在 DISTINCT 关键字。
	stmt.Distinct = p.match(DISTINCT)

	// 解析列。
	stmt.Columns = p.parseColumns()

	// 如果存在，解析FROM子句。
	if p.match(FROM) {
		stmt.From = p.parseFromClause()
	}

	// TODO: 解析其他子句，如 WHERE, GROUP BY, HAVING, ORDER BY, LIMIT等。

	return stmt
}

// parseColumns 解析列。
func (p *Parser) parseColumns() []ASTNode {
	var columns []ASTNode
	for {
		if p.match(STAR) {
			columns = append(columns, &Star{})
		} else {
			columns = append(columns, &Identifier{Name: p.expect(IDENTIFIER).Value})
		}
		if !p.match(COMMA) {
			break
		}
	}
	return columns
}

// parseFromClause 解析FROM子句。
func (p *Parser) parseFromClause() *FromClause {
	from := &FromClause{}
	from.TableName = p.expect(IDENTIFIER).Value

	// TODO: 解析JOINs, ALIAS等。

	return from
}

// parseInsertStatement 解析一个INSERT语句。
func (p *Parser) parseInsertStatement() *InsertStatement {
	stmt := &InsertStatement{}
	stmt.TableName = p.expect(IDENTIFIER).Value

	// 解析列名称（如果提供了）。
	if p.match(LEFT_PAREN) {
		for !p.match(RIGHT_PAREN) {
			stmt.Columns = append(stmt.Columns, p.expect(IDENTIFIER).Value)
			p.match(COMMA) // 可选的，如果还有更多的列。
		}
	}

	p.expect(VALUES)

	// 解析插入的值。
	for p.match(LEFT_PAREN) {
		var values []ASTNode
		for !p.match(RIGHT_PAREN) {
			values = append(values, p.parseExpression())
			p.match(COMMA) // 可选的，如果还有更多的值。
		}
		stmt.Values = append(stmt.Values, values)
		if !p.match(COMMA) { // 如果不再有更多的数据行。
			break
		}
	}

	return stmt
}

// parseUpdateStatement 解析一个UPDATE语句。
func (p *Parser) parseUpdateStatement() *UpdateStatement {
	stmt := &UpdateStatement{}
	stmt.TableName = p.expect(IDENTIFIER).Value

	p.expect(SET)
	for {
		updateExpr := &UpdateExpression{
			Column: p.expect(IDENTIFIER).Value,
		}
		p.expect(EQUALS)
		updateExpr.Value = p.parseExpression()
		stmt.Updates = append(stmt.Updates, updateExpr)
		if !p.match(COMMA) { // 如果不再有更多的列更新。
			break
		}
	}

	if p.match(WHERE) {
		stmt.Where = p.parseWhereClause()
	}

	return stmt
}
