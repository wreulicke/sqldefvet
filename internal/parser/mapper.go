// Code generated by tools/generator.go. DO NOT EDIT.

package parser

import (
	"github.com/sqldef/sqldef/parser"
)

type ColIdent struct {
	Val string
}

func mapColIdentPtr(src *parser.ColIdent) *ColIdent {
	return &ColIdent{
		Val: src.String(),
	}
}

type TableIdent struct {
	Val string
}

func mapTableIdentPtr(src *parser.TableIdent) *TableIdent {
	return &TableIdent{
		Val: src.String(),
	}
}

type DDL struct {
	Action DDLAction
	Table TableName
	NewName TableName
	IfExists bool
	TableSpec *TableSpec
	PartitionSpec *PartitionSpec
	IndexSpec *IndexSpec
	IndexCols []IndexColumn
	IndexExpr Expr
	ForeignKey *ForeignKeyDefinition
	Policy *Policy
	View *View
	Trigger *Trigger
	Type *Type
	Comment *Comment
	Extension *Extension
}

func mapDDLPtr(src *parser.DDL) *DDL {
	if src == nil {
		return nil
	}
	r := &DDL{
		Action: DDLAction(src.Action),
		Table: *mapTableNamePtr(&src.Table),
		NewName: *mapTableNamePtr(&src.NewName),
		IfExists: src.IfExists,
		TableSpec: mapTableSpecPtr(src.TableSpec),
		PartitionSpec: mapPartitionSpecPtr(src.PartitionSpec),
		IndexSpec: mapIndexSpecPtr(src.IndexSpec),
		IndexCols: mapIndexColumnSlice(src.IndexCols),
		IndexExpr: mapExpr(src.IndexExpr),
		ForeignKey: mapForeignKeyDefinitionPtr(src.ForeignKey),
		Policy: mapPolicyPtr(src.Policy),
		View: mapViewPtr(src.View),
		Trigger: mapTriggerPtr(src.Trigger),
		Type: mapTypePtr(src.Type),
		Comment: mapCommentPtr(src.Comment),
		Extension: mapExtensionPtr(src.Extension),
	}
	return r
}

type DDLAction int

func mapDDLAction(src parser.DDLAction) DDLAction {
	return DDLAction(src)
}

type TableName struct {
	Schema TableIdent
	Name TableIdent
}

func mapTableNamePtr(src *parser.TableName) *TableName {
	if src == nil {
		return nil
	}
	r := &TableName{
		Schema: *mapTableIdentPtr(&src.Schema),
		Name: *mapTableIdentPtr(&src.Name),
	}
	return r
}

func mapTableSpecPtr(src *parser.TableSpec) *TableSpec {
	if src == nil {
		return nil
	}
	r := &TableSpec{
		Columns: mapColumnDefinitionPtrSlice(src.Columns),
		Indexes: mapIndexDefinitionPtrSlice(src.Indexes),
		ForeignKeys: mapForeignKeyDefinitionPtrSlice(src.ForeignKeys),
		Checks: mapCheckDefinitionPtrSlice(src.Checks),
	}
	return r
}

type TableSpec struct {
	Columns []*ColumnDefinition
	Indexes []*IndexDefinition
	ForeignKeys []*ForeignKeyDefinition
	Checks []*CheckDefinition
}

func mapColumnDefinitionPtrSlice(src []*parser.ColumnDefinition) []*ColumnDefinition {
	if src == nil {
		return nil
	}
	result := make([]*ColumnDefinition, 0, len(src))
	for _, v := range src {
		result = append(result, mapColumnDefinitionPtr(v))
	}
	return result
}

func mapColumnDefinitionPtr(src *parser.ColumnDefinition) *ColumnDefinition {
	if src == nil {
		return nil
	}
	r := &ColumnDefinition{
		Name: *mapColIdentPtr(&src.Name),
		Type: *mapColumnTypePtr(&src.Type),
	}
	return r
}

type ColumnDefinition struct {
	Name ColIdent
	Type ColumnType
}

type ColumnType struct {
	Type string
	NotNull *BoolVal
	Autoincrement BoolVal
	Default *DefaultDefinition
	Srid *SridDefinition
	OnUpdate *SQLVal
	Comment *SQLVal
	Check *CheckDefinition
	Array BoolVal
	Length *SQLVal
	Unsigned BoolVal
	Zerofill BoolVal
	Scale *SQLVal
	DisplayWidth *SQLVal
	Charset string
	Collate string
	Timezone BoolVal
	EnumValues []string
	KeyOpt ColumnKeyOption
	References string
	ReferenceNames []ColIdent
	ReferenceOnDelete ColIdent
	ReferenceOnUpdate ColIdent
	Generated *GeneratedColumn
	Identity *IdentityOpt
}

func mapColumnTypePtr(src *parser.ColumnType) *ColumnType {
	if src == nil {
		return nil
	}
	r := &ColumnType{
		Type: src.Type,
		NotNull: mapBoolValPtr(src.NotNull),
		Autoincrement: BoolVal(src.Autoincrement),
		Default: mapDefaultDefinitionPtr(src.Default),
		Srid: mapSridDefinitionPtr(src.Srid),
		OnUpdate: mapSQLValPtr(src.OnUpdate),
		Comment: mapSQLValPtr(src.Comment),
		Check: mapCheckDefinitionPtr(src.Check),
		Array: BoolVal(src.Array),
		Length: mapSQLValPtr(src.Length),
		Unsigned: BoolVal(src.Unsigned),
		Zerofill: BoolVal(src.Zerofill),
		Scale: mapSQLValPtr(src.Scale),
		DisplayWidth: mapSQLValPtr(src.DisplayWidth),
		Charset: src.Charset,
		Collate: src.Collate,
		Timezone: BoolVal(src.Timezone),
		EnumValues: src.EnumValues,
		KeyOpt: ColumnKeyOption(src.KeyOpt),
		References: src.References,
		ReferenceNames: mapColIdentSlice(src.ReferenceNames),
		ReferenceOnDelete: *mapColIdentPtr(&src.ReferenceOnDelete),
		ReferenceOnUpdate: *mapColIdentPtr(&src.ReferenceOnUpdate),
		Generated: mapGeneratedColumnPtr(src.Generated),
		Identity: mapIdentityOptPtr(src.Identity),
	}
	return r
}

