package SqlPaser

import "strconv"

// parseGroupByClause 解析 GROUP BY 子句。
func (p *Parser) parseGroupByClause() *GroupByClause {
	groupBy := &GroupByClause{}
	// 期望紧跟着'BY'
	p.expect(BY)
	// 解析 GROUP BY 后的列
	for {
		column := p.expect(IDENTIFIER)
		groupBy.Columns = append(groupBy.Columns, &Identifier{Name: column.Value})

		if !p.match(COMMA) {
			break
		}
	}

	return groupBy
}

// parseHavingClause 解析 HAVING 子句。
func (p *Parser) parseHavingClause() *HavingClause {
	having := &HavingClause{}
	having.Condition = p.parseExpression()
	return having
}

// parseOrderByClause 解析 ORDER BY 子句。
func (p *Parser) parseOrderByClause() *OrderByClause {
	orderBy := &OrderByClause{}
	// 期望紧跟着'BY'
	p.expect(BY)
	// 解析 ORDER BY 后的列和排序方向（ASC 或 DESC）
	for {
		column := p.expect(IDENTIFIER)
		order := &OrderByExpression{Column: &Identifier{Name: column.Value}}

		if p.match(ASC) {
			order.Direction = ASC
		} else if p.match(DESC) {
			order.Direction = DESC
		}

		orderBy.Columns = append(orderBy.Columns, order)

		if !p.match(COMMA) {
			break
		}
	}

	return orderBy
}

// parseLimitClause 解析 LIMIT 子句。
func (p *Parser) parseLimitClause() *LimitClause {
	limit := &LimitClause{}
	limitValue := p.expect(NUMBER)
	limit.Count, _ = strconv.Atoi(limitValue.Value)

	// 如果有 OFFSET 关键字，解析它
	if p.match(OFFSET) {
		limit.Offset = limit.Count
	}

	return limit
}

// 新增的解析 JOIN 子句的函数
func (p *Parser) parseJoinClause(joinType TokenType) *JoinClause {
	join := &JoinClause{
		Type: joinType,
	}

	// 解析 JOIN 后的表名
	join.Table = &Identifier{Name: p.expect(IDENTIFIER).Value}

	// 检查是否有 AS 别名
	if p.match(AS) {
		p.advance() // 跳过 AS
		join.Alias = &AliasClause{Alias: p.expect(IDENTIFIER).Value}
	}

	// 解析 ON 子句
	if p.match(ON) {
		join.On = p.parseWhereClause()
	}

	return join
}
