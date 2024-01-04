// Package domain provides
package domain

import "time"

// Task .
type Task struct {
	ID        int64
	Name      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TaskParam .
type TaskParam struct {

	// implement query task condition, like page, per_page, sort_by ...etc
}