func mapBoolValPtr(src *parser.BoolVal) *BoolVal {
	if src == nil {
		return nil
	}
	r :=BoolVal(*src)
	return &r
}

type BoolVal bool

func mapBoolVal(src parser.BoolVal) BoolVal {
	return BoolVal(src)
}

func mapDefaultDefinitionPtr(src *parser.DefaultDefinition) *DefaultDefinition {
	if src == nil {
		return nil
	}
	r := &DefaultDefinition{
		ValueOrExpression: *mapDefaultValueOrExpressionPtr(&src.ValueOrExpression),
		ConstraintName: *mapColIdentPtr(&src.ConstraintName),
	}
	return r
}

type DefaultDefinition struct {
	ValueOrExpression DefaultValueOrExpression
	ConstraintName ColIdent
}

type DefaultValueOrExpression struct {
	Value *SQLVal
	Expr Expr
}

func mapDefaultValueOrExpressionPtr(src *parser.DefaultValueOrExpression) *DefaultValueOrExpression {
	if src == nil {
		return nil
	}
	r := &DefaultValueOrExpression{
		Value: mapSQLValPtr(src.Value),
		Expr: mapExpr(src.Expr),
	}
	return r
}

func mapSQLValPtr(src *parser.SQLVal) *SQLVal {
	if src == nil {
		return nil
	}
	r := &SQLVal{
		Type: ValType(src.Type),
		Val: src.Val,
	}
	return r
}

type SQLVal struct {
	Type ValType
	Val []uint8
}

type ValType int

func mapValType(src parser.ValType) ValType {
	return ValType(src)
}

type Expr interface {
}

func mapSridDefinitionPtr(src *parser.SridDefinition) *SridDefinition {
	if src == nil {
		return nil
	}
	r := &SridDefinition{
		Value: mapSQLValPtr(src.Value),
	}
	return r
}

type SridDefinition struct {
	Value *SQLVal
}

func mapCheckDefinitionPtr(src *parser.CheckDefinition) *CheckDefinition {
	if src == nil {
		return nil
	}
	r := &CheckDefinition{
		Where: *mapWherePtr(&src.Where),
		ConstraintName: *mapColIdentPtr(&src.ConstraintName),
		NotForReplication: src.NotForReplication,
		NoInherit: BoolVal(src.NoInherit),
	}
	return r
}

type CheckDefinition struct {
	Where Where
	ConstraintName ColIdent
	NotForReplication bool
	NoInherit BoolVal
}

type Where struct {
	Type string
	Expr Expr
}

func mapWherePtr(src *parser.Where) *Where {
	if src == nil {
		return nil
	}
	r := &Where{
		Type: src.Type,
		Expr: mapExpr(src.Expr),
	}
	return r
}

type ColumnKeyOption int

func mapColumnKeyOption(src parser.ColumnKeyOption) ColumnKeyOption {
	return ColumnKeyOption(src)
}

func mapColIdentSlice(src []parser.ColIdent) []ColIdent {
	if src == nil {
		return nil
	}
	result := make([]ColIdent, 0, len(src))
	for _, v := range src {
		result = append(result, *mapColIdentPtr(&v))
	}
	return result
}

func mapGeneratedColumnPtr(src *parser.GeneratedColumn) *GeneratedColumn {
	if src == nil {
		return nil
	}
	r := &GeneratedColumn{
		Expr: mapExpr(src.Expr),
		GeneratedType: src.GeneratedType,
	}
	return r
}

type GeneratedColumn struct {
	Expr Expr
	GeneratedType string
}

func mapIdentityOptPtr(src *parser.IdentityOpt) *IdentityOpt {
	if src == nil {
		return nil
	}
	r := &IdentityOpt{
		Behavior: src.Behavior,
		Sequence: mapSequencePtr(src.Sequence),
		NotForReplication: src.NotForReplication,
	}
	return r
}

type IdentityOpt struct {
	Behavior string
	Sequence *Sequence
	NotForReplication bool
}

func mapSequencePtr(src *parser.Sequence) *Sequence {
	if src == nil {
		return nil
	}
	r := &Sequence{
		Name: src.Name,
		IfNotExists: src.IfNotExists,
		Type: src.Type,
		IncrementBy: mapSQLValPtr(src.IncrementBy),
		MinValue: mapSQLValPtr(src.MinValue),
		NoMinValue: mapBoolValPtr(src.NoMinValue),
		MaxValue: mapSQLValPtr(src.MaxValue),
		NoMaxValue: mapBoolValPtr(src.NoMaxValue),
		StartWith: mapSQLValPtr(src.StartWith),
		Cache: mapSQLValPtr(src.Cache),
		Cycle: mapBoolValPtr(src.Cycle),
		NoCycle: mapBoolValPtr(src.NoCycle),
		OwnedBy: src.OwnedBy,
	}
	return r
}

type Sequence struct {
	Name string
	IfNotExists bool
	Type string
	IncrementBy *SQLVal
	MinValue *SQLVal
	NoMinValue *BoolVal
	MaxValue *SQLVal
	NoMaxValue *BoolVal
	StartWith *SQLVal
	Cache *SQLVal
	Cycle *BoolVal
	NoCycle *BoolVal
	OwnedBy string
}

func mapIndexDefinitionPtrSlice(src []*parser.IndexDefinition) []*IndexDefinition {
	if src == nil {
		return nil
	}
	result := make([]*IndexDefinition, 0, len(src))
	for _, v := range src {
		result = append(result, mapIndexDefinitionPtr(v))
	}
	return result
}

func mapIndexDefinitionPtr(src *parser.IndexDefinition) *IndexDefinition {
	if src == nil {
		return nil
	}
	r := &IndexDefinition{
		Info: mapIndexInfoPtr(src.Info),
		Columns: mapIndexColumnSlice(src.Columns),
		Options: mapIndexOptionPtrSlice(src.Options),
		Partition: mapIndexPartitionPtr(src.Partition),
		ConstraintOptions: mapConstraintOptionsPtr(src.ConstraintOptions),
	}
	return r
}

type IndexDefinition struct {
	Info *IndexInfo
	Columns []IndexColumn
	Options []*IndexOption
	Partition *IndexPartition
	ConstraintOptions *ConstraintOptions
}

func mapIndexInfoPtr(src *parser.IndexInfo) *IndexInfo {
	if src == nil {
		return nil
	}
	r := &IndexInfo{
		Type: src.Type,
		Name: *mapColIdentPtr(&src.Name),
		Primary: src.Primary,
		Spatial: src.Spatial,
		Unique: src.Unique,
		Fulltext: src.Fulltext,
		Clustered: BoolVal(src.Clustered),
	}
	return r
}

