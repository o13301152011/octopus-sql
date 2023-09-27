package PythonSqlPaser

// RegexRule 定义了正则表达式规则和对应的 Token 类型
type RegexRule struct {
	Regex string
	Token *TokenType
}

var SQL_REGEX = []RegexRule{
	{`(--|# )\+.*?(\r\n|\r|\n|$)`, SingleHint},
	{`/\*\+[\s\S]*?\*/`, MultilineHint},
	{`(--|# ).*?(\r\n|\r|\n|$)`, Single},
	{`/\*[\s\S]*?\*/`, Multiline},
	{`(\r\n|\r|\n)`, Newline},
	{`\s+?`, Whitespace},
	{`:=`, Assignment},
	{`::`, Punctuation},
	{`\*`, Wildcard},
	{"`(``|[^`])*`", Name},
	{"´(´´|[^´])*´", Name},
	{`((?<!\S)\$(?:[_A-ZÀ-Ü]\w*)?\$)[\s\S]*?\1`, Literal},

	{`\?`, Placeholder},
	{`%(\(\w+\))?s`, Placeholder},
	{`(?<!\w)[$:?]\w+`, Placeholder},

	{`\\\w+`, Command},

	{`(CASE|IN|VALUES|USING|FROM|AS)\b`, Keyword},

	{`(@|##|#)[A-ZÀ-Ü]\w+`, Name},

	{`[A-ZÀ-Ü]\w*(?=\s*\.)`, Name},

	{`(?<=\.)[A-ZÀ-Ü]\w*`, Name},
	{`[A-ZÀ-Ü]\w*(?=\()`, Name},
	{`-?0x[\dA-F]+`, Hexadecimal},
	{`-?\d+(\.\d+)?E-?\d+`, Float},
	{`(?![_A-ZÀ-Ü])-?(\d+(\.\d*)|\.\d+)(?![_A-ZÀ-Ü])`, Float},
	{`(?![_A-ZÀ-Ü])-?\d+(?![_A-ZÀ-Ü])`, Integer},
	{`'(''|\\'|[^'])*'`, Single},

	{`"(""|\\"|[^"])*"`, Symbol},
	{`(""|".*?[^\\]")`, Symbol},

	{`(?<![\w\])])(\[[^\]\[]+\])`, Name},
	{`((LEFT\s+|RIGHT\s+|FULL\s+)?(INNER\s+|OUTER\s+|STRAIGHT\s+)?|(CROSS\s+|NATURAL\s+)?)?JOIN\b`, Keyword},
	{`END(\s+IF|\s+LOOP|\s+WHILE)?\b`, Keyword},
	{`NOT\s+NULL\b`, Keyword},
	{`NULLS\s+(FIRST|LAST)\b`, Keyword},
	{`UNION\s+ALL\b`, Keyword},
	{`CREATE(\s+OR\s+REPLACE)?\b`, DDL},
	{`DOUBLE\s+PRECISION\b`, Builtin},
	{`GROUP\s+BY\b`, Keyword},
	{`ORDER\s+BY\b`, Keyword},
	{`HANDLER\s+FOR\b`, Keyword},
	{`(LATERAL\s+VIEW\s+)(EXPLODE|INLINE|PARSE_URL_TUPLE|POSEXPLODE|STACK)\b`, Keyword},
	{`(AT|WITH')\s+TIME\s+ZONE\s+'[^']+'`, TZCast},
	{`(NOT\s+)?(LIKE|ILIKE|RLIKE)\b`, Comparison},
	{`(NOT\s+)?(REGEXP)\b`, Comparison},
	// 添加 PROCESS_AS_KEYWORD 的处理，这里需要根据具体的需求来实现这个规则。
	{`\w[$#\w]*`, PROCESS_AS_KEYWORD}, // 临时将此规则映射到 Keyword，需要根据实际情况调整。
	{`[;:()\[\],\.]`, Punctuation},
	{`[<>=~!]+`, Comparison},
	{`[+/@#%^&|^-]+`, Operator},
}

