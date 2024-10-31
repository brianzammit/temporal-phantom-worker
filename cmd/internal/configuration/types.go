package configuration

type Activity struct {
	Type   string      `json:"type"`
	Result interface{} `json:"result"`
}

type Workflow struct {
	Type   string      `json:"type"`
	Result interface{} `json:"result"`
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