type IndexInfo struct {
	Type string
	Name ColIdent
	Primary bool
	Spatial bool
	Unique bool
	Fulltext bool
	Clustered BoolVal
}

func mapIndexColumnSlice(src []parser.IndexColumn) []IndexColumn {
	if src == nil {
		return nil
	}
	result := make([]IndexColumn, 0, len(src))
	for _, v := range src {
		result = append(result, *mapIndexColumnPtr(&v))
	}
	return result
}

type IndexColumn struct {
	Column ColIdent
	Length *SQLVal
	Direction string
	OperatorClass string
}

func mapIndexColumnPtr(src *parser.IndexColumn) *IndexColumn {
	if src == nil {
		return nil
	}
	r := &IndexColumn{
		Column: *mapColIdentPtr(&src.Column),
		Length: mapSQLValPtr(src.Length),
		Direction: src.Direction,
		OperatorClass: src.OperatorClass,
	}
	return r
}

func mapIndexOptionPtrSlice(src []*parser.IndexOption) []*IndexOption {
	if src == nil {
		return nil
	}
	result := make([]*IndexOption, 0, len(src))
	for _, v := range src {
		result = append(result, mapIndexOptionPtr(v))
	}
	return result
}

func mapIndexOptionPtr(src *parser.IndexOption) *IndexOption {
	if src == nil {
		return nil
	}
	r := &IndexOption{
		Name: src.Name,
		Value: mapSQLValPtr(src.Value),
	}
	return r
}

type IndexOption struct {
	Name string
	Value *SQLVal
}

func mapIndexPartitionPtr(src *parser.IndexPartition) *IndexPartition {
	if src == nil {
		return nil
	}
	r := &IndexPartition{
		Name: src.Name,
		Column: src.Column,
	}
	return r
}

type IndexPartition struct {
	Name string
	Column string
}

func mapConstraintOptionsPtr(src *parser.ConstraintOptions) *ConstraintOptions {
	if src == nil {
		return nil
	}
	r := &ConstraintOptions{
		Deferrable: src.Deferrable,
		InitiallyDeferred: src.InitiallyDeferred,
	}
	return r
}

type ConstraintOptions struct {
	Deferrable bool
	InitiallyDeferred bool
}

func mapForeignKeyDefinitionPtrSlice(src []*parser.ForeignKeyDefinition) []*ForeignKeyDefinition {
	if src == nil {
		return nil
	}
	result := make([]*ForeignKeyDefinition, 0, len(src))
	for _, v := range src {
		result = append(result, mapForeignKeyDefinitionPtr(v))
	}
	return result
}

func mapForeignKeyDefinitionPtr(src *parser.ForeignKeyDefinition) *ForeignKeyDefinition {
	if src == nil {
		return nil
	}
	r := &ForeignKeyDefinition{
		ConstraintName: *mapColIdentPtr(&src.ConstraintName),
		IndexName: *mapColIdentPtr(&src.IndexName),
		IndexColumns: mapColIdentSlice(src.IndexColumns),
		ReferenceName: *mapTableNamePtr(&src.ReferenceName),
		ReferenceColumns: mapColIdentSlice(src.ReferenceColumns),
		OnDelete: *mapColIdentPtr(&src.OnDelete),
		OnUpdate: *mapColIdentPtr(&src.OnUpdate),
		NotForReplication: src.NotForReplication,
		ConstraintOptions: mapConstraintOptionsPtr(src.ConstraintOptions),
	}
	return r
}

type ForeignKeyDefinition struct {
	ConstraintName ColIdent
	IndexName ColIdent
	IndexColumns []ColIdent
	ReferenceName TableName
	ReferenceColumns []ColIdent
	OnDelete ColIdent
	OnUpdate ColIdent
	NotForReplication bool
	ConstraintOptions *ConstraintOptions
}

func mapCheckDefinitionPtrSlice(src []*parser.CheckDefinition) []*CheckDefinition {
	if src == nil {
		return nil
	}
	result := make([]*CheckDefinition, 0, len(src))
	for _, v := range src {
		result = append(result, mapCheckDefinitionPtr(v))
	}
	return result
}

func mapPartitionSpecPtr(src *parser.PartitionSpec) *PartitionSpec {
	if src == nil {
		return nil
	}
	r := &PartitionSpec{
		Action: src.Action,
		Name: *mapColIdentPtr(&src.Name),
		Definitions: mapPartitionDefinitionPtrSlice(src.Definitions),
	}
	return r
}

type PartitionSpec struct {
	Action string
	Name ColIdent
	Definitions []*PartitionDefinition
}

func mapPartitionDefinitionPtrSlice(src []*parser.PartitionDefinition) []*PartitionDefinition {
	if src == nil {
		return nil
	}
	result := make([]*PartitionDefinition, 0, len(src))
	for _, v := range src {
		result = append(result, mapPartitionDefinitionPtr(v))
	}
	return result
}

func mapPartitionDefinitionPtr(src *parser.PartitionDefinition) *PartitionDefinition {
	if src == nil {
		return nil
	}
	r := &PartitionDefinition{
		Name: *mapColIdentPtr(&src.Name),
		Limit: mapExpr(src.Limit),
		Maxvalue: src.Maxvalue,
	}
	return r
}

type PartitionDefinition struct {
	Name ColIdent
	Limit Expr
	Maxvalue bool
}

func mapIndexSpecPtr(src *parser.IndexSpec) *IndexSpec {
	if src == nil {
		return nil
	}
	r := &IndexSpec{
		Name: *mapColIdentPtr(&src.Name),
		Type: *mapColIdentPtr(&src.Type),
		Unique: src.Unique,
		Primary: src.Primary,
		Constraint: src.Constraint,
		Clustered: src.Clustered,
		ColumnStore: src.ColumnStore,
		Included: mapColIdentSlice(src.Included),
		Where: mapWherePtr(src.Where),
		Options: mapIndexOptionPtrSlice(src.Options),
		Partition: mapIndexPartitionPtr(src.Partition),
		ConstraintOptions: mapConstraintOptionsPtr(src.ConstraintOptions),
	}
	return r
}

