// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/kolakdd/auth_template/models"
)

func newInvalidAccessToken(db *gorm.DB, opts ...gen.DOOption) invalidAccessToken {
	_invalidAccessToken := invalidAccessToken{}

	_invalidAccessToken.invalidAccessTokenDo.UseDB(db, opts...)
	_invalidAccessToken.invalidAccessTokenDo.UseModel(&models.InvalidAccessToken{})

	tableName := _invalidAccessToken.invalidAccessTokenDo.TableName()
	_invalidAccessToken.ALL = field.NewAsterisk(tableName)
	_invalidAccessToken.GUID = field.NewField(tableName, "guid")
	_invalidAccessToken.UserGUID = field.NewField(tableName, "user_guid")
	_invalidAccessToken.CreatedAt = field.NewTime(tableName, "created_at")

	_invalidAccessToken.fillFieldMap()

	return _invalidAccessToken
}

type invalidAccessToken struct {
	invalidAccessTokenDo

	ALL       field.Asterisk
	GUID      field.Field
	UserGUID  field.Field
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (i invalidAccessToken) Table(newTableName string) *invalidAccessToken {
	i.invalidAccessTokenDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i invalidAccessToken) As(alias string) *invalidAccessToken {
	i.invalidAccessTokenDo.DO = *(i.invalidAccessTokenDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *invalidAccessToken) updateTableName(table string) *invalidAccessToken {
	i.ALL = field.NewAsterisk(table)
	i.GUID = field.NewField(table, "guid")
	i.UserGUID = field.NewField(table, "user_guid")
	i.CreatedAt = field.NewTime(table, "created_at")

	i.fillFieldMap()

	return i
}

func (i *invalidAccessToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *invalidAccessToken) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 3)
	i.fieldMap["guid"] = i.GUID
	i.fieldMap["user_guid"] = i.UserGUID
	i.fieldMap["created_at"] = i.CreatedAt
}

func (i invalidAccessToken) clone(db *gorm.DB) invalidAccessToken {
	i.invalidAccessTokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i invalidAccessToken) replaceDB(db *gorm.DB) invalidAccessToken {
	i.invalidAccessTokenDo.ReplaceDB(db)
	return i
}

type invalidAccessTokenDo struct{ gen.DO }

type IInvalidAccessTokenDo interface {
	gen.SubQuery
	Debug() IInvalidAccessTokenDo
	WithContext(ctx context.Context) IInvalidAccessTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInvalidAccessTokenDo
	WriteDB() IInvalidAccessTokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInvalidAccessTokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInvalidAccessTokenDo
	Not(conds ...gen.Condition) IInvalidAccessTokenDo
	Or(conds ...gen.Condition) IInvalidAccessTokenDo
	Select(conds ...field.Expr) IInvalidAccessTokenDo
	Where(conds ...gen.Condition) IInvalidAccessTokenDo
	Order(conds ...field.Expr) IInvalidAccessTokenDo
	Distinct(cols ...field.Expr) IInvalidAccessTokenDo
	Omit(cols ...field.Expr) IInvalidAccessTokenDo
	Join(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo
	Group(cols ...field.Expr) IInvalidAccessTokenDo
	Having(conds ...gen.Condition) IInvalidAccessTokenDo
	Limit(limit int) IInvalidAccessTokenDo
	Offset(offset int) IInvalidAccessTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInvalidAccessTokenDo
	Unscoped() IInvalidAccessTokenDo
	Create(values ...*models.InvalidAccessToken) error
	CreateInBatches(values []*models.InvalidAccessToken, batchSize int) error
	Save(values ...*models.InvalidAccessToken) error
	First() (*models.InvalidAccessToken, error)
	Take() (*models.InvalidAccessToken, error)
	Last() (*models.InvalidAccessToken, error)
	Find() ([]*models.InvalidAccessToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.InvalidAccessToken, err error)
	FindInBatches(result *[]*models.InvalidAccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.InvalidAccessToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInvalidAccessTokenDo
	Assign(attrs ...field.AssignExpr) IInvalidAccessTokenDo
	Joins(fields ...field.RelationField) IInvalidAccessTokenDo
	Preload(fields ...field.RelationField) IInvalidAccessTokenDo
	FirstOrInit() (*models.InvalidAccessToken, error)
	FirstOrCreate() (*models.InvalidAccessToken, error)
	FindByPage(offset int, limit int) (result []*models.InvalidAccessToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInvalidAccessTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i invalidAccessTokenDo) Debug() IInvalidAccessTokenDo {
	return i.withDO(i.DO.Debug())
}

func (i invalidAccessTokenDo) WithContext(ctx context.Context) IInvalidAccessTokenDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i invalidAccessTokenDo) ReadDB() IInvalidAccessTokenDo {
	return i.Clauses(dbresolver.Read)
}

func (i invalidAccessTokenDo) WriteDB() IInvalidAccessTokenDo {
	return i.Clauses(dbresolver.Write)
}

func (i invalidAccessTokenDo) Session(config *gorm.Session) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Session(config))
}

