// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Job is an object representing the database table.
type Job struct {
	ID          string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	Location    string    `boil:"location" json:"location" toml:"location" yaml:"location"`
	Status      int       `boil:"status" json:"status" toml:"status" yaml:"status"`
	PersonID    string    `boil:"person_id" json:"person_id" toml:"person_id" yaml:"person_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *jobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JobColumns = struct {
	ID          string
	Title       string
	Description string
	Location    string
	Status      string
	PersonID    string
	CreatedAt   string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	Location:    "location",
	Status:      "status",
	PersonID:    "person_id",
	CreatedAt:   "created_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var JobWhere = struct {
	ID          whereHelperstring
	Title       whereHelperstring
	Description whereHelperstring
	Location    whereHelperstring
	Status      whereHelperint
	PersonID    whereHelperstring
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperstring{field: "\"job\".\"id\""},
	Title:       whereHelperstring{field: "\"job\".\"title\""},
	Description: whereHelperstring{field: "\"job\".\"description\""},
	Location:    whereHelperstring{field: "\"job\".\"location\""},
	Status:      whereHelperint{field: "\"job\".\"status\""},
	PersonID:    whereHelperstring{field: "\"job\".\"person_id\""},
	CreatedAt:   whereHelpertime_Time{field: "\"job\".\"created_at\""},
}

// JobRels is where relationship names are stored.
var JobRels = struct {
	Person        string
	JobSeekerFavs string
	JobSkillMaps  string
}{
	Person:        "Person",
	JobSeekerFavs: "JobSeekerFavs",
	JobSkillMaps:  "JobSkillMaps",
}

// jobR is where relationships are stored.
type jobR struct {
	Person        *Person
	JobSeekerFavs JobSeekerFavSlice
	JobSkillMaps  JobSkillMapSlice
}

// NewStruct creates a new relationship struct
func (*jobR) NewStruct() *jobR {
	return &jobR{}
}

// jobL is where Load methods for each relationship are stored.
type jobL struct{}

var (
	jobAllColumns            = []string{"id", "title", "description", "location", "status", "person_id", "created_at"}
	jobColumnsWithoutDefault = []string{"id", "title", "description", "location", "person_id", "created_at"}
	jobColumnsWithDefault    = []string{"status"}
	jobPrimaryKeyColumns     = []string{"id"}
)

type (
	// JobSlice is an alias for a slice of pointers to Job.
	// This should generally be used opposed to []Job.
	JobSlice []*Job
	// JobHook is the signature for custom Job hook methods
	JobHook func(context.Context, boil.ContextExecutor, *Job) error

	jobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jobType                 = reflect.TypeOf(&Job{})
	jobMapping              = queries.MakeStructMapping(jobType)
	jobPrimaryKeyMapping, _ = queries.BindMapping(jobType, jobMapping, jobPrimaryKeyColumns)
	jobInsertCacheMut       sync.RWMutex
	jobInsertCache          = make(map[string]insertCache)
	jobUpdateCacheMut       sync.RWMutex
	jobUpdateCache          = make(map[string]updateCache)
	jobUpsertCacheMut       sync.RWMutex
	jobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var jobBeforeInsertHooks []JobHook
var jobBeforeUpdateHooks []JobHook
var jobBeforeDeleteHooks []JobHook
var jobBeforeUpsertHooks []JobHook

var jobAfterInsertHooks []JobHook
var jobAfterSelectHooks []JobHook
var jobAfterUpdateHooks []JobHook
var jobAfterDeleteHooks []JobHook
var jobAfterUpsertHooks []JobHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Job) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Job) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Job) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Job) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Job) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Job) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Job) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Job) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Job) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range jobAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJobHook registers your hook function for all future operations.
func AddJobHook(hookPoint boil.HookPoint, jobHook JobHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		jobBeforeInsertHooks = append(jobBeforeInsertHooks, jobHook)
	case boil.BeforeUpdateHook:
		jobBeforeUpdateHooks = append(jobBeforeUpdateHooks, jobHook)
	case boil.BeforeDeleteHook:
		jobBeforeDeleteHooks = append(jobBeforeDeleteHooks, jobHook)
	case boil.BeforeUpsertHook:
		jobBeforeUpsertHooks = append(jobBeforeUpsertHooks, jobHook)
	case boil.AfterInsertHook:
		jobAfterInsertHooks = append(jobAfterInsertHooks, jobHook)
	case boil.AfterSelectHook:
		jobAfterSelectHooks = append(jobAfterSelectHooks, jobHook)
	case boil.AfterUpdateHook:
		jobAfterUpdateHooks = append(jobAfterUpdateHooks, jobHook)
	case boil.AfterDeleteHook:
		jobAfterDeleteHooks = append(jobAfterDeleteHooks, jobHook)
	case boil.AfterUpsertHook:
		jobAfterUpsertHooks = append(jobAfterUpsertHooks, jobHook)
	}
}

// One returns a single job record from the query.
func (q jobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Job, error) {
	o := &Job{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for job")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Job records from the query.
func (q jobQuery) All(ctx context.Context, exec boil.ContextExecutor) (JobSlice, error) {
	var o []*Job

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Job slice")
	}

	if len(jobAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Job records in the query.
func (q jobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count job rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q jobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if job exists")
	}

	return count > 0, nil
}

// Person pointed to by the foreign key.
func (o *Job) Person(mods ...qm.QueryMod) personQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PersonID),
	}

	queryMods = append(queryMods, mods...)

	query := People(queryMods...)
	queries.SetFrom(query.Query, "\"person\"")

	return query
}

// JobSeekerFavs retrieves all the job_seeker_fav's JobSeekerFavs with an executor.
func (o *Job) JobSeekerFavs(mods ...qm.QueryMod) jobSeekerFavQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"job_seeker_fav\".\"job_id\"=?", o.ID),
	)

	query := JobSeekerFavs(queryMods...)
	queries.SetFrom(query.Query, "\"job_seeker_fav\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"job_seeker_fav\".*"})
	}

	return query
}

// JobSkillMaps retrieves all the job_skill_map's JobSkillMaps with an executor.
func (o *Job) JobSkillMaps(mods ...qm.QueryMod) jobSkillMapQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"job_skill_map\".\"job_id\"=?", o.ID),
	)

	query := JobSkillMaps(queryMods...)
	queries.SetFrom(query.Query, "\"job_skill_map\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"job_skill_map\".*"})
	}

	return query
}

// LoadPerson allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (jobL) LoadPerson(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJob interface{}, mods queries.Applicator) error {
	var slice []*Job
	var object *Job

	if singular {
		object = maybeJob.(*Job)
	} else {
		slice = *maybeJob.(*[]*Job)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jobR{}
		}
		args = append(args, object.PersonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jobR{}
			}

			for _, a := range args {
				if a == obj.PersonID {
					continue Outer
				}
			}

			args = append(args, obj.PersonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`person`), qm.WhereIn(`person.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Person")
	}

	var resultSlice []*Person
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Person")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for person")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for person")
	}

	if len(jobAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Person = foreign
		if foreign.R == nil {
			foreign.R = &personR{}
		}
		foreign.R.Jobs = append(foreign.R.Jobs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PersonID == foreign.ID {
				local.R.Person = foreign
				if foreign.R == nil {
					foreign.R = &personR{}
				}
				foreign.R.Jobs = append(foreign.R.Jobs, local)
				break
			}
		}
	}

	return nil
}

