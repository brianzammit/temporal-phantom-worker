package configuration

type Activity struct {
	Type   string        `yaml:"type"`
	Result *interface{}  `yaml:"result"`
	Error  *ErrorDetails `yaml:"error"`
}

type Workflow struct {
	Type   string        `yaml:"type"`
	Result *interface{}  `yaml:"result"`
	Error  *ErrorDetails `yaml:"error"`
}

type ErrorDetails struct {
	Message string      `yaml:"message"`
	Type    string      `yaml:"type"`
	Details interface{} `yaml:"details"`
}

type Worker struct {
	Name       string     `yaml:"name"`
	TaskQueue  string     `yaml:"task_queue"`
	Workflows  []Workflow `yaml:"workflows"`
	Activities []Activity `yaml:"activities"`
}

type Server struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
}

type Config struct {
	Server  *Server  `yaml:"server"`
	Workers []Worker `yaml:"workers"`
}

func (a Activity) IsSuccess() bool {
	return a.Error == nil
}

func (w Workflow) IsSuccess() bool {
	return w.Error == nil
}
