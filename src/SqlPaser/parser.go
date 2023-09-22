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

// parseColumns 解析列。
func (p *Parser) parseColumns() []ASTNode {
	var columns []ASTNode
	for {
		if p.match(STAR) {
			columns = append(columns, &Star{})
		} else {
			expr := p.parseExpression() // 使用parseExpression来处理更复杂的表达式
			if p.match(AS) {
				alias := p.expect(IDENTIFIER).Value
				columns = append(columns, &AliasedExpression{Expr: expr, Alias: alias})
			} else {
				columns = append(columns, expr)
			}
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

	// 解析表的别名
	if p.match(AS) {
		p.advance() // 跳过AS
		alias := &AliasClause{Alias: p.expect(IDENTIFIER).Value}
		from.Alias = alias.Alias
	} else if p.peek().Type == IDENTIFIER { // 如果没有AS，但后面是标识符，也视为别名
		alias := &AliasClause{Alias: p.advance().Value}
		from.Alias = alias.Alias
	}

	// 解析所有的JOIN子句
	for p.match(INNER, LEFT, RIGHT, FULL) {
		join := &JoinClause{}
		join.Type = p.currentToken().Type
		p.expect(JOIN) // 期望下一个 token 是 JOIN 关键字

		join.Table = &Identifier{Name: p.expect(IDENTIFIER).Value}

		// 解析JOIN的别名
		if p.match(AS) {
			p.advance()
			alias := &AliasClause{Alias: p.expect(IDENTIFIER).Value}
			join.Alias = alias
		} else if p.peek().Type == IDENTIFIER {
			alias := &AliasClause{Alias: p.advance().Value}
			join.Alias = alias
		}

		if p.match(ON) {
			join.On = p.parseWhereClause()
		}
		from.Joins = append(from.Joins, join)
	}

	return from
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
