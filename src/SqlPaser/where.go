package SqlPaser

// parseWhereClause 解析WHERE子句。
func (p *Parser) parseWhereClause() *WhereClause {
	condition := p.parseExpression()
	return &WhereClause{Condition: condition}
}

// parseExpression 递归地解析表达式。
func (p *Parser) parseExpression() ASTNode {
	return p.parseOrExpression()
}

func (p *Parser) parseOrExpression() ASTNode {
	expr := p.parseAndExpression()
	for p.match(OR) {
		right := p.parseAndExpression()
		expr = &BinaryExpr{Left: expr, Operator: OR, Right: right}
	}
	return expr
}

func (p *Parser) parseAndExpression() ASTNode {
	expr := p.parseNotExpression()
	for p.match(AND) {
		right := p.parseNotExpression()
		expr = &BinaryExpr{Left: expr, Operator: AND, Right: right}
	}
	return expr
}

func (p *Parser) parseNotExpression() ASTNode {
	if p.match(NOT) {
		operand := p.parseComparisonExpression()
		return &UnaryExpr{Operator: NOT, Operand: operand}
	}
	return p.parseComparisonExpression()
}

func (p *Parser) parseComparisonExpression() ASTNode {
	expr := p.parseTerm()
	for {
		switch {
		case p.match(EQUALS):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: EQUALS, Right: right}
		case p.match(NOT_EQUALS):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: NOT_EQUALS, Right: right}
		case p.match(LESS_THAN):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: LESS_THAN, Right: right}
		case p.match(GREATER_THAN):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: GREATER_THAN, Right: right}
		case p.match(LESS_EQUALS):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: LESS_EQUALS, Right: right}
		case p.match(GREATER_EQUALS):
			right := p.parseTerm()
			expr = &BinaryExpr{Left: expr, Operator: GREATER_EQUALS, Right: right}
		default:
			return expr
		}
	}
}

func (p *Parser) parseTerm() ASTNode {
	// 为简化起见，这里我们只处理标识符和字面值。
	if p.match(IDENTIFIER) {
		return &Identifier{Name: p.currentToken().Value}
	} else if p.match(NUMBER, STRING) {
		return &LiteralValue{Type: p.currentToken().Type, Value: p.currentToken().Value}
	}
	// TODO: 抛出错误或者处理其他情况。
	return nil
}
