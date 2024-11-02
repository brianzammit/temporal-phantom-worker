package configuration

type Activity struct {
	Type   string        `json:"type"`
	Result *interface{}  `json:"result"`
	Error  *ErrorDetails `json:"error"`
}

type Workflow struct {
	Type   string        `json:"type"`
	Result *interface{}  `json:"result"`
	Error  *ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Details interface{} `json:"details"`
}

type Worker struct {
	Name       string     `json:"name"`
	TaskQueue  string     `yaml:"task_queue"`
	Workflows  []Workflow `yaml:"workflows"`
	Activities []Activity `yaml:"activities"`
}

type Config struct {
	Workers []Worker `yaml:"workers"`
}

func (a Activity) IsSuccess() bool {
	return a.Error == nil
}

func (w Workflow) IsSuccess() bool {
	return w.Error == nil
}