// LoadJobSeekerFavs allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (jobL) LoadJobSeekerFavs(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJob interface{}, mods queries.Applicator) error {
	var slice []*Job
	var object *Job

	if singular {
		object = maybeJob.(*Job)
	} else {
		slice = *maybeJob.(*[]*Job)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jobR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jobR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`job_seeker_fav`), qm.WhereIn(`job_seeker_fav.job_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load job_seeker_fav")
	}

	var resultSlice []*JobSeekerFav
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice job_seeker_fav")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on job_seeker_fav")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for job_seeker_fav")
	}

	if len(jobSeekerFavAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.JobSeekerFavs = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &jobSeekerFavR{}
			}
			foreign.R.Job = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.JobID {
				local.R.JobSeekerFavs = append(local.R.JobSeekerFavs, foreign)
				if foreign.R == nil {
					foreign.R = &jobSeekerFavR{}
				}
				foreign.R.Job = local
				break
			}
		}
	}

	return nil
}

// LoadJobSkillMaps allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (jobL) LoadJobSkillMaps(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJob interface{}, mods queries.Applicator) error {
	var slice []*Job
	var object *Job

	if singular {
		object = maybeJob.(*Job)
	} else {
		slice = *maybeJob.(*[]*Job)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jobR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jobR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`job_skill_map`), qm.WhereIn(`job_skill_map.job_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load job_skill_map")
	}

	var resultSlice []*JobSkillMap
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice job_skill_map")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on job_skill_map")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for job_skill_map")
	}

	if len(jobSkillMapAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.JobSkillMaps = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &jobSkillMapR{}
			}
			foreign.R.Job = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.JobID {
				local.R.JobSkillMaps = append(local.R.JobSkillMaps, foreign)
				if foreign.R == nil {
					foreign.R = &jobSkillMapR{}
				}
				foreign.R.Job = local
				break
			}
		}
	}

	return nil
}

