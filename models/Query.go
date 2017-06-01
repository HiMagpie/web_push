package models
import (
	"strings"
	"strconv"
)
const CommaSpace = ", "
const KEY_SELECT = "select"
const KEY_FROM = "from"
const KEY_JOIN = "join"
const KEY_WHERE = "where"
const KEY_GROUP = "group"
const KEY_HAVING = "having"
const KEY_ORDER = "order"
const KEY_LIMIT = "limit"
const KEY_OFFSET = "offset"
const KEY_UPDATE = "update"

var tokenKeys = []string{
	KEY_SELECT,
	KEY_FROM,
	KEY_JOIN,
	KEY_WHERE,
	KEY_GROUP,
	KEY_HAVING,
	KEY_ORDER,
	KEY_LIMIT,
	KEY_OFFSET,
}

type Query struct {
	//待绑定的参数
	args   []interface{}
	//组装的sql
	Tokens map[string]string
}

//Query的构造函数
func NewQuery() *Query {
	query := new(Query)
	query.Tokens = make(map[string]string)
	return query
}

//Select("*") Select("id,name,age")
func (this *Query ) Select(field string) *Query {
	this.Tokens[KEY_SELECT] = "SELECT " + field
	return this
}

//From("table") From("table as t")
func (this *Query ) From(table string) *Query {
	this.Tokens[KEY_FROM] = "FROM " + table
	return this
}

//Join("table1 as t1","t1.tid=t.id")
func (this *Query ) Join(table string, cond string) *Query {
	joinStr, ok := this.Tokens[KEY_JOIN]
	if ok {
		this.Tokens[KEY_JOIN] = joinStr + " INNER JOIN " + table + " ON " + cond
	}else {
		this.Tokens[KEY_JOIN] = "INNER JOIN " + table + " ON " + cond
	}
	return this
}

//LeftJoin("table2 as t2","t2.tid=t.id")
func (this *Query ) LeftJoin(table string, cond string) *Query {
	joinStr, ok := this.Tokens[KEY_JOIN]
	if ok {
		this.Tokens[KEY_JOIN] = joinStr + " LEFT JOIN " + table + " ON " + cond
	}else {
		this.Tokens[KEY_JOIN] = "LEFT JOIN " + table + " ON " + cond
	}
	return this
}

//RightJoin("table3 as t3","t3.tid=t.id")
func (this *Query ) RightJoin(table string, cond string) *Query {
	joinStr, ok := this.Tokens[KEY_JOIN]
	if ok {
		this.Tokens[KEY_JOIN] = joinStr + " Right JOIN " + table + " ON " + cond
	}else {
		this.Tokens[KEY_JOIN] = "Right JOIN " + table + " ON " + cond
	}
	return this
}

//WhereCond("id=?",1)
func (this *Query ) WhereCond(cond string, args ...interface{}) *Query {
	whereStr, ok := this.Tokens[KEY_WHERE]
	if ok {
		this.Tokens[KEY_WHERE] = "(" + whereStr + ")" + " AND " + cond
	}else {
		this.Tokens[KEY_WHERE] = cond
	}
	this.args = append(this.args, args...)
	return this
}

//Where("id",1)
func (this *Query ) Where(field string, value interface{}) *Query {
	cond := field + " = ?"
	return this.WhereCond(cond, value)
}

//根据slice生成"?,?,?"的字串
func (this *Query ) genQuestionStr(values []interface{}) string {
	questions := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		questions[i] = "?"
	}
	return strings.Join(questions, ",")
}

//WhereIn("id",[]{1,2,3,4})
func (this *Query ) WhereIn(field string, value interface{}) *Query {
	inters := transToInterfaceSlice(value)
	cond := field + " IN (" + this.genQuestionStr(inters) + ")"
	return this.WhereCond(cond, inters...)
}



//WhereNotIn("id",[]{1,2,3,4})
func (this *Query ) WhereNotIn(field string, value interface{}) *Query {
	inters := transToInterfaceSlice(value)
	cond := field + " NOT IN (" + this.genQuestionStr(inters) + ")"

	return this.WhereCond(cond, inters...)
}

