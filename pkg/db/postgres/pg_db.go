package postgres

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/alibaba/ioc-golang/autowire"
)

// +ioc:autowire=true
// +ioc:autowire:type=normal
// +ioc:autowire:type=singleton
// +ioc:autowire:paramType=Param
// +ioc:autowire:constructFunc=New
type PgDB struct {
	db *gorm.DB
}

func fromDB(db *gorm.DB) PgDBIOCInterface {
	return autowire.GetProxyFunction()(&PgDB{
		db: db,
	}).(PgDBIOCInterface)
}

// Session create new db session
func (db *PgDB) Session(config *gorm.Session) PgDBIOCInterface {
	return fromDB(db.db.Session(config))
}

// WithContext change current instance db's context to ctx
func (db *PgDB) WithContext(ctx context.Context) PgDBIOCInterface {
	return fromDB(db.db.WithContext(ctx))
}

// Debug start debug mode
func (db *PgDB) Debug() PgDBIOCInterface {
	return fromDB(db.db.Debug())
}

// Debug start debug mode
func (db *PgDB) GetDB() *gorm.DB {
	return db.db
}

// Set store value with key into current db instance's context
func (db *PgDB) Set(key string, value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Set(key, value))
}

// Get get value with key from current db instance's context
func (db *PgDB) Get(key string) (interface{}, bool) {
	return db.db.Get(key)
}

// InstanceSet store value with key into current db instance's context
func (db *PgDB) InstanceSet(key string, value interface{}) PgDBIOCInterface {
	return fromDB(db.db.InstanceSet(key, value))
}

// InstanceGet get value with key from current db instance's context
func (db *PgDB) InstanceGet(key string) (interface{}, bool) {
	return db.db.InstanceGet(key)
}

// AddError add error to db
func (db *PgDB) AddError(err error) error {
	return db.db.AddError(err)
}

// DB returns `*sql.DB`
func (db *PgDB) DB() (*sql.DB, error) {
	return db.db.DB()
}

// SetupJoinTable setup join table schema
func (db *PgDB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error {
	return db.db.SetupJoinTable(model, field, joinTable)
}

// Use use plugin
func (db *PgDB) Use(plugin gorm.Plugin) error {
	return db.db.Use(plugin)
}

func (db *PgDB) ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string {
	return db.db.ToSQL(queryFn)
}

// Model specify the model you would like to run db operations
//
//	// update all users's name to `hello`
//	db.Model(&User{}).Update("name", "hello")
//	// if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//	db.Model(&user).Update("name", "hello")
func (db *PgDB) Model(value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Model(value))
}

// Clauses Add clauses
func (db *PgDB) Clauses(conds ...clause.Expression) PgDBIOCInterface {
	return fromDB(db.db.Clauses(conds...))
}

// Table specify the table you would like to run db operations
func (db *PgDB) Table(name string, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Table(name, args...))
}

// Distinct specify distinct fields that you want querying
func (db *PgDB) Distinct(args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Distinct(args...))
}

// Select specify fields that you want when querying, creating, updating
func (db *PgDB) Select(query interface{}, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Select(query, args...))
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (db *PgDB) Omit(columns ...string) PgDBIOCInterface {
	return fromDB(db.db.Omit(columns...))
}

// Where add conditions
func (db *PgDB) Where(query interface{}, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Where(query, args...))
}

// Not add NOT conditions
func (db *PgDB) Not(query interface{}, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Not(query, args...))
}

// Or add OR conditions
func (db *PgDB) Or(query interface{}, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Or(query, args...))
}

// Joins specify Joins conditions
//
//	db.Joins("Account").Find(&user)
//	db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
//	db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
func (db *PgDB) Joins(query string, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Joins(query, args...))
}

// Group specify the group method on the find
func (db *PgDB) Group(name string) PgDBIOCInterface {
	return fromDB(db.db.Group(name))
}

// Having specify HAVING conditions for GROUP BY
func (db *PgDB) Having(query interface{}, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Having(query, args...))
}

// Order specify order when retrieve records from database
//
//	db.Order("name DESC")
//	db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (db *PgDB) Order(value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Order(value))
}

// Limit specify the number of records to be retrieved
func (db *PgDB) Limit(limit int) PgDBIOCInterface {
	return fromDB(db.db.Limit(limit))
}

// Offset specify the number of records to skip before starting to return the records
func (db *PgDB) Offset(offset int) PgDBIOCInterface {
	return fromDB(db.db.Offset(offset))
}

// Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically
//
//	func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//	    return db.Where("amount > ?", 1000)
//	}
//
//	func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//	    return func (db *gorm.DB) *gorm.DB {
//	        return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//	    }
//	}
//
//	db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
func (db *PgDB) Scopes(funcs ...func(db *gorm.DB) *gorm.DB) PgDBIOCInterface {
	return fromDB(db.db.Scopes(funcs...))
}

// Preload preload associations with given conditions
//
//	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *PgDB) Preload(query string, args ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Preload(query, args...))
}

func (db *PgDB) Attrs(attrs ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Attrs(attrs...))
}

func (db *PgDB) Assign(attrs ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Assign(attrs...))
}

func (db *PgDB) Unscoped() PgDBIOCInterface {
	return fromDB(db.db.Unscoped())
}

func (db *PgDB) Raw(sql string, values ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Raw(sql, values...))
}

func (db *PgDB) Error() error {
	return db.db.Error
}

// Create insert the value into database
func (db *PgDB) Create(value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Create(value))
}

// CreateInBatches insert the value in batches into database
func (db *PgDB) CreateInBatches(value interface{}, batchSize int) PgDBIOCInterface {
	return fromDB(db.db.CreateInBatches(value, batchSize))
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *PgDB) Save(value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Save(value))
}

// First find first record that match given conditions, order by primary key
func (db *PgDB) First(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.First(dest, conds...))
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (db *PgDB) Take(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Take(dest, conds...))
}

// Last find last record that match given conditions, order by primary key
func (db *PgDB) Last(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Last(dest, conds...))
}

// Find find records that match given conditions
func (db *PgDB) Find(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Find(dest, conds...))
}

// FindInBatches find records in batches
func (db *PgDB) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) PgDBIOCInterface {
	return fromDB(db.db.FindInBatches(dest, batchSize, fc))
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (db *PgDB) FirstOrInit(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.FirstOrInit(dest, conds...))
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (db *PgDB) FirstOrCreate(dest interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.FirstOrCreate(dest, conds...))
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *PgDB) Update(column string, value interface{}) PgDBIOCInterface {
	return fromDB(db.db.Update(column, value))
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *PgDB) Updates(values interface{}) PgDBIOCInterface {
	return fromDB(db.db.Updates(values))
}

func (db *PgDB) UpdateColumn(column string, value interface{}) PgDBIOCInterface {
	return fromDB(db.db.UpdateColumn(column, value))
}

func (db *PgDB) UpdateColumns(values interface{}) PgDBIOCInterface {
	return fromDB(db.db.UpdateColumns(values))
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *PgDB) Delete(value interface{}, conds ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Delete(value, conds...))
}

func (db *PgDB) Count(count *int64) PgDBIOCInterface {
	return fromDB(db.db.Count(count))
}

func (db *PgDB) Row() *sql.Row {
	return db.db.Row()
}

func (db *PgDB) Rows() (*sql.Rows, error) {
	return db.db.Rows()
}

// Scan scan value to a struct
func (db *PgDB) Scan(dest interface{}) PgDBIOCInterface {
	return fromDB(db.db.Scan(dest))
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (db *PgDB) Pluck(column string, dest interface{}) PgDBIOCInterface {
	return fromDB(db.db.Pluck(column, dest))
}

func (db *PgDB) ScanRows(rows *sql.Rows, dest interface{}) error {
	return db.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (db *PgDB) Connection(fc func(db *gorm.DB) error) (err error) {
	return db.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (db *PgDB) Transaction(fc func(db *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return db.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (db *PgDB) Begin(opts ...*sql.TxOptions) PgDBIOCInterface {
	return fromDB(db.db.Begin(opts...))
}

// Commit commit a transaction
func (db *PgDB) Commit() PgDBIOCInterface {
	return fromDB(db.db.Commit())
}

// Rollback rollback a transaction
func (db *PgDB) Rollback() {
	db.db.Rollback()
}

func (db *PgDB) SavePoint(name string) PgDBIOCInterface {
	return fromDB(db.db.SavePoint(name))
}

func (db *PgDB) RollbackTo(name string) PgDBIOCInterface {
	return fromDB(db.db.RollbackTo(name))
}

// Exec execute raw sql
func (db *PgDB) Exec(sql string, values ...interface{}) PgDBIOCInterface {
	return fromDB(db.db.Exec(sql, values...))
}

func (db *PgDB) Migrator() gorm.Migrator {
	return db.db.Migrator()
}

func (db *PgDB) AutoMigrate(dst ...interface{}) error {
	return db.db.AutoMigrate(dst...)
}

func (db *PgDB) Association(column string) *gorm.Association {
	return db.db.Association(column)
}