var KEYWORDS = map[string]*TokenType{
	"ABORT":          Keyword,
	"ABS":            Keyword,
	"ABSOLUTE":       Keyword,
	"ACCESS":         Keyword,
	"ADA":            Keyword,
	"ADD":            Keyword,
	"ADMIN":          Keyword,
	"AFTER":          Keyword,
	"AGGREGATE":      Keyword,
	"ALIAS":          Keyword,
	"ALL":            Keyword,
	"ALLOCATE":       Keyword,
	"ANALYSE":        Keyword,
	"ANALYZE":        Keyword,
	"ANY":            Keyword,
	"ARRAYLEN":       Keyword,
	"ARE":            Keyword,
	"ASC":            Order,
	"ASENSITIVE":     Keyword,
	"ASSERTION":      Keyword,
	"ASSIGNMENT":     Keyword,
	"ASYMMETRIC":     Keyword,
	"AT":             Keyword,
	"ATOMIC":         Keyword,
	"AUDIT":          Keyword,
	"AUTHORIZATION":  Keyword,
	"AUTO_INCREMENT": Keyword,
	"AVG":            Keyword,

	"BACKWARD":   Keyword,
	"BEFORE":     Keyword,
	"BEGIN":      Keyword,
	"BETWEEN":    Keyword,
	"BITVAR":     Keyword,
	"BIT_LENGTH": Keyword,
	"BOTH":       Keyword,
	"BREADTH":    Keyword,

	// "C": Keyword,  # most likely this is an alias
	"CACHE":                 Keyword,
	"CALL":                  Keyword,
	"CALLED":                Keyword,
	"CARDINALITY":           Keyword,
	"CASCADE":               Keyword,
	"CASCADED":              Keyword,
	"CAST":                  Keyword,
	"CATALOG":               Keyword,
	"CATALOG_NAME":          Keyword,
	"CHAIN":                 Keyword,
	"CHARACTERISTICS":       Keyword,
	"CHARACTER_LENGTH":      Keyword,
	"CHARACTER_SET_CATALOG": Keyword,
	"CHARACTER_SET_NAME":    Keyword,
	"CHARACTER_SET_SCHEMA":  Keyword,
	"CHAR_LENGTH":           Keyword,
	"CHARSET":               Keyword,
	"CHECK":                 Keyword,
	"CHECKED":               Keyword,
	"CHECKPOINT":            Keyword,
	"CLASS":                 Keyword,
	"CLASS_ORIGIN":          Keyword,
	"CLOB":                  Keyword,
	"CLOSE":                 Keyword,
	"CLUSTER":               Keyword,
	"COALESCE":              Keyword,
	"COBOL":                 Keyword,
	"COLLATE":               Keyword,
	"COLLATION":             Keyword,
	"COLLATION_CATALOG":     Keyword,
	"COLLATION_NAME":        Keyword,
	"COLLATION_SCHEMA":      Keyword,
	"COLLECT":               Keyword,
	"COLUMN":                Keyword,
	"COLUMN_NAME":           Keyword,
	"COMPRESS":              Keyword,
	"COMMAND_FUNCTION":      Keyword,
	"COMMAND_FUNCTION_CODE": Keyword,
	"COMMENT":               Keyword,
	"COMMIT":                DML,
	"COMMITTED":             Keyword,
	"COMPLETION":            Keyword,
	"CONCURRENTLY":          Keyword,
	"CONDITION_NUMBER":      Keyword,
	"CONNECT":               Keyword,
	"CONNECTION":            Keyword,
	"CONNECTION_NAME":       Keyword,
	"CONSTRAINT":            Keyword,
	"CONSTRAINTS":           Keyword,
	"CONSTRAINT_CATALOG":    Keyword,
	"CONSTRAINT_NAME":       Keyword,
	"CONSTRAINT_SCHEMA":     Keyword,
	"CONSTRUCTOR":           Keyword,
	"CONTAINS":              Keyword,
	"CONTINUE":              Keyword,
	"CONVERSION":            Keyword,
	"CONVERT":               Keyword,
	"COPY":                  Keyword,
	"CORRESPONDING":         Keyword,
	"COUNT":                 Keyword,
	"CREATEDB":              Keyword,
	"CREATEUSER":            Keyword,
	"CROSS":                 Keyword,
	"CUBE":                  Keyword,
	"CURRENT":               Keyword,
	"CURRENT_DATE":          Keyword,
	"CURRENT_PATH":          Keyword,
	"CURRENT_ROLE":          Keyword,
	"CURRENT_TIME":          Keyword,
	"CURRENT_TIMESTAMP":     Keyword,
	"CURRENT_USER":          Keyword,
	"CURSOR":                Keyword,
	"CURSOR_NAME":           Keyword,
	"CYCLE":                 Keyword,

	"DATA":                        Keyword,
	"DATABASE":                    Keyword,
	"DATETIME_INTERVAL_CODE":      Keyword,
	"DATETIME_INTERVAL_PRECISION": Keyword,
	"DAY":                         Keyword,
	"DEALLOCATE":                  Keyword,
	"DECLARE":                     Keyword,
	"DEFAULT":                     Keyword,
	"DEFAULTS":                    Keyword,
	"DEFERRABLE":                  Keyword,
	"DEFERRED":                    Keyword,
	"DEFINED":                     Keyword,
	"DEFINER":                     Keyword,
	"DELIMITER":                   Keyword,
	"DELIMITERS":                  Keyword,
	"DEREF":                       Keyword,
	"DESC":                        Order,
	"DESCRIBE":                    Keyword,
	"DESCRIPTOR":                  Keyword,
	"DESTROY":                     Keyword,
	"DESTRUCTOR":                  Keyword,
	"DETERMINISTIC":               Keyword,
	"DIAGNOSTICS":                 Keyword,
	"DICTIONARY":                  Keyword,
	"DISABLE":                     Keyword,
	"DISCONNECT":                  Keyword,
	"DISPATCH":                    Keyword,
	"DIV":                         Operator,
	"DO":                          Keyword,
	"DOMAIN":                      Keyword,
	"DYNAMIC":                     Keyword,
	"DYNAMIC_FUNCTION":            Keyword,
	"DYNAMIC_FUNCTION_CODE":       Keyword,

	"EACH":      Keyword,
	"ENABLE":    Keyword,
	"ENCODING":  Keyword,
	"ENCRYPTED": Keyword,
	"END-EXEC":  Keyword,
	"ENGINE":    Keyword,
	"EQUALS":    Keyword,
	"ESCAPE":    Keyword,
	"EVERY":     Keyword,
	"EXCEPT":    Keyword,
	"EXCEPTION": Keyword,
	"EXCLUDING": Keyword,
	"EXCLUSIVE": Keyword,
	"EXEC":      Keyword,
	"EXECUTE":   Keyword,
	"EXISTING":  Keyword,
	"EXISTS":    Keyword,
	"EXPLAIN":   Keyword,
	"EXTERNAL":  Keyword,
	"EXTRACT":   Keyword,

	"FALSE":    Keyword,
	"FETCH":    Keyword,
	"FILE":     Keyword,
	"FINAL":    Keyword,
	"FIRST":    Keyword,
	"FORCE":    Keyword,
	"FOREACH":  Keyword,
	"FOREIGN":  Keyword,
	"FORTRAN":  Keyword,
	"FORWARD":  Keyword,
	"FOUND":    Keyword,
	"FREE":     Keyword,
	"FREEZE":   Keyword,
	"FULL":     Keyword,
	"FUNCTION": Keyword,

	//# "G": Keyword,
	"GENERAL":   Keyword,
	"GENERATED": Keyword,
	"GET":       Keyword,
	"GLOBAL":    Keyword,
	"GO":        Keyword,
	"GOTO":      Keyword,
	"GRANT":     Keyword,
	"GRANTED":   Keyword,
	"GROUPING":  Keyword,

	"HAVING":    Keyword,
	"HIERARCHY": Keyword,
	"HOLD":      Keyword,
	"HOUR":      Keyword,
	"HOST":      Keyword,

	"IDENTIFIED": Keyword,
	"IDENTITY":   Keyword,
	"IGNORE":     Keyword,
	"ILIKE":      Keyword,
	"IMMEDIATE":  Keyword,
	"IMMUTABLE":  Keyword,

	"IMPLEMENTATION": Keyword,
	"IMPLICIT":       Keyword,
	"INCLUDING":      Keyword,
	"INCREMENT":      Keyword,
	"INDEX":          Keyword,

	"INDICATOR":    Keyword,
	"INFIX":        Keyword,
	"INHERITS":     Keyword,
	"INITIAL":      Keyword,
	"INITIALIZE":   Keyword,
	"INITIALLY":    Keyword,
	"INOUT":        Keyword,
	"INPUT":        Keyword,
	"INSENSITIVE":  Keyword,
	"INSTANTIABLE": Keyword,
	"INSTEAD":      Keyword,
	"INTERSECT":    Keyword,
	"INTO":         Keyword,
	"INVOKER":      Keyword,
	"IS":           Keyword,
	"ISNULL":       Keyword,
	"ISOLATION":    Keyword,
	"ITERATE":      Keyword,

	//# "K": Keyword,
	"KEY":        Keyword,
	"KEY_MEMBER": Keyword,
	"KEY_TYPE":   Keyword,

	"LANCOMPILER":    Keyword,
	"LANGUAGE":       Keyword,
	"LARGE":          Keyword,
	"LAST":           Keyword,
	"LATERAL":        Keyword,
	"LEADING":        Keyword,
	"LENGTH":         Keyword,
	"LESS":           Keyword,
	"LEVEL":          Keyword,
	"LIMIT":          Keyword,
	"LISTEN":         Keyword,
	"LOAD":           Keyword,
	"LOCAL":          Keyword,
	"LOCALTIME":      Keyword,
	"LOCALTIMESTAMP": Keyword,
	"LOCATION":       Keyword,
	"LOCATOR":        Keyword,
	"LOCK":           Keyword,
	"LOWER":          Keyword,

	//# "M": Keyword,
	"MAP":                  Keyword,
	"MATCH":                Keyword,
	"MAXEXTENTS":           Keyword,
	"MAXVALUE":             Keyword,
	"MESSAGE_LENGTH":       Keyword,
	"MESSAGE_OCTET_LENGTH": Keyword,
	"MESSAGE_TEXT":         Keyword,
	"METHOD":               Keyword,
	"MINUTE":               Keyword,
	"MINUS":                Keyword,
	"MINVALUE":             Keyword,
	"MOD":                  Keyword,
	"MODE":                 Keyword,
	"MODIFIES":             Keyword,
	"MODIFY":               Keyword,
	"MONTH":                Keyword,
	"MORE":                 Keyword,
	"MOVE":                 Keyword,
	"MUMPS":                Keyword,

	"NAMES":        Keyword,
	"NATIONAL":     Keyword,
	"NATURAL":      Keyword,
	"NCHAR":        Keyword,
	"NCLOB":        Keyword,
	"NEW":          Keyword,
	"NEXT":         Keyword,
	"NO":           Keyword,
	"NOAUDIT":      Keyword,
	"NOCOMPRESS":   Keyword,
	"NOCREATEDB":   Keyword,
	"NOCREATEUSER": Keyword,
	"NONE":         Keyword,
	"NOT":          Keyword,
	"NOTFOUND":     Keyword,
	"NOTHING":      Keyword,
	"NOTIFY":       Keyword,
	"NOTNULL":      Keyword,
	"NOWAIT":       Keyword,
	"NULL":         Keyword,
	"NULLABLE":     Keyword,
	"NULLIF":       Keyword,

	"OBJECT":       Keyword,
	"OCTET_LENGTH": Keyword,
	"OF":           Keyword,
	"OFF":          Keyword,
	"OFFLINE":      Keyword,
	"OFFSET":       Keyword,
	"OIDS":         Keyword,
	"OLD":          Keyword,
	"ONLINE":       Keyword,
	"ONLY":         Keyword,
	"OPEN":         Keyword,
	"OPERATION":    Keyword,
	"OPERATOR":     Keyword,
	"OPTION":       Keyword,
	"OPTIONS":      Keyword,
	"ORDINALITY":   Keyword,
	"OUT":          Keyword,
	"OUTPUT":       Keyword,
	"OVERLAPS":     Keyword,
	"OVERLAY":      Keyword,
	"OVERRIDING":   Keyword,
	"OWNER":        Keyword,

	"QUARTER": Keyword,

	"PAD":                        Keyword,
	"PARAMETER":                  Keyword,
	"PARAMETERS":                 Keyword,
	"PARAMETER_MODE":             Keyword,
	"PARAMETER_NAME":             Keyword,
	"PARAMETER_ORDINAL_POSITION": Keyword,
	"PARAMETER_SPECIFIC_CATALOG": Keyword,
	"PARAMETER_SPECIFIC_NAME":    Keyword,
	"PARAMETER_SPECIFIC_SCHEMA":  Keyword,
	"PARTIAL":                    Keyword,
	"PASCAL":                     Keyword,
	"PCTFREE":                    Keyword,
	"PENDANT":                    Keyword,
	"PLACING":                    Keyword,
	"PLI":                        Keyword,
	"POSITION":                   Keyword,
	"POSTFIX":                    Keyword,
	"PRECISION":                  Keyword,
	"PREFIX":                     Keyword,
	"PREORDER":                   Keyword,
	"PREPARE":                    Keyword,
	"PRESERVE":                   Keyword,
	"PRIMARY":                    Keyword,
	"PRIOR":                      Keyword,
	"PRIVILEGES":                 Keyword,
	"PROCEDURAL":                 Keyword,
	"PROCEDURE":                  Keyword,
	"PUBLIC":                     Keyword,

	"RAISE":                 Keyword,
	"RAW":                   Keyword,
	"READ":                  Keyword,
	"READS":                 Keyword,
	"RECHECK":               Keyword,
	"RECURSIVE":             Keyword,
	"REF":                   Keyword,
	"REFERENCES":            Keyword,
	"REFERENCING":           Keyword,
	"REINDEX":               Keyword,
	"RELATIVE":              Keyword,
	"RENAME":                Keyword,
	"REPEATABLE":            Keyword,
	"RESET":                 Keyword,
	"RESOURCE":              Keyword,
	"RESTART":               Keyword,
	"RESTRICT":              Keyword,
	"RESULT":                Keyword,
	"RETURN":                Keyword,
	"RETURNED_LENGTH":       Keyword,
	"RETURNED_OCTET_LENGTH": Keyword,
	"RETURNED_SQLSTATE":     Keyword,
	"RETURNING":             Keyword,
	"RETURNS":               Keyword,
	"REVOKE":                Keyword,
	"RIGHT":                 Keyword,
	"ROLE":                  Keyword,
	"ROLLBACK":              DML,
	"ROLLUP":                Keyword,
	"ROUTINE":               Keyword,
	"ROUTINE_CATALOG":       Keyword,
	"ROUTINE_NAME":          Keyword,
	"ROUTINE_SCHEMA":        Keyword,
	"ROW":                   Keyword,
	"ROWS":                  Keyword,
	"ROW_COUNT":             Keyword,
	"RULE":                  Keyword,

	"SAVE_POINT":    Keyword,
	"SCALE":         Keyword,
	"SCHEMA":        Keyword,
	"SCHEMA_NAME":   Keyword,
	"SCOPE":         Keyword,
	"SCROLL":        Keyword,
	"SEARCH":        Keyword,
	"SECOND":        Keyword,
	"SECURITY":      Keyword,
	"SELF":          Keyword,
	"SENSITIVE":     Keyword,
	"SEQUENCE":      Keyword,
	"SERIALIZABLE":  Keyword,
	"SERVER_NAME":   Keyword,
	"SESSION":       Keyword,
	"SESSION_USER":  Keyword,
	"SETOF":         Keyword,
	"SETS":          Keyword,
	"SHARE":         Keyword,
	"SHOW":          Keyword,
	"SIMILAR":       Keyword,
	"SIMPLE":        Keyword,
	"SIZE":          Keyword,
	"SOME":          Keyword,
	"SOURCE":        Keyword,
	"SPACE":         Keyword,
	"SPECIFIC":      Keyword,
	"SPECIFICTYPE":  Keyword,
	"SPECIFIC_NAME": Keyword,
	"SQL":           Keyword,
	"SQLBUF":        Keyword,
	"SQLCODE":       Keyword,
	"SQLERROR":      Keyword,
	"SQLEXCEPTION":  Keyword,
	"SQLSTATE":      Keyword,
	"SQLWARNING":    Keyword,
	"STABLE":        Keyword,
	"START":         DML,
	//# "STATE": Keyword,
	"STATEMENT":       Keyword,
	"STATIC":          Keyword,
	"STATISTICS":      Keyword,
	"STDIN":           Keyword,
	"STDOUT":          Keyword,
	"STORAGE":         Keyword,
	"STRICT":          Keyword,
	"STRUCTURE":       Keyword,
	"STYPE":           Keyword,
	"SUBCLASS_ORIGIN": Keyword,
	"SUBLIST":         Keyword,
	"SUBSTRING":       Keyword,
	"SUCCESSFUL":      Keyword,
	"SUM":             Keyword,
	"SYMMETRIC":       Keyword,
	"SYNONYM":         Keyword,
	"SYSID":           Keyword,
	"SYSTEM":          Keyword,
	"SYSTEM_USER":     Keyword,

	"TABLE":                    Keyword,
	"TABLE_NAME":               Keyword,
	"TEMP":                     Keyword,
	"TEMPLATE":                 Keyword,
	"TEMPORARY":                Keyword,
	"TERMINATE":                Keyword,
	"THAN":                     Keyword,
	"TIMESTAMP":                Keyword,
	"TIMEZONE_HOUR":            Keyword,
	"TIMEZONE_MINUTE":          Keyword,
	"TO":                       Keyword,
	"TOAST":                    Keyword,
	"TRAILING":                 Keyword,
	"TRANSATION":               Keyword,
	"TRANSACTIONS_COMMITTED":   Keyword,
	"TRANSACTIONS_ROLLED_BACK": Keyword,
	"TRANSATION_ACTIVE":        Keyword,
	"TRANSFORM":                Keyword,
	"TRANSFORMS":               Keyword,
	"TRANSLATE":                Keyword,
	"TRANSLATION":              Keyword,
	"TREAT":                    Keyword,
	"TRIGGER":                  Keyword,
	"TRIGGER_CATALOG":          Keyword,
	"TRIGGER_NAME":             Keyword,
	"TRIGGER_SCHEMA":           Keyword,
	"TRIM":                     Keyword,
	"TRUE":                     Keyword,
	"TRUNCATE":                 Keyword,
	"TRUSTED":                  Keyword,
	"TYPE":                     Keyword,

	"UID":                       Keyword,
	"UNCOMMITTED":               Keyword,
	"UNDER":                     Keyword,
	"UNENCRYPTED":               Keyword,
	"UNION":                     Keyword,
	"UNIQUE":                    Keyword,
	"UNKNOWN":                   Keyword,
	"UNLISTEN":                  Keyword,
	"UNNAMED":                   Keyword,
	"UNNEST":                    Keyword,
	"UNTIL":                     Keyword,
	"UPPER":                     Keyword,
	"USAGE":                     Keyword,
	"USE":                       Keyword,
	"USER":                      Keyword,
	"USER_DEFINED_TYPE_CATALOG": Keyword,
	"USER_DEFINED_TYPE_NAME":    Keyword,
	"USER_DEFINED_TYPE_SCHEMA":  Keyword,
	"USING":                     Keyword,

	"VACUUM":    Keyword,
	"VALID":     Keyword,
	"VALIDATE":  Keyword,
	"VALIDATOR": Keyword,
	"VALUES":    Keyword,
	"VARIABLE":  Keyword,
	"VERBOSE":   Keyword,
	"VERSION":   Keyword,
	"VIEW":      Keyword,
	"VOLATILE":  Keyword,

	"WEEK":     Keyword,
	"WHENEVER": Keyword,
	"WITH":     CTE,
	"WITHOUT":  Keyword,
	"WORK":     Keyword,
	"WRITE":    Keyword,

	"YEAR": Keyword,

	"ZONE": Keyword,

	//# Name.Builtin
	"ARRAY":          Builtin,
	"BIGINT":         Builtin,
	"BINARY":         Builtin,
	"BIT":            Builtin,
	"BLOB":           Builtin,
	"BOOLEAN":        Builtin,
	"CHAR":           Builtin,
	"CHARACTER":      Builtin,
	"DATE":           Builtin,
	"DEC":            Builtin,
	"DECIMAL":        Builtin,
	"FILE_TYPE":      Builtin,
	"FLOAT":          Builtin,
	"INT":            Builtin,
	"INT8":           Builtin,
	"INTEGER":        Builtin,
	"INTERVAL":       Builtin,
	"LONG":           Builtin,
	"NATURALN":       Builtin,
	"NVARCHAR":       Builtin,
	"NUMBER":         Builtin,
	"NUMERIC":        Builtin,
	"PLS_INTEGER":    Builtin,
	"POSITIVE":       Builtin,
	"POSITIVEN":      Builtin,
	"REAL":           Builtin,
	"ROWID":          Builtin,
	"ROWLABEL":       Builtin,
	"ROWNUM":         Builtin,
	"SERIAL":         Builtin,
	"SERIAL8":        Builtin,
	"SIGNED":         Builtin,
	"SIGNTYPE":       Builtin,
	"SIMPLE_DOUBLE":  Builtin,
	"SIMPLE_FLOAT":   Builtin,
	"SIMPLE_INTEGER": Builtin,
	"SMALLINT":       Builtin,
	"SYS_REFCURSOR":  Builtin,
	"SYSDATE":        Name,
	"TEXT":           Builtin,
	"TINYINT":        Builtin,
	"UNSIGNED":       Builtin,
	"UROWID":         Builtin,
	"UTL_FILE":       Builtin,
	"VARCHAR":        Builtin,
	"VARCHAR2":       Builtin,
	"VARYING":        Builtin,
}

