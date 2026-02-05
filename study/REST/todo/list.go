package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

/**
многосоставная операция, для конкурентного доступа должна быть атомарной
именно поэтому, мы не можем создать отдельно блокировку на чтение, и отдельно на запись
при попадании в условие, несколько горутин могут пройти его одновременно
*/

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock() // defer всегда разблокирует мьютекс, т.к. выполнится при любом завершении метода

	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil

}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	// обязательно переопредлеяем мапу, чтобы извне функции не было доступа к оригинальной
	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v // tmp не является разделяемым ресурсом, т.к. обладает областью видимости метода
	}

	return tmp
}

func (l *List) ListUncompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	unCompletedTasks := make(map[string]Task, len(l.tasks))

	for title, task := range l.tasks {
		if !task.Completed {
			unCompletedTasks[title] = task
		}
	}

	return unCompletedTasks
}

func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Complete()
	l.tasks[task.Title] = task

	return l.tasks[title], nil

}

func (l *List) UncompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Uncomplete()
	return l.tasks[title], nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title) // изменение состояние разделяемого ресурса приравнивается к записи в этот ресурс

	return nil
}