type IndexSpec struct {
	Name ColIdent
	Type ColIdent
	Unique bool
	Primary bool
	Constraint bool
	Clustered bool
	ColumnStore bool
	Included []ColIdent
	Where *Where
	Options []*IndexOption
	Partition *IndexPartition
	ConstraintOptions *ConstraintOptions
}

func mapPolicyPtr(src *parser.Policy) *Policy {
	if src == nil {
		return nil
	}
	r := &Policy{
		Name: *mapColIdentPtr(&src.Name),
		Permissive: Permissive(src.Permissive),
		Scope: src.Scope,
		To: mapColIdentSlice(src.To),
		Using: mapWherePtr(src.Using),
		WithCheck: mapWherePtr(src.WithCheck),
	}
	return r
}

type Policy struct {
	Name ColIdent
	Permissive Permissive
	Scope []uint8
	To []ColIdent
	Using *Where
	WithCheck *Where
}

type Permissive string

func mapPermissive(src parser.Permissive) Permissive {
	return Permissive(src)
}

func mapViewPtr(src *parser.View) *View {
	if src == nil {
		return nil
	}
	r := &View{
		Type: src.Type,
		SecurityType: src.SecurityType,
		Name: *mapTableNamePtr(&src.Name),
		Definition: mapSelectStatement(src.Definition),
	}
	return r
}

type View struct {
	Type string
	SecurityType string
	Name TableName
	Definition SelectStatement
}

type SelectStatement interface {
}

func mapTriggerPtr(src *parser.Trigger) *Trigger {
	if src == nil {
		return nil
	}
	r := &Trigger{
		Name: *mapColIdentPtr(&src.Name),
		TableName: *mapTableNamePtr(&src.TableName),
		Time: src.Time,
		Event: src.Event,
		Body: mapStatementSlice(src.Body),
	}
	return r
}

type Trigger struct {
	Name ColIdent
	TableName TableName
	Time string
	Event []string
	Body []Statement
}

func mapStatementSlice(src []parser.Statement) []Statement {
	if src == nil {
		return nil
	}
	result := make([]Statement, 0, len(src))
	for _, v := range src {
		result = append(result, mapStatement(v))
	}
	return result
}

type Statement interface {
}

func mapTypePtr(src *parser.Type) *Type {
	if src == nil {
		return nil
	}
	r := &Type{
		Name: *mapTableNamePtr(&src.Name),
		Type: *mapColumnTypePtr(&src.Type),
	}
	return r
}

type Type struct {
	Name TableName
	Type ColumnType
}

func mapCommentPtr(src *parser.Comment) *Comment {
	if src == nil {
		return nil
	}
	r := &Comment{
		ObjectType: src.ObjectType,
		Object: src.Object,
		Comment: src.Comment,
	}
	return r
}

type Comment struct {
	ObjectType string
	Object string
	Comment string
}

func mapExtensionPtr(src *parser.Extension) *Extension {
	if src == nil {
		return nil
	}
	r := &Extension{
		Name: src.Name,
	}
	return r
}

type Extension struct {
	Name string
}

type Union struct {
	Type string
	Left SelectStatement
	Right SelectStatement
	OrderBy []*Order
	Limit *Limit
	Lock string
}

func mapUnionPtr(src *parser.Union) *Union {
	if src == nil {
		return nil
	}
	r := &Union{
		Type: src.Type,
		Left: mapSelectStatement(src.Left),
		Right: mapSelectStatement(src.Right),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
		Limit: mapLimitPtr(src.Limit),
		Lock: src.Lock,
	}
	return r
}

func mapOrderPtrSlice(src []*parser.Order) []*Order {
	if src == nil {
		return nil
	}
	result := make([]*Order, 0, len(src))
	for _, v := range src {
		result = append(result, mapOrderPtr(v))
	}
	return result
}

func mapOrderPtr(src *parser.Order) *Order {
	if src == nil {
		return nil
	}
	r := &Order{
		Expr: mapExpr(src.Expr),
		Direction: src.Direction,
	}
	return r
}

type Order struct {
	Expr Expr
	Direction string
}

func mapLimitPtr(src *parser.Limit) *Limit {
	if src == nil {
		return nil
	}
	r := &Limit{
		Offset: mapExpr(src.Offset),
		Rowcount: mapExpr(src.Rowcount),
	}
	return r
}

type Limit struct {
	Offset Expr
	Rowcount Expr
}

type Select struct {
	Cache string
	Comments [][]uint8
	Distinct string
	Hints string
	SelectExprs []SelectExpr
	From []TableExpr
	Where *Where
	GroupBy []Expr
	Having *Where
	OrderBy []*Order
	Limit *Limit
	Lock string
}

func mapSelectPtr(src *parser.Select) *Select {
	if src == nil {
		return nil
	}
	r := &Select{
		Cache: src.Cache,
		Comments: src.Comments,
		Distinct: src.Distinct,
		Hints: src.Hints,
		SelectExprs: mapSelectExprSlice(src.SelectExprs),
		From: mapTableExprSlice(src.From),
		Where: mapWherePtr(src.Where),
		GroupBy: mapExprSlice(src.GroupBy),
		Having: mapWherePtr(src.Having),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
		Limit: mapLimitPtr(src.Limit),
		Lock: src.Lock,
	}
	return r
}

func mapSelectExprSlice(src []parser.SelectExpr) []SelectExpr {
	if src == nil {
		return nil
	}
	result := make([]SelectExpr, 0, len(src))
	for _, v := range src {
		result = append(result, mapSelectExpr(v))
	}
	return result
}

type SelectExpr interface {
}

func mapTableExprSlice(src []parser.TableExpr) []TableExpr {
	if src == nil {
		return nil
	}
	result := make([]TableExpr, 0, len(src))
	for _, v := range src {
		result = append(result, mapTableExpr(v))
	}
	return result
}

type TableExpr interface {
}

func mapExprSlice(src []parser.Expr) []Expr {
	if src == nil {
		return nil
	}
	result := make([]Expr, 0, len(src))
	for _, v := range src {
		result = append(result, mapExpr(v))
	}
	return result
}

type Stream struct {
	Comments [][]uint8
	SelectExpr SelectExpr
	Table TableName
}

func mapStreamPtr(src *parser.Stream) *Stream {
	if src == nil {
		return nil
	}
	r := &Stream{
		Comments: src.Comments,
		SelectExpr: mapSelectExpr(src.SelectExpr),
		Table: *mapTableNamePtr(&src.Table),
	}
	return r
}

