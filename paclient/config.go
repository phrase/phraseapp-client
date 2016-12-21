package paclient

type Config struct {
	Host        string                            `yaml:"host,omitempty"`
	AccessToken string                            `yaml:"access_token,omitempty"`
	ProjectID   string                            `yaml:"project_id"`
	FileFormat  string                            `yaml:"file_format,omitempty"`
	PerPage     int                               `yaml:"per_page,omitempty"`
	Defaults    map[string]map[string]interface{} `yaml:"defaults,omitempty"`
	Push        PushYAML                          `yaml:"push,omitempty"`
	Pull        PullYAML                          `yaml:"pull,omitempty"`
}

type PushYAML struct {
	Sources []SourcesYAML `yaml:"sources,omitempty"`
}

type PullYAML struct {
	Targets []TargetsYAML `yaml:"targets,omitempty"`
}

type SourcesYAML struct {
	File   string                 `yaml:"file,omitempty"`
	Params map[string]interface{} `yaml:"params,omitempty"`
}

type TargetsYAML SourcesYAML