var KEYWORDS_COMMON = map[string]*TokenType{
	"SELECT":  DML,
	"INSERT":  DML,
	"DELETE":  DML,
	"UPDATE":  DML,
	"UPSERT":  DML,
	"REPLACE": DML,
	"MERGE":   DML,
	"DROP":    DDL,
	"CREATE":  DDL,
	"ALTER":   DDL,

	"WHERE":         Keyword,
	"FROM":          Keyword,
	"INNER":         Keyword,
	"JOIN":          Keyword,
	"STRAIGHT_JOIN": Keyword,
	"AND":           Keyword,
	"OR":            Keyword,
	"LIKE":          Keyword,
	"ON":            Keyword,
	"IN":            Keyword,
	"SET":           Keyword,

	"BY":    Keyword,
	"GROUP": Keyword,
	"ORDER": Keyword,
	"LEFT":  Keyword,
	"OUTER": Keyword,
	"FULL":  Keyword,

	"IF":    Keyword,
	"END":   Keyword,
	"THEN":  Keyword,
	"LOOP":  Keyword,
	"AS":    Keyword,
	"ELSE":  Keyword,
	"FOR":   Keyword,
	"WHILE": Keyword,

	"CASE":     Keyword,
	"WHEN":     Keyword,
	"MIN":      Keyword,
	"MAX":      Keyword,
	"DISTINCT": Keyword,
}