//OrWhereCond("id=?",1)
func (this *Query ) OrWhereCond(cond string, args ...interface{}) *Query {
	whereStr, ok := this.Tokens[KEY_WHERE]
	if ok {
		this.Tokens[KEY_WHERE] = "(" + whereStr + ")" + " OR " + cond
	}else {
		this.Tokens[KEY_WHERE] = cond
	}
	this.args = append(this.args, args...)
	return this
}

//OrWhere("id",1)
func (this *Query ) OrWhere(field string, value interface{}) *Query {
	cond := field + " = ?"
	return this.OrWhereCond(cond, value)
}

// OrWhereIn("id",[]{1,2,3,4})
func (this *Query ) OrWhereIn(field string, value interface{}) *Query {
	inters := transToInterfaceSlice(value)
	cond := field + " IN (" + this.genQuestionStr(inters) + ")"
	return this.OrWhereCond(cond, inters...)
}

// OrWhereNotIn("id",[]{1,2,3,4})
func (this *Query ) OrWhereNotIn(field string, value interface{}) *Query {
	inters := transToInterfaceSlice(value)
	cond := field + " NOT IN (" + this.genQuestionStr(inters) + ")"
	return this.OrWhereCond(cond, inters...)
}

//GroupBy("id,name") GroupBy("id","name")
func (this *Query ) GroupBy(fields ...string) *Query {
	this.Tokens[KEY_GROUP] = "GROUP BY " + strings.Join(fields, CommaSpace)
	return this
}

// Having("count(a) > 2")
func (this *Query ) Having(cond string) *Query {
	this.Tokens[KEY_HAVING] = "HAVING " + cond
	return this
}

//Limit(10)
func (this *Query ) Limit(limit int) *Query {
	this.Tokens[KEY_LIMIT] = "LIMIT " + strconv.Itoa(limit)
	return this
}

//Offset(10)
func (this *Query ) Offset(offset int) *Query {
	this.Tokens[KEY_OFFSET] = "OFFSET " + strconv.Itoa(offset)
	return this
}

//OrderBy("id desc","ctime asc")
func (this *Query ) OrderBy(fields ...string) *Query {
	this.Tokens[KEY_ORDER] = "ORDER BY " + strings.Join(fields, CommaSpace)
	return this
}

//Update("table",map{"count":3})
func (this *Query ) Update(table string,columns map[string]interface{}) *Query {
	sql := "UPDATE "+table+" SET"
	for field,val := range columns{
		sql = sql+" "+field+"=?,"
		this.args = append(this.args,val)
	}
	this.Tokens[KEY_UPDATE] = strings.TrimRight(sql, ",")
	return this
}

//返回组装后的sql语句
func (this *Query ) String() string {
	//如果是执行的修改操作
	if updateSql,ok := this.Tokens[KEY_UPDATE];ok{
		if whereStr, ok := this.Tokens[KEY_WHERE]; ok {
			updateSql = updateSql+" WHERE " + whereStr
		}
		return updateSql
	}
	//执行的是查询操作
	var sql []string
	if _, ok := this.Tokens[KEY_SELECT]; !ok {
		this.Tokens[KEY_SELECT] = "SELECT *"
	}
	if whereStr, ok := this.Tokens[KEY_WHERE]; ok {
		this.Tokens[KEY_WHERE] = "WHERE " + whereStr
	}
	for _, v := range tokenKeys {
		if key, ok := this.Tokens[v]; ok {
			sql = append(sql, key)
		}
	}
	return strings.Join(sql, " ")
}

//返回query的args参数
func (this *Query ) GetArgs() []interface{} {
	return this.args
}


func transToInterfaceSlice(value interface{}) []interface{} {
	inters := make([]interface{}, 0)
	switch tmp := value.(type) {
	case []string:
		for _, v := range tmp {
			inters = append(inters, v)
		}
	case []int:
		for _, v := range tmp {
			inters = append(inters, v)
		}
	default:
		// @TODO 无法识别类型
		panic("无法识别类型")
	}

	return inters
}