func (i invalidAccessTokenDo) Clauses(conds ...clause.Expression) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i invalidAccessTokenDo) Returning(value interface{}, columns ...string) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i invalidAccessTokenDo) Not(conds ...gen.Condition) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i invalidAccessTokenDo) Or(conds ...gen.Condition) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i invalidAccessTokenDo) Select(conds ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i invalidAccessTokenDo) Where(conds ...gen.Condition) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i invalidAccessTokenDo) Order(conds ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i invalidAccessTokenDo) Distinct(cols ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i invalidAccessTokenDo) Omit(cols ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i invalidAccessTokenDo) Join(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i invalidAccessTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i invalidAccessTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i invalidAccessTokenDo) Group(cols ...field.Expr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i invalidAccessTokenDo) Having(conds ...gen.Condition) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i invalidAccessTokenDo) Limit(limit int) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i invalidAccessTokenDo) Offset(offset int) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i invalidAccessTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i invalidAccessTokenDo) Unscoped() IInvalidAccessTokenDo {
	return i.withDO(i.DO.Unscoped())
}

func (i invalidAccessTokenDo) Create(values ...*models.InvalidAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i invalidAccessTokenDo) CreateInBatches(values []*models.InvalidAccessToken, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i invalidAccessTokenDo) Save(values ...*models.InvalidAccessToken) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i invalidAccessTokenDo) First() (*models.InvalidAccessToken, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.InvalidAccessToken), nil
	}
}

func (i invalidAccessTokenDo) Take() (*models.InvalidAccessToken, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.InvalidAccessToken), nil
	}
}

func (i invalidAccessTokenDo) Last() (*models.InvalidAccessToken, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.InvalidAccessToken), nil
	}
}

func (i invalidAccessTokenDo) Find() ([]*models.InvalidAccessToken, error) {
	result, err := i.DO.Find()
	return result.([]*models.InvalidAccessToken), err
}

func (i invalidAccessTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.InvalidAccessToken, err error) {
	buf := make([]*models.InvalidAccessToken, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i invalidAccessTokenDo) FindInBatches(result *[]*models.InvalidAccessToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i invalidAccessTokenDo) Attrs(attrs ...field.AssignExpr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i invalidAccessTokenDo) Assign(attrs ...field.AssignExpr) IInvalidAccessTokenDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i invalidAccessTokenDo) Joins(fields ...field.RelationField) IInvalidAccessTokenDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i invalidAccessTokenDo) Preload(fields ...field.RelationField) IInvalidAccessTokenDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i invalidAccessTokenDo) FirstOrInit() (*models.InvalidAccessToken, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.InvalidAccessToken), nil
	}
}

func (i invalidAccessTokenDo) FirstOrCreate() (*models.InvalidAccessToken, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.InvalidAccessToken), nil
	}
}

func (i invalidAccessTokenDo) FindByPage(offset int, limit int) (result []*models.InvalidAccessToken, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i invalidAccessTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i invalidAccessTokenDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i invalidAccessTokenDo) Delete(models ...*models.InvalidAccessToken) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *invalidAccessTokenDo) withDO(do gen.Dao) *invalidAccessTokenDo {
	i.DO = *do.(*gen.DO)
	return i
}