var KEYWORDS_ORACLE = map[string]*TokenType{
	"ARCHIVE":    Keyword,
	"ARCHIVELOG": Keyword,

	"BACKUP": Keyword,
	"BECOME": Keyword,
	"BLOCK":  Keyword,
	"BODY":   Keyword,

	"CANCEL":      Keyword,
	"CHANGE":      Keyword,
	"COMPILE":     Keyword,
	"CONTENTS":    Keyword,
	"CONTROLFILE": Keyword,

	"DATAFILE": Keyword,
	"DBA":      Keyword,
	"DISMOUNT": Keyword,
	"DOUBLE":   Keyword,
	"DUMP":     Keyword,

	"ELSIF":      Keyword,
	"EVENTS":     Keyword,
	"EXCEPTIONS": Keyword,
	"EXPLAIN":    Keyword,
	"EXTENT":     Keyword,
	"EXTERNALLY": Keyword,

	"FLUSH":     Keyword,
	"FREELIST":  Keyword,
	"FREELISTS": Keyword,

	//# groups seems too common as table name
	//# "GROUPS": Keyword,

	"INDICATOR": Keyword,
	"INITRANS":  Keyword,
	"INSTANCE":  Keyword,

	"LAYER":   Keyword,
	"LINK":    Keyword,
	"LISTS":   Keyword,
	"LOGFILE": Keyword,

	"MANAGE":        Keyword,
	"MANUAL":        Keyword,
	"MAXDATAFILES":  Keyword,
	"MAXINSTANCES":  Keyword,
	"MAXLOGFILES":   Keyword,
	"MAXLOGHISTORY": Keyword,
	"MAXLOGMEMBERS": Keyword,
	"MAXTRANS":      Keyword,
	"MINEXTENTS":    Keyword,
	"MODULE":        Keyword,
	"MOUNT":         Keyword,

	"NOARCHIVELOG": Keyword,
	"NOCACHE":      Keyword,
	"NOCYCLE":      Keyword,
	"NOMAXVALUE":   Keyword,
	"NOMINVALUE":   Keyword,
	"NOORDER":      Keyword,
	"NORESETLOGS":  Keyword,
	"NORMAL":       Keyword,
	"NOSORT":       Keyword,

	"OPTIMAL": Keyword,
	"OWN":     Keyword,

	"PACKAGE":     Keyword,
	"PARALLEL":    Keyword,
	"PCTINCREASE": Keyword,
	"PCTUSED":     Keyword,
	"PLAN":        Keyword,
	"PRIVATE":     Keyword,
	"PROFILE":     Keyword,

	"QUOTA": Keyword,

	"RECOVER":    Keyword,
	"RESETLOGS":  Keyword,
	"RESTRICTED": Keyword,
	"REUSE":      Keyword,
	"ROLES":      Keyword,

	"SAVEPOINT":    Keyword,
	"SCN":          Keyword,
	"SECTION":      Keyword,
	"SEGMENT":      Keyword,
	"SHARED":       Keyword,
	"SNAPSHOT":     Keyword,
	"SORT":         Keyword,
	"STATEMENT_ID": Keyword,
	"STOP":         Keyword,
	"SWITCH":       Keyword,

	"TABLES":      Keyword,
	"TABLESPACE":  Keyword,
	"THREAD":      Keyword,
	"TIME":        Keyword,
	"TRACING":     Keyword,
	"TRANSACTION": Keyword,
	"TRIGGERS":    Keyword,

	"UNLIMITED": Keyword,
	"UNLOCK":    Keyword,
}