type Insert struct {
	Action string
	Comments [][]uint8
	Ignore string
	Table TableName
	Partitions []ColIdent
	Columns []ColIdent
	Rows InsertRows
	OnDup []*UpdateExpr
}

func mapInsertPtr(src *parser.Insert) *Insert {
	if src == nil {
		return nil
	}
	r := &Insert{
		Action: src.Action,
		Comments: src.Comments,
		Ignore: src.Ignore,
		Table: *mapTableNamePtr(&src.Table),
		Partitions: mapColIdentSlice(src.Partitions),
		Columns: mapColIdentSlice(src.Columns),
		Rows: mapInsertRows(src.Rows),
		OnDup: mapUpdateExprPtrSlice(src.OnDup),
	}
	return r
}

type InsertRows interface {
}

func mapUpdateExprPtrSlice(src []*parser.UpdateExpr) []*UpdateExpr {
	if src == nil {
		return nil
	}
	result := make([]*UpdateExpr, 0, len(src))
	for _, v := range src {
		result = append(result, mapUpdateExprPtr(v))
	}
	return result
}

func mapUpdateExprPtr(src *parser.UpdateExpr) *UpdateExpr {
	if src == nil {
		return nil
	}
	r := &UpdateExpr{
		Name: mapColNamePtr(src.Name),
		Expr: mapExpr(src.Expr),
	}
	return r
}

type UpdateExpr struct {
	Name *ColName
	Expr Expr
}

func mapColNamePtr(src *parser.ColName) *ColName {
	if src == nil {
		return nil
	}
	r := &ColName{
		Name: *mapColIdentPtr(&src.Name),
		Qualifier: *mapTableNamePtr(&src.Qualifier),
	}
	return r
}

type ColName struct {
	Name ColIdent
	Qualifier TableName
}

type Update struct {
	Comments [][]uint8
	TableExprs []TableExpr
	Exprs []*UpdateExpr
	Where *Where
	OrderBy []*Order
	Limit *Limit
}

func mapUpdatePtr(src *parser.Update) *Update {
	if src == nil {
		return nil
	}
	r := &Update{
		Comments: src.Comments,
		TableExprs: mapTableExprSlice(src.TableExprs),
		Exprs: mapUpdateExprPtrSlice(src.Exprs),
		Where: mapWherePtr(src.Where),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
		Limit: mapLimitPtr(src.Limit),
	}
	return r
}

type Delete struct {
	Comments [][]uint8
	Targets []TableName
	TableExprs []TableExpr
	Partitions []ColIdent
	Where *Where
	OrderBy []*Order
	Limit *Limit
}

func mapDeletePtr(src *parser.Delete) *Delete {
	if src == nil {
		return nil
	}
	r := &Delete{
		Comments: src.Comments,
		Targets: mapTableNameSlice(src.Targets),
		TableExprs: mapTableExprSlice(src.TableExprs),
		Partitions: mapColIdentSlice(src.Partitions),
		Where: mapWherePtr(src.Where),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
		Limit: mapLimitPtr(src.Limit),
	}
	return r
}

func mapTableNameSlice(src []parser.TableName) []TableName {
	if src == nil {
		return nil
	}
	result := make([]TableName, 0, len(src))
	for _, v := range src {
		result = append(result, *mapTableNamePtr(&v))
	}
	return result
}

type Set struct {
	Comments [][]uint8
	Exprs []*SetExpr
	Scope string
}

func mapSetPtr(src *parser.Set) *Set {
	if src == nil {
		return nil
	}
	r := &Set{
		Comments: src.Comments,
		Exprs: mapSetExprPtrSlice(src.Exprs),
		Scope: src.Scope,
	}
	return r
}

func mapSetExprPtrSlice(src []*parser.SetExpr) []*SetExpr {
	if src == nil {
		return nil
	}
	result := make([]*SetExpr, 0, len(src))
	for _, v := range src {
		result = append(result, mapSetExprPtr(v))
	}
	return result
}

func mapSetExprPtr(src *parser.SetExpr) *SetExpr {
	if src == nil {
		return nil
	}
	r := &SetExpr{
		Name: *mapColIdentPtr(&src.Name),
		Expr: mapExpr(src.Expr),
	}
	return r
}

type SetExpr struct {
	Name ColIdent
	Expr Expr
}

type Declare struct {
	Type DeclareType
	Variables []*LocalVariable
	Cursor *CursorDefinition
}

func mapDeclarePtr(src *parser.Declare) *Declare {
	if src == nil {
		return nil
	}
	r := &Declare{
		Type: DeclareType(src.Type),
		Variables: mapLocalVariablePtrSlice(src.Variables),
		Cursor: mapCursorDefinitionPtr(src.Cursor),
	}
	return r
}

type DeclareType int

func mapDeclareType(src parser.DeclareType) DeclareType {
	return DeclareType(src)
}

func mapLocalVariablePtrSlice(src []*parser.LocalVariable) []*LocalVariable {
	if src == nil {
		return nil
	}
	result := make([]*LocalVariable, 0, len(src))
	for _, v := range src {
		result = append(result, mapLocalVariablePtr(v))
	}
	return result
}

func mapLocalVariablePtr(src *parser.LocalVariable) *LocalVariable {
	if src == nil {
		return nil
	}
	r := &LocalVariable{
		Name: *mapColIdentPtr(&src.Name),
		DataType: *mapColumnTypePtr(&src.DataType),
	}
	return r
}

type LocalVariable struct {
	Name ColIdent
	DataType ColumnType
}

func mapCursorDefinitionPtr(src *parser.CursorDefinition) *CursorDefinition {
	if src == nil {
		return nil
	}
	r := &CursorDefinition{
		Name: *mapColIdentPtr(&src.Name),
		Scroll: src.Scroll,
		Select: mapSelectStatement(src.Select),
	}
	return r
}

type CursorDefinition struct {
	Name ColIdent
	Scroll bool
	Select SelectStatement
}

type Cursor struct {
	Action string
	Fetch string
	CursorName ColIdent
	Into ColIdent
}

func mapCursorPtr(src *parser.Cursor) *Cursor {
	if src == nil {
		return nil
	}
	r := &Cursor{
		Action: src.Action,
		Fetch: src.Fetch,
		CursorName: *mapColIdentPtr(&src.CursorName),
		Into: *mapColIdentPtr(&src.Into),
	}
	return r
}

