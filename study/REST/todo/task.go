package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time // указатель, т.к. поле может быть nil (nil на момент создания)
}

func (t *Task) Complete() {
	completeTime := time.Now()
	t.Completed = true
	t.CompletedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