// PostgreSQL Syntax
var KEYWORDS_PLPGSQL = map[string]*TokenType{
	"CONFLICT":      Keyword,
	"WINDOW":        Keyword,
	"PARTITION":     Keyword,
	"OVER":          Keyword,
	"PERFORM":       Keyword,
	"NOTICE":        Keyword,
	"PLPGSQL":       Keyword,
	"INHERIT":       Keyword,
	"INDEXES":       Keyword,
	"ON_ERROR_STOP": Keyword,

	"BYTEA":             Keyword,
	"BIGSERIAL":         Keyword,
	"BIT VARYING":       Keyword,
	"BOX":               Keyword,
	"CHARACTER":         Keyword,
	"CHARACTER VARYING": Keyword,
	"CIDR":              Keyword,
	"CIRCLE":            Keyword,
	"DOUBLE PRECISION":  Keyword,
	"INET":              Keyword,
	"JSON":              Keyword,
	"JSONB":             Keyword,
	"LINE":              Keyword,
	"LSEG":              Keyword,
	"MACADDR":           Keyword,
	"MONEY":             Keyword,
	"PATH":              Keyword,
	"PG_LSN":            Keyword,
	"POINT":             Keyword,
	"POLYGON":           Keyword,
	"SMALLSERIAL":       Keyword,
	"TSQUERY":           Keyword,
	"TSVECTOR":          Keyword,
	"TXID_SNAPSHOT":     Keyword,
	"UUID":              Keyword,
	"XML":               Keyword,

	"FOR":  Keyword,
	"IN":   Keyword,
	"LOOP": Keyword,
}