type BeginEnd struct {
	Statements []Statement
}

func mapBeginEndPtr(src *parser.BeginEnd) *BeginEnd {
	if src == nil {
		return nil
	}
	r := &BeginEnd{
		Statements: mapStatementSlice(src.Statements),
	}
	return r
}

type While struct {
	Condition Expr
	Statements []Statement
	Keyword string
}

func mapWhilePtr(src *parser.While) *While {
	if src == nil {
		return nil
	}
	r := &While{
		Condition: mapExpr(src.Condition),
		Statements: mapStatementSlice(src.Statements),
		Keyword: src.Keyword,
	}
	return r
}

type If struct {
	Condition Expr
	IfStatements []Statement
	ElseStatements []Statement
	Keyword string
}

func mapIfPtr(src *parser.If) *If {
	if src == nil {
		return nil
	}
	r := &If{
		Condition: mapExpr(src.Condition),
		IfStatements: mapStatementSlice(src.IfStatements),
		ElseStatements: mapStatementSlice(src.ElseStatements),
		Keyword: src.Keyword,
	}
	return r
}

type Show struct {
	Type string
	OnTable TableName
	ShowTablesOpt *ShowTablesOpt
	Scope string
}

func mapShowPtr(src *parser.Show) *Show {
	if src == nil {
		return nil
	}
	r := &Show{
		Type: src.Type,
		OnTable: *mapTableNamePtr(&src.OnTable),
		ShowTablesOpt: mapShowTablesOptPtr(src.ShowTablesOpt),
		Scope: src.Scope,
	}
	return r
}

func mapShowTablesOptPtr(src *parser.ShowTablesOpt) *ShowTablesOpt {
	if src == nil {
		return nil
	}
	r := &ShowTablesOpt{
		Extended: src.Extended,
		Full: src.Full,
		DbName: src.DbName,
		Filter: mapShowFilterPtr(src.Filter),
	}
	return r
}

type ShowTablesOpt struct {
	Extended string
	Full string
	DbName string
	Filter *ShowFilter
}

func mapShowFilterPtr(src *parser.ShowFilter) *ShowFilter {
	if src == nil {
		return nil
	}
	r := &ShowFilter{
		Like: src.Like,
		Filter: mapExpr(src.Filter),
	}
	return r
}

type ShowFilter struct {
	Like string
	Filter Expr
}

type Use struct {
	DBName TableIdent
}

func mapUsePtr(src *parser.Use) *Use {
	if src == nil {
		return nil
	}
	r := &Use{
		DBName: *mapTableIdentPtr(&src.DBName),
	}
	return r
}

type Begin struct {
}

func mapBeginPtr(src *parser.Begin) *Begin {
	if src == nil {
		return nil
	}
	r := &Begin{
	}
	return r
}

type Commit struct {
}

func mapCommitPtr(src *parser.Commit) *Commit {
	if src == nil {
		return nil
	}
	r := &Commit{
	}
	return r
}

type Rollback struct {
}

func mapRollbackPtr(src *parser.Rollback) *Rollback {
	if src == nil {
		return nil
	}
	r := &Rollback{
	}
	return r
}

type OtherRead struct {
}

func mapOtherReadPtr(src *parser.OtherRead) *OtherRead {
	if src == nil {
		return nil
	}
	r := &OtherRead{
	}
	return r
}

type OtherAdmin struct {
}

func mapOtherAdminPtr(src *parser.OtherAdmin) *OtherAdmin {
	if src == nil {
		return nil
	}
	r := &OtherAdmin{
	}
	return r
}

type SetBoolOption struct {
	OptionNames []string
	Value *SQLVal
}

func mapSetBoolOptionPtr(src *parser.SetBoolOption) *SetBoolOption {
	if src == nil {
		return nil
	}
	r := &SetBoolOption{
		OptionNames: src.OptionNames,
		Value: mapSQLValPtr(src.Value),
	}
	return r
}

type AndExpr struct {
	Left Expr
	Right Expr
}

func mapAndExprPtr(src *parser.AndExpr) *AndExpr {
	if src == nil {
		return nil
	}
	r := &AndExpr{
		Left: mapExpr(src.Left),
		Right: mapExpr(src.Right),
	}
	return r
}

type OrExpr struct {
	Left Expr
	Right Expr
}

func mapOrExprPtr(src *parser.OrExpr) *OrExpr {
	if src == nil {
		return nil
	}
	r := &OrExpr{
		Left: mapExpr(src.Left),
		Right: mapExpr(src.Right),
	}
	return r
}

type NotExpr struct {
	Expr Expr
}

func mapNotExprPtr(src *parser.NotExpr) *NotExpr {
	if src == nil {
		return nil
	}
	r := &NotExpr{
		Expr: mapExpr(src.Expr),
	}
	return r
}

type ParenExpr struct {
	Expr Expr
}

func mapParenExprPtr(src *parser.ParenExpr) *ParenExpr {
	if src == nil {
		return nil
	}
	r := &ParenExpr{
		Expr: mapExpr(src.Expr),
	}
	return r
}

type ComparisonExpr struct {
	Operator string
	Left Expr
	Right Expr
	Escape Expr
	All bool
	Any bool
}

func mapComparisonExprPtr(src *parser.ComparisonExpr) *ComparisonExpr {
	if src == nil {
		return nil
	}
	r := &ComparisonExpr{
		Operator: src.Operator,
		Left: mapExpr(src.Left),
		Right: mapExpr(src.Right),
		Escape: mapExpr(src.Escape),
		All: src.All,
		Any: src.Any,
	}
	return r
}

type RangeCond struct {
	Operator string
	Left Expr
	From Expr
	To Expr
}

func mapRangeCondPtr(src *parser.RangeCond) *RangeCond {
	if src == nil {
		return nil
	}
	r := &RangeCond{
		Operator: src.Operator,
		Left: mapExpr(src.Left),
		From: mapExpr(src.From),
		To: mapExpr(src.To),
	}
	return r
}

type IsExpr struct {
	Operator string
	Expr Expr
}

func mapIsExprPtr(src *parser.IsExpr) *IsExpr {
	if src == nil {
		return nil
	}
	r := &IsExpr{
		Operator: src.Operator,
		Expr: mapExpr(src.Expr),
	}
	return r
}

type ExistsExpr struct {
	Subquery *Subquery
}

