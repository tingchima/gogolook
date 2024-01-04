// Package postgres provides
package postgres

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tingchima/gogolook/internal/domain"
	"github.com/tingchima/gogolook/testdata"
)

// TestTaskRepo_ListTasks .
func TestTaskRepo_ListTasks(t *testing.T) {

	conn := getTestDBConn()

	err := setupTestData(conn, testdata.Path(testdata.TestDataTasks))
	require.NoError(t, err)

	repo := NewRepository(conn)

	_, err = repo.ListTasks(context.Background(), domain.TaskParam{})
	require.NoError(t, err)
}

// TestTaskRepo_CreateTask .
func TestTaskRepo_CreateTask(t *testing.T) {

	repo := NewRepository(getTestDBConn())

	// Args .
	type Args struct {
		Task domain.Task
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	got, err := repo.CreateTask(context.Background(), args.Task)
	require.NoError(t, err)

	assert.Equal(t, args.Task.Name, got.Name)
	assert.Equal(t, args.Task.Status, got.Status)
}

// TestTaskRepo_UpdateTask .
func TestTaskRepo_UpdateTask(t *testing.T) {

	repo := NewRepository(getTestDBConn())

	// Args .
	type Args struct {
		Task domain.Task
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	createdTask, err := repo.CreateTask(context.Background(), args.Task)
	require.NoError(t, err)

	updates := domain.Task{
		ID:     createdTask.ID,
		Name:   "updated_name",
		Status: true,
	}

	updatedTask, err := repo.UpdateTask(context.Background(), updates)
	require.NoError(t, err)

	assert.Equal(t, updates.Name, updatedTask.Name)
	assert.Equal(t, updates.Status, updatedTask.Status)
}

// TestTaskRepo_DeleteTask .
func TestTaskRepo_DeleteTask(t *testing.T) {

	repo := NewRepository(getTestDBConn())

	// Args .
	type Args struct {
		Task domain.Task
	}

	var args Args

	err := faker.FakeData(&args)
	require.NoError(t, err)

	createdTask, err := repo.CreateTask(context.Background(), args.Task)
	require.NoError(t, err)

	err = repo.DeleteTaskByID(context.Background(), createdTask.ID)
	require.NoError(t, err)
}