// Hive Syntax
var KEYWORDS_HQL = map[string]*TokenType{
	"EXPLODE":    Keyword,
	"DIRECTORY":  Keyword,
	"DISTRIBUTE": Keyword,
	"INCLUDE":    Keyword,
	"LOCATE":     Keyword,
	"OVERWRITE":  Keyword,
	"POSEXPLODE": Keyword,

	"ARRAY_CONTAINS":  Keyword,
	"CMP":             Keyword,
	"COLLECT_LIST":    Keyword,
	"CONCAT":          Keyword,
	"CONDITION":       Keyword,
	"DATE_ADD":        Keyword,
	"DATE_SUB":        Keyword,
	"DECODE":          Keyword,
	"DBMS_OUTPUT":     Keyword,
	"ELEMENTS":        Keyword,
	"EXCHANGE":        Keyword,
	"EXTENDED":        Keyword,
	"FLOOR":           Keyword,
	"FOLLOWING":       Keyword,
	"FROM_UNIXTIME":   Keyword,
	"FTP":             Keyword,
	"HOUR":            Keyword,
	"INLINE":          Keyword,
	"INSTR":           Keyword,
	"LEN":             Keyword,
	"MAP":             Builtin,
	"MAXELEMENT":      Keyword,
	"MAXINDEX":        Keyword,
	"MAX_PART_DATE":   Keyword,
	"MAX_PART_INT":    Keyword,
	"MAX_PART_STRING": Keyword,
	"MINELEMENT":      Keyword,
	"MININDEX":        Keyword,
	"MIN_PART_DATE":   Keyword,
	"MIN_PART_INT":    Keyword,
	"MIN_PART_STRING": Keyword,
	"NOW":             Keyword,
	"NVL":             Keyword,
	"NVL2":            Keyword,
	"PARSE_URL_TUPLE": Keyword,
	"PART_LOC":        Keyword,
	"PART_COUNT":      Keyword,
	"PART_COUNT_BY":   Keyword,
	"PRINT":           Keyword,
	"PUT_LINE":        Keyword,
	"RANGE":           Keyword,
	"REDUCE":          Keyword,
	"REGEXP_REPLACE":  Keyword,
	"RESIGNAL":        Keyword,
	"RTRIM":           Keyword,
	"SIGN":            Keyword,
	"SIGNAL":          Keyword,
	"SIN":             Keyword,
	"SPLIT":           Keyword,
	"SQRT":            Keyword,
	"STACK":           Keyword,
	"STR":             Keyword,
	"STRING":          Builtin,
	"STRUCT":          Builtin,
	"SUBSTR":          Keyword,
	"SUMMARY":         Keyword,
	"TBLPROPERTIES":   Keyword,
	"TIMESTAMP":       Builtin,
	"TIMESTAMP_ISO":   Keyword,
	"TO_CHAR":         Keyword,
	"TO_DATE":         Keyword,
	"TO_TIMESTAMP":    Keyword,
	"TRUNC":           Keyword,
	"UNBOUNDED":       Keyword,
	"UNIQUEJOIN":      Keyword,
	"UNIX_TIMESTAMP":  Keyword,
	"UTC_TIMESTAMP":   Keyword,
	"VIEWS":           Keyword,

	"EXIT":  Keyword,
	"BREAK": Keyword,
	"LEAVE": Keyword,
}

var KEYWORDS_MSACCESS = map[string]*TokenType{
	"DISTINCTROW": Keyword,
}