func mapExistsExprPtr(src *parser.ExistsExpr) *ExistsExpr {
	if src == nil {
		return nil
	}
	r := &ExistsExpr{
		Subquery: mapSubqueryPtr(src.Subquery),
	}
	return r
}

func mapSubqueryPtr(src *parser.Subquery) *Subquery {
	if src == nil {
		return nil
	}
	r := &Subquery{
		Select: mapSelectStatement(src.Select),
	}
	return r
}

type Subquery struct {
	Select SelectStatement
}

type NullVal struct {
}

func mapNullValPtr(src *parser.NullVal) *NullVal {
	if src == nil {
		return nil
	}
	r := &NullVal{
	}
	return r
}

type NewQualifierColName struct {
	Name ColIdent
}

func mapNewQualifierColNamePtr(src *parser.NewQualifierColName) *NewQualifierColName {
	if src == nil {
		return nil
	}
	r := &NewQualifierColName{
		Name: *mapColIdentPtr(&src.Name),
	}
	return r
}

type BinaryExpr struct {
	Operator string
	Left Expr
	Right Expr
}

func mapBinaryExprPtr(src *parser.BinaryExpr) *BinaryExpr {
	if src == nil {
		return nil
	}
	r := &BinaryExpr{
		Operator: src.Operator,
		Left: mapExpr(src.Left),
		Right: mapExpr(src.Right),
	}
	return r
}

type UnaryExpr struct {
	Operator string
	Expr Expr
}

func mapUnaryExprPtr(src *parser.UnaryExpr) *UnaryExpr {
	if src == nil {
		return nil
	}
	r := &UnaryExpr{
		Operator: src.Operator,
		Expr: mapExpr(src.Expr),
	}
	return r
}

type IntervalExpr struct {
	Expr Expr
	Unit string
}

func mapIntervalExprPtr(src *parser.IntervalExpr) *IntervalExpr {
	if src == nil {
		return nil
	}
	r := &IntervalExpr{
		Expr: mapExpr(src.Expr),
		Unit: src.Unit,
	}
	return r
}

type CollateExpr struct {
	Expr Expr
	Charset string
}

func mapCollateExprPtr(src *parser.CollateExpr) *CollateExpr {
	if src == nil {
		return nil
	}
	r := &CollateExpr{
		Expr: mapExpr(src.Expr),
		Charset: src.Charset,
	}
	return r
}

type FuncExpr struct {
	Qualifier TableIdent
	Name ColIdent
	Distinct bool
	Exprs []SelectExpr
	Over *OverExpr
}

func mapFuncExprPtr(src *parser.FuncExpr) *FuncExpr {
	if src == nil {
		return nil
	}
	r := &FuncExpr{
		Qualifier: *mapTableIdentPtr(&src.Qualifier),
		Name: *mapColIdentPtr(&src.Name),
		Distinct: src.Distinct,
		Exprs: mapSelectExprSlice(src.Exprs),
		Over: mapOverExprPtr(src.Over),
	}
	return r
}

func mapOverExprPtr(src *parser.OverExpr) *OverExpr {
	if src == nil {
		return nil
	}
	r := &OverExpr{
		PartitionBy: mapPartitionPtrSlice(src.PartitionBy),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
	}
	return r
}

type OverExpr struct {
	PartitionBy []*Partition
	OrderBy []*Order
}

func mapPartitionPtrSlice(src []*parser.Partition) []*Partition {
	if src == nil {
		return nil
	}
	result := make([]*Partition, 0, len(src))
	for _, v := range src {
		result = append(result, mapPartitionPtr(v))
	}
	return result
}

func mapPartitionPtr(src *parser.Partition) *Partition {
	if src == nil {
		return nil
	}
	r := &Partition{
		Expr: mapExpr(src.Expr),
	}
	return r
}

type Partition struct {
	Expr Expr
}

type CaseExpr struct {
	Expr Expr
	Whens []*When
	Else Expr
}

func mapCaseExprPtr(src *parser.CaseExpr) *CaseExpr {
	if src == nil {
		return nil
	}
	r := &CaseExpr{
		Expr: mapExpr(src.Expr),
		Whens: mapWhenPtrSlice(src.Whens),
		Else: mapExpr(src.Else),
	}
	return r
}

func mapWhenPtrSlice(src []*parser.When) []*When {
	if src == nil {
		return nil
	}
	result := make([]*When, 0, len(src))
	for _, v := range src {
		result = append(result, mapWhenPtr(v))
	}
	return result
}

func mapWhenPtr(src *parser.When) *When {
	if src == nil {
		return nil
	}
	r := &When{
		Cond: mapExpr(src.Cond),
		Val: mapExpr(src.Val),
	}
	return r
}

type When struct {
	Cond Expr
	Val Expr
}

type ValuesFuncExpr struct {
	Name *ColName
}

func mapValuesFuncExprPtr(src *parser.ValuesFuncExpr) *ValuesFuncExpr {
	if src == nil {
		return nil
	}
	r := &ValuesFuncExpr{
		Name: mapColNamePtr(src.Name),
	}
	return r
}

type CastExpr struct {
	Expr Expr
	Type *ConvertType
}

func mapCastExprPtr(src *parser.CastExpr) *CastExpr {
	if src == nil {
		return nil
	}
	r := &CastExpr{
		Expr: mapExpr(src.Expr),
		Type: mapConvertTypePtr(src.Type),
	}
	return r
}

func mapConvertTypePtr(src *parser.ConvertType) *ConvertType {
	if src == nil {
		return nil
	}
	r := &ConvertType{
		Type: src.Type,
		Length: mapSQLValPtr(src.Length),
		Scale: mapSQLValPtr(src.Scale),
		Operator: src.Operator,
		Charset: src.Charset,
	}
	return r
}

type ConvertType struct {
	Type string
	Length *SQLVal
	Scale *SQLVal
	Operator string
	Charset string
}

type ConvertExpr struct {
	Expr Expr
	Type *ConvertType
}

func mapConvertExprPtr(src *parser.ConvertExpr) *ConvertExpr {
	if src == nil {
		return nil
	}
	r := &ConvertExpr{
		Expr: mapExpr(src.Expr),
		Type: mapConvertTypePtr(src.Type),
	}
	return r
}

type SubstrExpr struct {
	Name Expr
	From Expr
	To Expr
}

func mapSubstrExprPtr(src *parser.SubstrExpr) *SubstrExpr {
	if src == nil {
		return nil
	}
	r := &SubstrExpr{
		Name: mapExpr(src.Name),
		From: mapExpr(src.From),
		To: mapExpr(src.To),
	}
	return r
}

