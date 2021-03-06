// Code generated by https://github.com/src-d/go-kallax. DO NOT EDIT.
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package main

import (
	"database/sql"
	"fmt"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

type modelSaveFunc func(*kallax.Store) error

// NewTestModel returns a new instance of TestModel.
func NewTestModel() (record *TestModel) {
	return new(TestModel)
}

// GetID returns the primary key of the model.
func (r *TestModel) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *TestModel) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "data":
		return &r.Data, nil
	case "data2":
		return &r.Data2, nil
	case "counter":
		return &r.Counter, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in TestModel: %s", col)
	}
}

// Value returns the value of the given column.
func (r *TestModel) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "data":
		return r.Data, nil
	case "data2":
		return r.Data2, nil
	case "counter":
		return r.Counter, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in TestModel: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *TestModel) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model TestModel has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *TestModel) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model TestModel has no relationships")
}

// TestModelStore is the entity to access the records of the type TestModel
// in the database.
type TestModelStore struct {
	*kallax.Store
}

// NewTestModelStore creates a new instance of TestModelStore
// using a SQL database.
func NewTestModelStore(db *sql.DB) *TestModelStore {
	return &TestModelStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *TestModelStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *TestModelStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *TestModelStore) Debug() *TestModelStore {
	return &TestModelStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *TestModelStore) DebugWith(logger kallax.LoggerFunc) *TestModelStore {
	return &TestModelStore{s.Store.DebugWith(logger)}
}

// DisableCacher turns off prepared statements, which can be useful in some scenarios.
func (s *TestModelStore) DisableCacher() *TestModelStore {
	return &TestModelStore{s.Store.DisableCacher()}
}

// Insert inserts a TestModel in the database. A non-persisted object is
// required for this operation.
func (s *TestModelStore) Insert(record *TestModel) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Insert(Schema.TestModel.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *TestModelStore) Update(record *TestModel, cols ...kallax.SchemaField) (updated int64, err error) {
	record.SetSaving(true)
	defer record.SetSaving(false)

	return s.Store.Update(Schema.TestModel.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *TestModelStore) Save(record *TestModel) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *TestModelStore) Delete(record *TestModel) error {
	return s.Store.Delete(Schema.TestModel.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *TestModelStore) Find(q *TestModelQuery) (*TestModelResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewTestModelResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *TestModelStore) MustFind(q *TestModelQuery) *TestModelResultSet {
	return NewTestModelResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *TestModelStore) Count(q *TestModelQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *TestModelStore) MustCount(q *TestModelQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *TestModelStore) FindOne(q *TestModelQuery) (*TestModel, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *TestModelStore) FindAll(q *TestModelQuery) ([]*TestModel, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *TestModelStore) MustFindOne(q *TestModelQuery) *TestModel {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the TestModel with the data in the database and
// makes it writable.
func (s *TestModelStore) Reload(record *TestModel) error {
	return s.Store.Reload(Schema.TestModel.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *TestModelStore) Transaction(callback func(*TestModelStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&TestModelStore{store})
	})
}

// TestModelQuery is the object used to create queries for the TestModel
// entity.
type TestModelQuery struct {
	*kallax.BaseQuery
}

// NewTestModelQuery returns a new instance of TestModelQuery.
func NewTestModelQuery() *TestModelQuery {
	return &TestModelQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.TestModel.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *TestModelQuery) Select(columns ...kallax.SchemaField) *TestModelQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *TestModelQuery) SelectNot(columns ...kallax.SchemaField) *TestModelQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *TestModelQuery) Copy() *TestModelQuery {
	return &TestModelQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *TestModelQuery) Order(cols ...kallax.ColumnOrder) *TestModelQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *TestModelQuery) BatchSize(size uint64) *TestModelQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *TestModelQuery) Limit(n uint64) *TestModelQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *TestModelQuery) Offset(n uint64) *TestModelQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *TestModelQuery) Where(cond kallax.Condition) *TestModelQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *TestModelQuery) FindByID(v ...int64) *TestModelQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.TestModel.ID, values...))
}

// FindByData adds a new filter to the query that will require that
// the Data property is equal to the passed value.
func (q *TestModelQuery) FindByData(v byte) *TestModelQuery {
	return q.Where(kallax.Eq(Schema.TestModel.Data, v))
}

// FindByData2 adds a new filter to the query that will require that
// the Data2 property is equal to the passed value.
func (q *TestModelQuery) FindByData2(v byte) *TestModelQuery {
	return q.Where(kallax.Eq(Schema.TestModel.Data2, v))
}

// FindByCounter adds a new filter to the query that will require that
// the Counter property is equal to the passed value.
func (q *TestModelQuery) FindByCounter(cond kallax.ScalarCond, v int64) *TestModelQuery {
	return q.Where(cond(Schema.TestModel.Counter, v))
}

// TestModelResultSet is the set of results returned by a query to the
// database.
type TestModelResultSet struct {
	ResultSet kallax.ResultSet
	last      *TestModel
	lastErr   error
}

// NewTestModelResultSet creates a new result set for rows of the type
// TestModel.
func NewTestModelResultSet(rs kallax.ResultSet) *TestModelResultSet {
	return &TestModelResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *TestModelResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.TestModel.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*TestModel)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *TestModel")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *TestModelResultSet) Get() (*TestModel, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *TestModelResultSet) ForEach(fn func(*TestModel) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *TestModelResultSet) All() ([]*TestModel, error) {
	var result []*TestModel
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *TestModelResultSet) One() (*TestModel, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *TestModelResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *TestModelResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	TestModel *schemaTestModel
}

type schemaTestModel struct {
	*kallax.BaseSchema
	ID      kallax.SchemaField
	Data    kallax.SchemaField
	Data2   kallax.SchemaField
	Counter kallax.SchemaField
}

var Schema = &schema{
	TestModel: &schemaTestModel{
		BaseSchema: kallax.NewBaseSchema(
			"test_models",
			"__testmodel",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(TestModel)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("data"),
			kallax.NewSchemaField("data2"),
			kallax.NewSchemaField("counter"),
		),
		ID:      kallax.NewSchemaField("id"),
		Data:    kallax.NewSchemaField("data"),
		Data2:   kallax.NewSchemaField("data2"),
		Counter: kallax.NewSchemaField("counter"),
	},
}
