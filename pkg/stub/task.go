package stub

type Task struct {
	Type   string
	result *interface{}
	error  *Error
}

type Error struct {
	Type    string
	Message string
	Details interface{}
}

func NewSuccessTask(taskType string, taskResult interface{}) Task {
	return Task{
		Type:   taskType,
		result: &taskResult,
	}
}

func NewErrorTask(taskType string, taskError Error) Task {
	return Task{
		Type:  taskType,
		error: &taskError,
	}
}

func (t Task) isSuccess() bool {
	return t.error == nil
}
