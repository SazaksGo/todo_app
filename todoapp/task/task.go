package task

import "time"

type Task struct {
	ID           int
	Title        string
	Description  string
	Complete     bool
	Created_at   time.Time
	Completed_at *time.Time
}

func NewTask(title, description string) Task {
	return Task{
		Title:        title,
		Description:  description,
		Complete:     false,
		Created_at:   time.Now(),
		Completed_at: nil,
	}

}

func (t *Task) Completed() {

	CompleteTime := time.Now()

	t.Complete = true
	t.Completed_at = &CompleteTime
}

func (t *Task) UnComplete() {

	t.Complete = false
	t.Completed_at = nil
}
