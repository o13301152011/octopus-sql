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
		case p.match(BETWEEN):
			lowerBound := p.parseTerm()
			p.expect(AND)
			upperBound := p.parseTerm()
			expr = &BetweenExpr{Operand: expr, LowerBound: lowerBound, UpperBound: upperBound}
		default:
			return expr
		}
	}
}

func (p *Parser) parseTerm() ASTNode {
	expr := p.parseFactor()
	for {
		switch {
		case p.match(PLUS):
			right := p.parseFactor()
			expr = &BinaryExpr{Left: expr, Operator: PLUS, Right: right}
		case p.match(MINUS):
			right := p.parseFactor()
			expr = &BinaryExpr{Left: expr, Operator: MINUS, Right: right}
		default:
			return expr
		}
	}
}

func (p *Parser) parseFactor() ASTNode {
	expr := p.parsePrimary()
	for {
		switch {
		case p.match(MULTIPLY):
			right := p.parsePrimary()
			expr = &BinaryExpr{Left: expr, Operator: MULTIPLY, Right: right}
		case p.match(DIVIDE):
			right := p.parsePrimary()
			expr = &BinaryExpr{Left: expr, Operator: DIVIDE, Right: right}
		default:
			return expr
		}
	}
}

func (p *Parser) parsePrimary() ASTNode {
	if token := p.expect(NUMBER); token != nil {
		return &NumberLiteral{Value: token.Value}
	}
	if token := p.expect(STRING); token != nil {
		return &StringLiteral{Value: token.Value}
	}
	if p.match(NULL) {
		return &NullLiteral{}
	}
	if token := p.expect(TRUE); token != nil {
		return &BooleanLiteral{Value: true}
	}
	if token := p.expect(FALSE); token != nil {
		return &BooleanLiteral{Value: false}
	}
	if p.peek().Type == FUNCTION {
		return p.parseFunctionCall()
	}
	if token := p.expect(IDENTIFIER); token != nil {
		return &Identifier{Name: token.Value}
	}
	if p.match(LEFT_PAREN) {
		if p.isSubquery() {
			return p.parseSubquery()
		}
		expr := p.parseExpression()
		p.expect(RIGHT_PAREN)
		return expr
	}

	// 如果没有匹配到任何已知的模式，抛出错误或返回nil
	return nil
}

func (p *Parser) parseFunctionCall() *FunctionCall {
	funcName := p.currentToken().Value
	p.expect(FUNCTION)
	p.expect(LEFT_PAREN)

	// 检查是否存在 DISTINCT 关键字
	isDistinct := false
	if p.match(DISTINCT) {
		isDistinct = true
	}

	var args []ASTNode
	if !p.match(RIGHT_PAREN) { // 如果参数列表不是空的
		for {
			args = append(args, p.parseExpression())
			if !p.match(COMMA) {
				break
			}
		}
		p.expect(RIGHT_PAREN)
	}
	return &FunctionCall{Name: funcName, Args: args, Distinct: isDistinct}
}

func (p *Parser) isSubquery() bool {
	// 这里只是一个简单的检查，实际上可能需要更复杂的逻辑来检查是否是子查询
	return p.peek().Type == SELECT
}

func (p *Parser) parseSubquery() *Subquery {
	p.expect(LEFT_PAREN)
	stmt := p.parseSelectStatement() // 这里你应该已经有一个解析SELECT语句的函数
	p.expect(RIGHT_PAREN)
	return &Subquery{Statement: stmt}
}
