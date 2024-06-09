// this file follows up mapper.go

package parser

import "github.com/sqldef/sqldef/parser"

func mapExpr(expr parser.Expr) Expr {
	switch e := expr.(type) {
	case *parser.AndExpr:
		return mapAndExprPtr(e)
	case *parser.OrExpr:
		return mapOrExprPtr(e)
	case *parser.NotExpr:
		return mapNotExprPtr(e)
	case *parser.ParenExpr:
		return mapParenExprPtr(e)
	case *parser.ComparisonExpr:
		return mapComparisonExprPtr(e)
	case *parser.RangeCond:
		return mapRangeCondPtr(e)
	case *parser.IsExpr:
		return mapIsExprPtr(e)
	case *parser.ExistsExpr:
		return mapExistsExprPtr(e)
	case *parser.SQLVal:
		return mapSQLValPtr(e)
	case *parser.NullVal:
		return mapNullValPtr(e)
	case parser.BoolVal:
		return mapBoolVal(e)
	case *parser.ColName:
		return mapColNamePtr(e)
	case *parser.NewQualifierColName:
		return mapNewQualifierColNamePtr(e)
	case parser.ValTuple:
		return mapExprSlice(e)
	case *parser.Subquery:
		return mapSubqueryPtr(e)
	case *parser.ListArg:
		return e
	case *parser.BinaryExpr:
		return mapBinaryExprPtr(e)
	case *parser.UnaryExpr:
		return mapUnaryExprPtr(e)
	case *parser.IntervalExpr:
		return mapIntervalExprPtr(e)
	case *parser.CollateExpr:
		return mapCollateExprPtr(e)
	case *parser.FuncExpr:
		return mapFuncExprPtr(e)
	case *parser.CaseExpr:
		return mapCaseExprPtr(e)
	case *parser.ValuesFuncExpr:
		return mapValuesFuncExprPtr(e)
	case *parser.CastExpr:
		return mapCastExprPtr(e)
	case *parser.ConvertExpr:
		return mapConvertExprPtr(e)
	case *parser.SubstrExpr:
		return mapSubstrExprPtr(e)
	case *parser.ConvertUsingExpr:
		return mapConvertUsingExprPtr(e)
	case *parser.MatchExpr:
		return mapMatchExprPtr(e)
	case *parser.GroupConcatExpr:
		return mapGroupConcatExprPtr(e)
	case *parser.OverExpr:
		return mapOverExprPtr(e)
	case *parser.Default:
		return mapDefaultPtr(e)
	case *parser.ArrayConstructor:
		return mapArrayConstructorPtr(e)
	case *parser.FuncCallExpr:
		return mapFuncCallExprPtr(e)
	case *parser.NextSeqValExpr:
		return mapNextSeqValExprPtr(e)
	default:
		panic("unsupported expr")
	}
}

func mapSimpleTableExpr(expr parser.SimpleTableExpr) SimpleTableExpr {
	switch e := expr.(type) {
	case *parser.TableName:
		return mapTableNamePtr(e)
	case *parser.Subquery:
		return mapSubqueryPtr(e)
	default:
		panic("unsupported simple table expr")
	}
}

func mapStatement(stmt parser.Statement) Statement {
	switch s := stmt.(type) {
	case *parser.Union:
		return mapUnionPtr(s)
	case *parser.Select:
		return mapSelectPtr(s)
	case *parser.Stream:
		return mapStreamPtr(s)
	case *parser.Insert:
		return mapInsertPtr(s)
	case *parser.Update:
		return mapUpdatePtr(s)
	case *parser.Delete:
		return mapDeletePtr(s)
	case *parser.Set:
		return mapSetPtr(s)
	case *parser.Declare:
		return mapDeclarePtr(s)
	case *parser.Cursor:
		return mapCursorPtr(s)
	case *parser.BeginEnd:
		return mapBeginEndPtr(s)
	case *parser.While:
		return mapWhilePtr(s)
	case *parser.If:
		return mapIfPtr(s)
	case *parser.DDL:
		return mapDDLPtr(s)
	case *parser.Show:
		return mapShowPtr(s)
	case *parser.Use:
		return mapUsePtr(s)
	case *parser.Begin:
		return mapBeginPtr(s)
	case *parser.Commit:
		return mapCommitPtr(s)
	case *parser.Rollback:
		return mapRollbackPtr(s)
	case *parser.OtherRead:
		return mapOtherReadPtr(s)
	case *parser.OtherAdmin:
		return mapOtherAdminPtr(s)
	case *parser.SetBoolOption:
		return mapSetBoolOptionPtr(s)
	default:
		panic("unsupported statement")
	}
}

func mapSelectStatement(stmt parser.SelectStatement) SelectStatement {
	switch s := stmt.(type) {
	case *parser.Select:
		return mapSelectPtr(s)
	case *parser.Union:
		return mapUnionPtr(s)
	case *parser.ParenSelect:
		return mapParenSelectPtr(s)
	default:
		panic("unsupported select statement")
	}
}

func mapSelectExpr(expr parser.SelectExpr) SelectExpr {
	switch e := expr.(type) {
	case *parser.StarExpr:
		return mapStarExprPtr(e)
	case *parser.AliasedExpr:
		return mapAliasedExprPtr(e)
	default:
		panic("unsupported select expr")
	}
}

func mapInsertRows(rows parser.InsertRows) InsertRows {
	switch r := rows.(type) {
	case *parser.Select:
		return mapSelectPtr(r)
	case *parser.Union:
		return mapUnionPtr(r)
	case parser.Values:
		return r
	case *parser.ParenSelect:
		return mapParenSelectPtr(r)
	default:
		panic("unsupported insert rows")
	}
}

func mapTableExpr(expr parser.TableExpr) TableExpr {
	switch e := expr.(type) {
	case *parser.AliasedTableExpr:
		return mapAliasedTableExprPtr(e)
	case *parser.JoinTableExpr:
		return mapJoinTableExprPtr(e)
	case *parser.ParenTableExpr:
		return mapParenTableExprPtr(e)
	default:
		panic("unsupported table expr")
	}
}

func mapArrayElement(elem parser.ArrayElement) ArrayElement {
	switch e := elem.(type) {
	case *parser.SQLVal:
		return mapSQLValPtr(e)
	case *parser.CastExpr:
		return mapCastExprPtr(e)
	default:
		panic("unsupported array element")
	}
}

type ValTuple []Expr

// todo: fix name
func mapExprSlicePtr(exprs *parser.ValTuple) *[]Expr {
	result := make([]Expr, len(*exprs))
	for i, expr := range *exprs {
		result[i] = mapExpr(expr)
	}
	return &result
}
