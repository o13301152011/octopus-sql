package SqlPaser

// ASTNode 代表抽象语法树中的一个节点。
type ASTNode interface {
	// 这个接口可以稍后扩展，例如添加一个方法来转换节点为字符串等。
}

// FromClause 代表FROM子句。
type FromClause struct {
	TableName string
	Alias     string // 可选的表别名。
	Joins     []*JoinClause
}

// JoinClause 代表JOIN子句。
type JoinClause struct {
	Type      string //例如 INNER, LEFT OUTER 等。
	TableName string
	Alias     string       // 可选的表别名。
	On        *WhereClause // JOIN的ON条件。
}

// WhereClause 代表WHERE子句。
type WhereClause struct {
	Condition ASTNode
}

// GroupByClause 代表GROUP BY子句。
type GroupByClause struct {
	Columns []ASTNode
}

// HavingClause 代表HAVING子句。
type HavingClause struct {
	Condition ASTNode
}

// OrderByClause 代表ORDER BY子句。
type OrderByClause struct {
	Columns []*OrderByExpression
}

// OrderByExpression 代表一个ORDER BY中的表达式。
type OrderByExpression struct {
	Column   ASTNode
	Ordering string //例如 ASC 或 DESC。
}

// LimitClause 代表LIMIT子句。
type LimitClause struct {
	Offset int
	Count  int
}

// SelectStatement 代表一个SELECT语句。
type SelectStatement struct {
	Distinct bool
	Columns  []ASTNode // 这可以是`*`或特定的列。
	From     *FromClause
	Where    *WhereClause
	GroupBy  *GroupByClause
	Having   *HavingClause
	OrderBy  *OrderByClause
	Limit    *LimitClause
}

// InsertStatement 代表一个INSERT语句。
type InsertStatement struct {
	TableName string
	Columns   []string    // 要插入的列名称。
	Values    [][]ASTNode // 每个子数组代表一行的值。
}

// UpdateStatement 代表一个UPDATE语句。
type UpdateStatement struct {
	TableName string
	Updates   []*UpdateExpression
	Where     *WhereClause
}

// UpdateExpression 代表UPDATE语句中的一个更新表达式。
type UpdateExpression struct {
	Column string
	Value  ASTNode
}

// DeleteStatement 代表一个DELETE语句。
type DeleteStatement struct {
	TableName string
	Where     *WhereClause
}

// CreateTableStatement 代表一个CREATE TABLE语句。
type CreateTableStatement struct {
	TableName string
	Columns   []*ColumnDefinition
}

// ColumnDefinition 代表CREATE TABLE语句中的列定义。
type ColumnDefinition struct {
	Name    string
	Type    string
	NotNull bool
	// ... 其他属性，例如默认值、是否是主键等。
}

// DropTableStatement 代表一个DROP TABLE语句。
type DropTableStatement struct {
	TableName string
}

// CreateIndexStatement 代表一个CREATE INDEX语句。
type CreateIndexStatement struct {
	IndexName string
	TableName string
	Columns   []string
	IsUnique  bool
}

// DropIndexStatement 代表一个DROP INDEX语句。
type DropIndexStatement struct {
	IndexName string
}

// ... 更多的语句结构可以根据需要添加。

// Identifier 代表一个标识符。
type Identifier struct {
	Name string
}

// BinaryExpr 表示两个值之间的二元表达式，如 `column1 = 'value'`。
type BinaryExpr struct {
	Left     ASTNode
	Operator TokenType
	Right    ASTNode
}

// UnaryExpr 表示单目运算表达式，如 `NOT condition`。
type UnaryExpr struct {
	Operator TokenType
	Operand  ASTNode
}

// LiteralValue 表示一个字面值，如数字、字符串等。
type LiteralValue struct {
	Type  TokenType
	Value string
}

// ColumnName 表示一个列名，可能包括表名或别名，如 `table.column`。
type ColumnName struct {
	Table string
	Name  string
}

// FunctionCall 表示一个函数调用。
type FunctionCall struct {
	Name     string
	Args     []ASTNode
	Distinct bool
}

// Star 表示 `*`，用于SELECT语句中。
type Star struct{}

// SubQuery 表示一个子查询。
type SubQuery struct {
	Select *SelectStatement
}

// BetweenExpr 表示BETWEEN表达式。
type BetweenExpr struct {
	Value ASTNode
	From  ASTNode
	To    ASTNode
}

// InExpr 表示IN表达式。
type InExpr struct {
	Value  ASTNode
	Values []ASTNode
}

// LikeExpr 表示LIKE表达式。
type LikeExpr struct {
	Value   ASTNode
	Pattern ASTNode
}

// CaseExpr 表示CASE表达式。
type CaseExpr struct {
	Expr     ASTNode
	Branches []*CaseBranch
	Default  ASTNode
}

// CaseBranch 表示CASE表达式中的一个分支。
type CaseBranch struct {
	Condition ASTNode
	Result    ASTNode
}

// AlterTableStatement 代表一个ALTER TABLE语句。
type AlterTableStatement struct {
	TableName string
	Actions   []*AlterAction
}

// AlterAction 代表ALTER TABLE中的一个动作。
type AlterAction struct {
	Type     AlterActionType
	Column   *ColumnDefinition
	NewName  string
	Position int // 对于某些数据库，可能需要指定新列的位置。
}

// AlterActionType 代表ALTER TABLE动作的类型。
type AlterActionType int

const (
	ADD_COLUMN AlterActionType = iota
	DROP_COLUMN
	MODIFY_COLUMN
	RENAME_COLUMN
	// ... 其他可能的动作。
)

// TransactionStatement 代表事务相关的语句，如START TRANSACTION、COMMIT、ROLLBACK。
type TransactionStatement struct {
	Type TransactionType
}

// TransactionType 代表事务操作的类型。
type TransactionType int

// WithClause 代表WITH子句，即公共表表达式。
type WithClause struct {
	CTEs []*CommonTableExpression
}

// CommonTableExpression 代表公共表表达式。
type CommonTableExpression struct {
	Name    string
	Columns []string
	Select  *SelectStatement
}

// NestedSubQuery 代表嵌套的子查询。
type NestedSubQuery struct {
	Select *SelectStatement
	Alias  string
}

// ... 其他已存在的结构和定义。
