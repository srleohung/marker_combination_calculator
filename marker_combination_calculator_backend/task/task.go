package task

import "time"

type Task struct {
	inProgress bool
	progress   int
	cancel     chan bool
}

func New() *Task {
	return &Task{
		inProgress: false,
		progress:   0,
		cancel:     make(chan bool),
	}
}

func (t *Task) Start() string {
	if t.inProgress {
		return "Another task in progress"
	} else {
		t.inProgress = !t.inProgress
		go t.calculate()
		return "Task started"
	}
}

func (t *Task) calculate() {
	for t.progress = 0; t.progress < 100; t.progress++ {
		select {
		case <-t.cancel:
		case <-time.After(50 * time.Millisecond):
			continue
		}
		break
	}
	t.inProgress = !t.inProgress
}

func (t *Task) GetProgress() int {
	return t.progress
}

func (t *Task) GetResult() string {
	if t.progress == 100 {
		return "Task completed"
	} else if t.inProgress {
		return "Task in progress"
	}
	return "Task does not exist"
}

func (t *Task) Cancel() string {
	if !t.inProgress {
		return "Task not in progress"
	} else if len(t.cancel) > 0 {
		return "Was asked to cancel before the api"
	}
	t.cancel <- true
	return "Cancel success"
}
