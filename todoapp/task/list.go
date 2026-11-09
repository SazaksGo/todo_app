package task

import (
	"fmt"
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.tasks[task.Title]; ok {

		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task

	return nil
}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, fmt.Errorf("task with title %s not found", title)
	}
	return task, nil

}

func (l *List) ListNotCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	NKT := make(map[string]Task, len(l.tasks))
	for title, tasks := range l.tasks {
		if !tasks.Complete {
			NKT[title] = tasks
		}
	}
	return NKT
}
func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.Completed()

	l.tasks[title] = task
	return l.tasks[title], nil
}

func (l *List) UnCompleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.UnComplete()
	l.tasks[title] = task
	return nil

}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.tasks[title]
	if !ok {
		return fmt.Errorf("task with title %s not found", title)
	}
	delete(l.tasks, title)
	return nil
}
