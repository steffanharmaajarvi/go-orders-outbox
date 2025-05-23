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

// Outbox is an object representing the database table.
type Outbox struct {
	ID            string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	AggregateType string     `boil:"aggregate_type" json:"aggregate_type" toml:"aggregate_type" yaml:"aggregate_type"`
	AggregateID   string     `boil:"aggregate_id" json:"aggregate_id" toml:"aggregate_id" yaml:"aggregate_id"`
	Type          string     `boil:"type" json:"type" toml:"type" yaml:"type"`
	Payload       types.JSON `boil:"payload" json:"payload" toml:"payload" yaml:"payload"`
	OccurredAt    null.Time  `boil:"occurred_at" json:"occurred_at,omitempty" toml:"occurred_at" yaml:"occurred_at,omitempty"`
	SentAt        null.Time  `boil:"sent_at" json:"sent_at,omitempty" toml:"sent_at" yaml:"sent_at,omitempty"`

	R *outboxR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L outboxL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OutboxColumns = struct {
	ID            string
	AggregateType string
	AggregateID   string
	Type          string
	Payload       string
	OccurredAt    string
	SentAt        string
}{
	ID:            "id",
	AggregateType: "aggregate_type",
	AggregateID:   "aggregate_id",
	Type:          "type",
	Payload:       "payload",
	OccurredAt:    "occurred_at",
	SentAt:        "sent_at",
}

var OutboxTableColumns = struct {
	ID            string
	AggregateType string
	AggregateID   string
	Type          string
	Payload       string
	OccurredAt    string
	SentAt        string
}{
	ID:            "outbox.id",
	AggregateType: "outbox.aggregate_type",
	AggregateID:   "outbox.aggregate_id",
	Type:          "outbox.type",
	Payload:       "outbox.payload",
	OccurredAt:    "outbox.occurred_at",
	SentAt:        "outbox.sent_at",
}

// Generated where

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var OutboxWhere = struct {
	ID            whereHelperstring
	AggregateType whereHelperstring
	AggregateID   whereHelperstring
	Type          whereHelperstring
	Payload       whereHelpertypes_JSON
	OccurredAt    whereHelpernull_Time
	SentAt        whereHelpernull_Time
}{
	ID:            whereHelperstring{field: "\"outbox\".\"id\""},
	AggregateType: whereHelperstring{field: "\"outbox\".\"aggregate_type\""},
	AggregateID:   whereHelperstring{field: "\"outbox\".\"aggregate_id\""},
	Type:          whereHelperstring{field: "\"outbox\".\"type\""},
	Payload:       whereHelpertypes_JSON{field: "\"outbox\".\"payload\""},
	OccurredAt:    whereHelpernull_Time{field: "\"outbox\".\"occurred_at\""},
	SentAt:        whereHelpernull_Time{field: "\"outbox\".\"sent_at\""},
}

// OutboxRels is where relationship names are stored.
var OutboxRels = struct {
}{}

// outboxR is where relationships are stored.
type outboxR struct {
}

// NewStruct creates a new relationship struct
func (*outboxR) NewStruct() *outboxR {
	return &outboxR{}
}

// outboxL is where Load methods for each relationship are stored.
type outboxL struct{}

var (
	outboxAllColumns            = []string{"id", "aggregate_type", "aggregate_id", "type", "payload", "occurred_at", "sent_at"}
	outboxColumnsWithoutDefault = []string{"aggregate_type", "aggregate_id", "type", "payload"}
	outboxColumnsWithDefault    = []string{"id", "occurred_at", "sent_at"}
	outboxPrimaryKeyColumns     = []string{"id"}
	outboxGeneratedColumns      = []string{}
)

type (
	// OutboxSlice is an alias for a slice of pointers to Outbox.
	// This should almost always be used instead of []Outbox.
	OutboxSlice []*Outbox
	// OutboxHook is the signature for custom Outbox hook methods
	OutboxHook func(context.Context, boil.ContextExecutor, *Outbox) error

	outboxQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	outboxType                 = reflect.TypeOf(&Outbox{})
	outboxMapping              = queries.MakeStructMapping(outboxType)
	outboxPrimaryKeyMapping, _ = queries.BindMapping(outboxType, outboxMapping, outboxPrimaryKeyColumns)
	outboxInsertCacheMut       sync.RWMutex
	outboxInsertCache          = make(map[string]insertCache)
	outboxUpdateCacheMut       sync.RWMutex
	outboxUpdateCache          = make(map[string]updateCache)
	outboxUpsertCacheMut       sync.RWMutex
	outboxUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var outboxAfterSelectMu sync.Mutex
var outboxAfterSelectHooks []OutboxHook

var outboxBeforeInsertMu sync.Mutex
var outboxBeforeInsertHooks []OutboxHook
var outboxAfterInsertMu sync.Mutex
var outboxAfterInsertHooks []OutboxHook

var outboxBeforeUpdateMu sync.Mutex
var outboxBeforeUpdateHooks []OutboxHook
var outboxAfterUpdateMu sync.Mutex
var outboxAfterUpdateHooks []OutboxHook

var outboxBeforeDeleteMu sync.Mutex
var outboxBeforeDeleteHooks []OutboxHook
var outboxAfterDeleteMu sync.Mutex
var outboxAfterDeleteHooks []OutboxHook

var outboxBeforeUpsertMu sync.Mutex
var outboxBeforeUpsertHooks []OutboxHook
var outboxAfterUpsertMu sync.Mutex
var outboxAfterUpsertHooks []OutboxHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Outbox) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Outbox) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Outbox) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Outbox) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Outbox) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Outbox) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Outbox) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Outbox) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Outbox) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range outboxAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOutboxHook registers your hook function for all future operations.
func AddOutboxHook(hookPoint boil.HookPoint, outboxHook OutboxHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		outboxAfterSelectMu.Lock()
		outboxAfterSelectHooks = append(outboxAfterSelectHooks, outboxHook)
		outboxAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		outboxBeforeInsertMu.Lock()
		outboxBeforeInsertHooks = append(outboxBeforeInsertHooks, outboxHook)
		outboxBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		outboxAfterInsertMu.Lock()
		outboxAfterInsertHooks = append(outboxAfterInsertHooks, outboxHook)
		outboxAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		outboxBeforeUpdateMu.Lock()
		outboxBeforeUpdateHooks = append(outboxBeforeUpdateHooks, outboxHook)
		outboxBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		outboxAfterUpdateMu.Lock()
		outboxAfterUpdateHooks = append(outboxAfterUpdateHooks, outboxHook)
		outboxAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		outboxBeforeDeleteMu.Lock()
		outboxBeforeDeleteHooks = append(outboxBeforeDeleteHooks, outboxHook)
		outboxBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		outboxAfterDeleteMu.Lock()
		outboxAfterDeleteHooks = append(outboxAfterDeleteHooks, outboxHook)
		outboxAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		outboxBeforeUpsertMu.Lock()
		outboxBeforeUpsertHooks = append(outboxBeforeUpsertHooks, outboxHook)
		outboxBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		outboxAfterUpsertMu.Lock()
		outboxAfterUpsertHooks = append(outboxAfterUpsertHooks, outboxHook)
		outboxAfterUpsertMu.Unlock()
	}
}

