// Package postgres provides
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/tingchima/gogolook/internal/domain"
	"github.com/tingchima/gogolook/internal/domain/common"
)

var (
	ErrNotFoundTask = errors.New("task not found")
)

// repoTask .
type repoTask struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Status    bool         `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// toTask convert repo struct to domain struct
func (row repoTask) toTask() domain.Task {

	return domain.Task{
		ID:        row.ID,
		Name:      row.Name,
		Status:    row.Status,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt.Time,
	}
}

// table name
const repoTableTask = "tasks"

type repoFieldNameTask struct {
	ID        string
	Name      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

var repoFieldTask = repoFieldNameTask{
	ID:        "id",
	Name:      "name",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (r *repoFieldNameTask) fields() []string {
	return []string{
		r.ID,
		r.Name,
		r.Status,
		r.CreatedAt,
		r.UpdatedAt,
	}
}

// 列出任務
func (r *Postgres) ListTasks(ctx context.Context, param domain.TaskParam) ([]domain.Task, error) {

	wheres := squirrel.And{
		// implement select tasks condition
	}

	query, args, err := r.stmtBuilder.Select(repoFieldTask.fields()...).
		From(repoTableTask).
		Where(wheres).
		ToSql()
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	var rows []repoTask

	if err = r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	tasks := make([]domain.Task, len(rows))

	for i := range rows {
		tasks[i] = rows[i].toTask()
	}

	return tasks, nil
}

// 建立任務
func (r *Postgres) CreateTask(ctx context.Context, param domain.Task) (*domain.Task, error) {

	insertBuilder := r.stmtBuilder.Insert(repoTableTask).Columns(
		repoFieldTask.Name,
		repoFieldTask.Status,
	)

	insertBuilder = insertBuilder.Values(
		param.Name,
		param.Status,
	)

	query, args, err := insertBuilder.
		Suffix(fmt.Sprintf("returning %s", strings.Join(repoFieldTask.fields(), ", "))).
		ToSql()
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	var row repoTask

	err = r.db.GetContext(ctx, &row, query, args...)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	task := row.toTask()

	return &task, nil
}

// 修改任務
func (r *Postgres) UpdateTask(ctx context.Context, param domain.Task) (*domain.Task, error) {

	where := squirrel.And{
		squirrel.Eq{repoFieldTask.ID: param.ID},
	}

	updates := map[string]any{
		repoFieldTask.Name:      param.Name,
		repoFieldTask.Status:    param.Status,
		repoFieldTask.UpdatedAt: time.Now().UTC(),
	}

	query, args, err := r.stmtBuilder.Update(repoTableTask).
		Where(where).
		SetMap(updates).
		Suffix(fmt.Sprintf("returning %s", strings.Join(repoFieldTask.fields(), ", "))).
		ToSql()
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	var row repoTask

	err = r.db.GetContext(ctx, &row, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrNotFoundTask
			return nil, common.NewError(common.ErrCodeResourceNotFound, err, common.WithMsg(err.Error()))
		}
		return nil, common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	task := row.toTask()

	return &task, nil
}

// 透過ID刪除任務
func (r *Postgres) DeleteTaskByID(ctx context.Context, id int64) error {

	where := squirrel.And{
		squirrel.Eq{repoFieldTask.ID: id},
	}

	query, args, err := r.stmtBuilder.Delete(repoTableTask).Where(where).ToSql()
	if err != nil {
		return common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	affects, err := result.RowsAffected()
	if err != nil {
		return common.NewError(common.ErrCodeInternalProcess, err, common.WithMsg(err.Error()))
	}

	if affects == 0 {
		err := ErrNotFoundTask
		return common.NewError(common.ErrCodeResourceNotFound, err, common.WithMsg(err.Error()))
	}

	return nil
}
