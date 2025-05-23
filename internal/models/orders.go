// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Order is an object representing the database table.
type Order struct {
	ID          string        `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID      string        `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Status      string        `boil:"status" json:"status" toml:"status" yaml:"status"`
	TotalAmount types.Decimal `boil:"total_amount" json:"total_amount" toml:"total_amount" yaml:"total_amount"`
	CreatedAt   null.Time     `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time     `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *orderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderColumns = struct {
	ID          string
	UserID      string
	Status      string
	TotalAmount string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	Status:      "status",
	TotalAmount: "total_amount",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var OrderTableColumns = struct {
	ID          string
	UserID      string
	Status      string
	TotalAmount string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "orders.id",
	UserID:      "orders.user_id",
	Status:      "orders.status",
	TotalAmount: "orders.total_amount",
	CreatedAt:   "orders.created_at",
	UpdatedAt:   "orders.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod    { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod   { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod   { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) SIMILAR(x string) qm.QueryMod { return qm.Where(w.field+" SIMILAR TO ?", x) }
func (w whereHelperstring) NSIMILAR(x string) qm.QueryMod {
	return qm.Where(w.field+" NOT SIMILAR TO ?", x)
}
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_Decimal struct{ field string }

func (w whereHelpertypes_Decimal) EQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Decimal) NEQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Decimal) LT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Decimal) LTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Decimal) GT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Decimal) GTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var OrderWhere = struct {
	ID          whereHelperstring
	UserID      whereHelperstring
	Status      whereHelperstring
	TotalAmount whereHelpertypes_Decimal
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	ID:          whereHelperstring{field: "\"orders\".\"id\""},
	UserID:      whereHelperstring{field: "\"orders\".\"user_id\""},
	Status:      whereHelperstring{field: "\"orders\".\"status\""},
	TotalAmount: whereHelpertypes_Decimal{field: "\"orders\".\"total_amount\""},
	CreatedAt:   whereHelpernull_Time{field: "\"orders\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"orders\".\"updated_at\""},
}

// OrderRels is where relationship names are stored.
var OrderRels = struct {
}{}

// orderR is where relationships are stored.
type orderR struct {
}

// NewStruct creates a new relationship struct
func (*orderR) NewStruct() *orderR {
	return &orderR{}
}

// orderL is where Load methods for each relationship are stored.
type orderL struct{}

var (
	orderAllColumns            = []string{"id", "user_id", "status", "total_amount", "created_at", "updated_at"}
	orderColumnsWithoutDefault = []string{"user_id", "total_amount"}
	orderColumnsWithDefault    = []string{"id", "status", "created_at", "updated_at"}
	orderPrimaryKeyColumns     = []string{"id"}
	orderGeneratedColumns      = []string{}
)