// One returns a single outbox record from the query.
func (q outboxQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Outbox, error) {
	o := &Outbox{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for outbox")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Outbox records from the query.
func (q outboxQuery) All(ctx context.Context, exec boil.ContextExecutor) (OutboxSlice, error) {
	var o []*Outbox

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Outbox slice")
	}

	if len(outboxAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Outbox records in the query.
func (q outboxQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count outbox rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q outboxQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if outbox exists")
	}

	return count > 0, nil
}

// Outboxes retrieves all the records using an executor.
func Outboxes(mods ...qm.QueryMod) outboxQuery {
	mods = append(mods, qm.From("\"outbox\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"outbox\".*"})
	}

	return outboxQuery{q}
}

// FindOutbox retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOutbox(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Outbox, error) {
	outboxObj := &Outbox{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"outbox\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, outboxObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from outbox")
	}

	if err = outboxObj.doAfterSelectHooks(ctx, exec); err != nil {
		return outboxObj, err
	}

	return outboxObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Outbox) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no outbox provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(outboxColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	outboxInsertCacheMut.RLock()
	cache, cached := outboxInsertCache[key]
	outboxInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			outboxAllColumns,
			outboxColumnsWithDefault,
			outboxColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(outboxType, outboxMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(outboxType, outboxMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"outbox\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"outbox\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into outbox")
	}

	if !cached {
		outboxInsertCacheMut.Lock()
		outboxInsertCache[key] = cache
		outboxInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Outbox.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Outbox) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	outboxUpdateCacheMut.RLock()
	cache, cached := outboxUpdateCache[key]
	outboxUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			outboxAllColumns,
			outboxPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update outbox, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"outbox\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, outboxPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(outboxType, outboxMapping, append(wl, outboxPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update outbox row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for outbox")
	}

	if !cached {
		outboxUpdateCacheMut.Lock()
		outboxUpdateCache[key] = cache
		outboxUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q outboxQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for outbox")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for outbox")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OutboxSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outboxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"outbox\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, outboxPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in outbox slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all outbox")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Outbox) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no outbox provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(outboxColumnsWithDefault, o)

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

	outboxUpsertCacheMut.RLock()
	cache, cached := outboxUpsertCache[key]
	outboxUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			outboxAllColumns,
			outboxColumnsWithDefault,
			outboxColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			outboxAllColumns,
			outboxPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert outbox, could not build update column list")
		}

		ret := strmangle.SetComplement(outboxAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(outboxPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert outbox, could not build conflict column list")
			}

			conflict = make([]string, len(outboxPrimaryKeyColumns))
			copy(conflict, outboxPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"outbox\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(outboxType, outboxMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(outboxType, outboxMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert outbox")
	}

	if !cached {
		outboxUpsertCacheMut.Lock()
		outboxUpsertCache[key] = cache
		outboxUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Outbox record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Outbox) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Outbox provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), outboxPrimaryKeyMapping)
	sql := "DELETE FROM \"outbox\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from outbox")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for outbox")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q outboxQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no outboxQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from outbox")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for outbox")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OutboxSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(outboxBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outboxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"outbox\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, outboxPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from outbox slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for outbox")
	}

	if len(outboxAfterDeleteHooks) != 0 {
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
func (o *Outbox) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOutbox(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OutboxSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OutboxSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outboxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"outbox\".* FROM \"outbox\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, outboxPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OutboxSlice")
	}

	*o = slice

	return nil
}

// OutboxExists checks if the Outbox row exists.
func OutboxExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"outbox\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if outbox exists")
	}

	return exists, nil
}

// Exists checks if the Outbox row exists.
func (o *Outbox) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OutboxExists(ctx, exec, o.ID)
}