// SetPerson of the job to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Jobs.
func (o *Job) SetPerson(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Person) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"job\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"person_id"}),
		strmangle.WhereClause("\"", "\"", 2, jobPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PersonID = related.ID
	if o.R == nil {
		o.R = &jobR{
			Person: related,
		}
	} else {
		o.R.Person = related
	}

	if related.R == nil {
		related.R = &personR{
			Jobs: JobSlice{o},
		}
	} else {
		related.R.Jobs = append(related.R.Jobs, o)
	}

	return nil
}

// AddJobSeekerFavs adds the given related objects to the existing relationships
// of the job, optionally inserting them as new records.
// Appends related to o.R.JobSeekerFavs.
// Sets related.R.Job appropriately.
func (o *Job) AddJobSeekerFavs(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*JobSeekerFav) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.JobID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"job_seeker_fav\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"job_id"}),
				strmangle.WhereClause("\"", "\"", 2, jobSeekerFavPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.JobID = o.ID
		}
	}

	if o.R == nil {
		o.R = &jobR{
			JobSeekerFavs: related,
		}
	} else {
		o.R.JobSeekerFavs = append(o.R.JobSeekerFavs, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jobSeekerFavR{
				Job: o,
			}
		} else {
			rel.R.Job = o
		}
	}
	return nil
}

// AddJobSkillMaps adds the given related objects to the existing relationships
// of the job, optionally inserting them as new records.
// Appends related to o.R.JobSkillMaps.
// Sets related.R.Job appropriately.
func (o *Job) AddJobSkillMaps(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*JobSkillMap) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.JobID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"job_skill_map\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"job_id"}),
				strmangle.WhereClause("\"", "\"", 2, jobSkillMapPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.JobID = o.ID
		}
	}

	if o.R == nil {
		o.R = &jobR{
			JobSkillMaps: related,
		}
	} else {
		o.R.JobSkillMaps = append(o.R.JobSkillMaps, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jobSkillMapR{
				Job: o,
			}
		} else {
			rel.R.Job = o
		}
	}
	return nil
}

// Jobs retrieves all the records using an executor.
func Jobs(mods ...qm.QueryMod) jobQuery {
	mods = append(mods, qm.From("\"job\""))
	return jobQuery{NewQuery(mods...)}
}

// FindJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJob(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Job, error) {
	jobObj := &Job{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"job\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, jobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from job")
	}

	return jobObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Job) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	jobInsertCacheMut.RLock()
	cache, cached := jobInsertCache[key]
	jobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			jobAllColumns,
			jobColumnsWithDefault,
			jobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jobType, jobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"job\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"job\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into job")
	}

	if !cached {
		jobInsertCacheMut.Lock()
		jobInsertCache[key] = cache
		jobInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Job.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Job) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	jobUpdateCacheMut.RLock()
	cache, cached := jobUpdateCache[key]
	jobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			jobAllColumns,
			jobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update job, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"job\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, append(wl, jobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update job row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for job")
	}

	if !cached {
		jobUpdateCacheMut.Lock()
		jobUpdateCache[key] = cache
		jobUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q jobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for job")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for job")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JobSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"job\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, jobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in job slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all job")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Job) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(jobColumnsWithDefault, o)

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

	jobUpsertCacheMut.RLock()
	cache, cached := jobUpsertCache[key]
	jobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			jobAllColumns,
			jobColumnsWithDefault,
			jobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			jobAllColumns,
			jobPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert job, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jobPrimaryKeyColumns))
			copy(conflict, jobPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"job\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jobType, jobMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert job")
	}

	if !cached {
		jobUpsertCacheMut.Lock()
		jobUpsertCache[key] = cache
		jobUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Job record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Job) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Job provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jobPrimaryKeyMapping)
	sql := "DELETE FROM \"job\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from job")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for job")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q jobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no jobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from job")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JobSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(jobBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"job\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from job slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job")
	}

	if len(jobAfterDeleteHooks) != 0 {
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
func (o *Job) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JobSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"job\".* FROM \"job\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JobSlice")
	}

	*o = slice

	return nil
}

// JobExists checks if the Job row exists.
func JobExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"job\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if job exists")
	}

	return exists, nil
}
