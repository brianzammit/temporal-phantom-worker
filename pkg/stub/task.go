package stub

type Task struct {
	Type   string
	result *ResultTemplate
	error  *Error
}

type Error struct {
	Type    string
	Message *ResultTemplate
	Details *ResultTemplate
}

func NewSuccessTask(taskType string, taskResult interface{}) (*Task, error) {
	resultTemplate := ResultTemplate{
		ResultSpec: taskResult,
	}

	err := resultTemplate.init()
	if err != nil {
		return nil, err
	}

	return &Task{
		Type:   taskType,
		result: &resultTemplate,
	}, nil
}

func NewErrorTask(taskType string, errorType string, message interface{}, details interface{}) (*Task, error) {
	messageTemplate := ResultTemplate{
		ResultSpec: message,
	}
	err := messageTemplate.init()
	if err != nil {
		return nil, err
	}

	detailsTemplate := ResultTemplate{
		ResultSpec: details,
	}
	err = detailsTemplate.init()
	if err != nil {
		return nil, err
	}

	return &Task{
		Type: taskType,
		error: &Error{
			Type:    errorType,
			Message: &messageTemplate,
			Details: &detailsTemplate,
		},
	}, nil
}

func (t Task) isSuccess() bool {
	return t.error == nil
}