type (
	// OrderSlice is an alias for a slice of pointers to Order.
	// This should almost always be used instead of []Order.
	OrderSlice []*Order
	// OrderHook is the signature for custom Order hook methods
	OrderHook func(context.Context, boil.ContextExecutor, *Order) error

	orderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderType                 = reflect.TypeOf(&Order{})
	orderMapping              = queries.MakeStructMapping(orderType)
	orderPrimaryKeyMapping, _ = queries.BindMapping(orderType, orderMapping, orderPrimaryKeyColumns)
	orderInsertCacheMut       sync.RWMutex
	orderInsertCache          = make(map[string]insertCache)
	orderUpdateCacheMut       sync.RWMutex
	orderUpdateCache          = make(map[string]updateCache)
	orderUpsertCacheMut       sync.RWMutex
	orderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderAfterSelectMu sync.Mutex
var orderAfterSelectHooks []OrderHook

var orderBeforeInsertMu sync.Mutex
var orderBeforeInsertHooks []OrderHook
var orderAfterInsertMu sync.Mutex
var orderAfterInsertHooks []OrderHook

var orderBeforeUpdateMu sync.Mutex
var orderBeforeUpdateHooks []OrderHook
var orderAfterUpdateMu sync.Mutex
var orderAfterUpdateHooks []OrderHook

var orderBeforeDeleteMu sync.Mutex
var orderBeforeDeleteHooks []OrderHook
var orderAfterDeleteMu sync.Mutex
var orderAfterDeleteHooks []OrderHook

var orderBeforeUpsertMu sync.Mutex
var orderBeforeUpsertHooks []OrderHook
var orderAfterUpsertMu sync.Mutex
var orderAfterUpsertHooks []OrderHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Order) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Order) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Order) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Order) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Order) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Order) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Order) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Order) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Order) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderHook registers your hook function for all future operations.
func AddOrderHook(hookPoint boil.HookPoint, orderHook OrderHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderAfterSelectMu.Lock()
		orderAfterSelectHooks = append(orderAfterSelectHooks, orderHook)
		orderAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderBeforeInsertMu.Lock()
		orderBeforeInsertHooks = append(orderBeforeInsertHooks, orderHook)
		orderBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderAfterInsertMu.Lock()
		orderAfterInsertHooks = append(orderAfterInsertHooks, orderHook)
		orderAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		orderBeforeUpdateMu.Lock()
		orderBeforeUpdateHooks = append(orderBeforeUpdateHooks, orderHook)
		orderBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		orderAfterUpdateMu.Lock()
		orderAfterUpdateHooks = append(orderAfterUpdateHooks, orderHook)
		orderAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		orderBeforeDeleteMu.Lock()
		orderBeforeDeleteHooks = append(orderBeforeDeleteHooks, orderHook)
		orderBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		orderAfterDeleteMu.Lock()
		orderAfterDeleteHooks = append(orderAfterDeleteHooks, orderHook)
		orderAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		orderBeforeUpsertMu.Lock()
		orderBeforeUpsertHooks = append(orderBeforeUpsertHooks, orderHook)
		orderBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderAfterUpsertMu.Lock()
		orderAfterUpsertHooks = append(orderAfterUpsertHooks, orderHook)
		orderAfterUpsertMu.Unlock()
	}
}

// One returns a single order record from the query.
func (q orderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Order, error) {
	o := &Order{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for orders")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Order records from the query.
func (q orderQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderSlice, error) {
	var o []*Order

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Order slice")
	}

	if len(orderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Order records in the query.
func (q orderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count orders rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if orders exists")
	}

	return count > 0, nil
}

// Orders retrieves all the records using an executor.
func Orders(mods ...qm.QueryMod) orderQuery {
	mods = append(mods, qm.From("\"orders\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"orders\".*"})
	}

	return orderQuery{q}
}

// FindOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrder(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Order, error) {
	orderObj := &Order{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"orders\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from orders")
	}

	if err = orderObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderObj, err
	}

	return orderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Order) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orders provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderInsertCacheMut.RLock()
	cache, cached := orderInsertCache[key]
	orderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderType, orderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"orders\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"orders\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into orders")
	}

	if !cached {
		orderInsertCacheMut.Lock()
		orderInsertCache[key] = cache
		orderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Order.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Order) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderUpdateCacheMut.RLock()
	cache, cached := orderUpdateCache[key]
	orderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update orders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"orders\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, append(wl, orderPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update orders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for orders")
	}

	if !cached {
		orderUpdateCacheMut.Lock()
		orderUpdateCache[key] = cache
		orderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for orders")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"orders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all order")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Order) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no orders provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderUpsertCacheMut.RLock()
	cache, cached := orderUpsertCache[key]
	orderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert orders, could not build update column list")
		}

		ret := strmangle.SetComplement(orderAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(orderPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert orders, could not build conflict column list")
			}

			conflict = make([]string, len(orderPrimaryKeyColumns))
			copy(conflict, orderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"orders\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderType, orderMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert orders")
	}

	if !cached {
		orderUpsertCacheMut.Lock()
		orderUpsertCache[key] = cache
		orderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Order record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Order) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Order provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderPrimaryKeyMapping)
	sql := "DELETE FROM \"orders\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for orders")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	if len(orderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Order) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"orders\".* FROM \"orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderSlice")
	}

	*o = slice

	return nil
}

// OrderExists checks if the Order row exists.
func OrderExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"orders\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if orders exists")
	}

	return exists, nil
}

// Exists checks if the Order row exists.
func (o *Order) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderExists(ctx, exec, o.ID)
}