type ConvertUsingExpr struct {
	Expr Expr
	Type string
}

func mapConvertUsingExprPtr(src *parser.ConvertUsingExpr) *ConvertUsingExpr {
	if src == nil {
		return nil
	}
	r := &ConvertUsingExpr{
		Expr: mapExpr(src.Expr),
		Type: src.Type,
	}
	return r
}

type MatchExpr struct {
	Columns []SelectExpr
	Expr Expr
	Option string
}

func mapMatchExprPtr(src *parser.MatchExpr) *MatchExpr {
	if src == nil {
		return nil
	}
	r := &MatchExpr{
		Columns: mapSelectExprSlice(src.Columns),
		Expr: mapExpr(src.Expr),
		Option: src.Option,
	}
	return r
}

type GroupConcatExpr struct {
	Distinct string
	Exprs []SelectExpr
	OrderBy []*Order
	Separator string
}

func mapGroupConcatExprPtr(src *parser.GroupConcatExpr) *GroupConcatExpr {
	if src == nil {
		return nil
	}
	r := &GroupConcatExpr{
		Distinct: src.Distinct,
		Exprs: mapSelectExprSlice(src.Exprs),
		OrderBy: mapOrderPtrSlice(src.OrderBy),
		Separator: src.Separator,
	}
	return r
}

type Default struct {
	ColName string
}

func mapDefaultPtr(src *parser.Default) *Default {
	if src == nil {
		return nil
	}
	r := &Default{
		ColName: src.ColName,
	}
	return r
}

type ArrayConstructor struct {
	Elements []ArrayElement
}

func mapArrayConstructorPtr(src *parser.ArrayConstructor) *ArrayConstructor {
	if src == nil {
		return nil
	}
	r := &ArrayConstructor{
		Elements: mapArrayElementSlice(src.Elements),
	}
	return r
}

func mapArrayElementSlice(src []parser.ArrayElement) []ArrayElement {
	if src == nil {
		return nil
	}
	result := make([]ArrayElement, 0, len(src))
	for _, v := range src {
		result = append(result, mapArrayElement(v))
	}
	return result
}

type ArrayElement interface {
}

type FuncCallExpr struct {
	Name ColIdent
	Exprs []Expr
}

func mapFuncCallExprPtr(src *parser.FuncCallExpr) *FuncCallExpr {
	if src == nil {
		return nil
	}
	r := &FuncCallExpr{
		Name: *mapColIdentPtr(&src.Name),
		Exprs: mapExprSlice(src.Exprs),
	}
	return r
}

type NextSeqValExpr struct {
	SequenceName TableIdent
}

func mapNextSeqValExprPtr(src *parser.NextSeqValExpr) *NextSeqValExpr {
	if src == nil {
		return nil
	}
	r := &NextSeqValExpr{
		SequenceName: *mapTableIdentPtr(&src.SequenceName),
	}
	return r
}

type AliasedTableExpr struct {
	Expr SimpleTableExpr
	Partitions []ColIdent
	As TableIdent
	TableHints []string
	IndexHints *IndexHints
}

func mapAliasedTableExprPtr(src *parser.AliasedTableExpr) *AliasedTableExpr {
	if src == nil {
		return nil
	}
	r := &AliasedTableExpr{
		Expr: mapSimpleTableExpr(src.Expr),
		Partitions: mapColIdentSlice(src.Partitions),
		As: *mapTableIdentPtr(&src.As),
		TableHints: src.TableHints,
		IndexHints: mapIndexHintsPtr(src.IndexHints),
	}
	return r
}

type SimpleTableExpr interface {
}

func mapIndexHintsPtr(src *parser.IndexHints) *IndexHints {
	if src == nil {
		return nil
	}
	r := &IndexHints{
		Type: src.Type,
		Indexes: mapColIdentSlice(src.Indexes),
	}
	return r
}

type IndexHints struct {
	Type string
	Indexes []ColIdent
}

type ParenTableExpr struct {
	Exprs []TableExpr
}

func mapParenTableExprPtr(src *parser.ParenTableExpr) *ParenTableExpr {
	if src == nil {
		return nil
	}
	r := &ParenTableExpr{
		Exprs: mapTableExprSlice(src.Exprs),
	}
	return r
}

type JoinTableExpr struct {
	LeftExpr TableExpr
	Join string
	RightExpr TableExpr
	Condition JoinCondition
}

func mapJoinTableExprPtr(src *parser.JoinTableExpr) *JoinTableExpr {
	if src == nil {
		return nil
	}
	r := &JoinTableExpr{
		LeftExpr: mapTableExpr(src.LeftExpr),
		Join: src.Join,
		RightExpr: mapTableExpr(src.RightExpr),
		Condition: *mapJoinConditionPtr(&src.Condition),
	}
	return r
}

type JoinCondition struct {
	On Expr
	Using []ColIdent
}

func mapJoinConditionPtr(src *parser.JoinCondition) *JoinCondition {
	if src == nil {
		return nil
	}
	r := &JoinCondition{
		On: mapExpr(src.On),
		Using: mapColIdentSlice(src.Using),
	}
	return r
}

type ParenSelect struct {
	Select SelectStatement
}

func mapParenSelectPtr(src *parser.ParenSelect) *ParenSelect {
	if src == nil {
		return nil
	}
	r := &ParenSelect{
		Select: mapSelectStatement(src.Select),
	}
	return r
}

type StarExpr struct {
	TableName TableName
}

func mapStarExprPtr(src *parser.StarExpr) *StarExpr {
	if src == nil {
		return nil
	}
	r := &StarExpr{
		TableName: *mapTableNamePtr(&src.TableName),
	}
	return r
}

type AliasedExpr struct {
	Expr Expr
	As ColIdent
}

func mapAliasedExprPtr(src *parser.AliasedExpr) *AliasedExpr {
	if src == nil {
		return nil
	}
	r := &AliasedExpr{
		Expr: mapExpr(src.Expr),
		As: *mapColIdentPtr(&src.As),
	}
	return r
}

func mapExprSliceSlice(src []parser.ValTuple) []ValTuple {
	if src == nil {
		return nil
	}
	result := make([]ValTuple, 0, len(src))
	for _, v := range src {
		result = append(result, *mapExprSlicePtr(&v))
	}
	return result
}