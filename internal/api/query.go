package api

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	entity     string
	selectExpr string
	where      string
	orderBy    string
	startPos   int
	maxResults int
}

func NewQuery(entity string) *QueryBuilder {
	return &QueryBuilder{
		entity:     entity,
		selectExpr: "*",
		maxResults: 1000,
		startPos:   1,
	}
}

func (q *QueryBuilder) Select(fields string) *QueryBuilder {
	if fields != "" {
		q.selectExpr = fields
	}
	return q
}

func (q *QueryBuilder) Where(clause string) *QueryBuilder {
	q.where = clause
	return q
}

func (q *QueryBuilder) OrderBy(expr string) *QueryBuilder {
	q.orderBy = expr
	return q
}

func (q *QueryBuilder) StartPosition(pos int) *QueryBuilder {
	q.startPos = pos
	return q
}

func (q *QueryBuilder) MaxResults(n int) *QueryBuilder {
	if n > 0 {
		q.maxResults = n
	}
	return q
}

func (q *QueryBuilder) Build() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "SELECT %s FROM %s", q.selectExpr, q.entity)
	if q.where != "" {
		fmt.Fprintf(&sb, " WHERE %s", q.where)
	}
	if q.orderBy != "" {
		fmt.Fprintf(&sb, " ORDERBY %s", q.orderBy)
	}
	if q.startPos > 1 {
		fmt.Fprintf(&sb, " STARTPOSITION %d", q.startPos)
	}
	if q.maxResults > 0 {
		fmt.Fprintf(&sb, " MAXRESULTS %d", q.maxResults)
	}
	return sb.String()
}
