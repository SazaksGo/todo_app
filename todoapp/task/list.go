package task

import "fmt"

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Title]; ok {
		return fmt.Errorf("task with title %s already exists", task.Title)
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) ListNotCompletedTasks() map[string]Task {
	NKT := make(map[string]Task, len(l.tasks))
	for title, tasks := range l.tasks {
		if !tasks.Complete {
			NKT[title] = tasks
		}
	}
	return NKT
}
func (l *List) CompleteTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return fmt.Errorf("task with title %s not found", title)
	}
	task.Completed()

	l.tasks[title] = task
	return nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return fmt.Errorf("task with title %s not found", title)
	}
	delete(l.tasks, title)
	return nil
}
