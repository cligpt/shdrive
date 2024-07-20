package config

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type MetaData struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Ollama Ollama `yaml:"ollama"`
	Studio Studio `yaml:"studio"`
}

type Ollama struct {
	Url string `yaml:"url"`
}

type Studio struct {
	Url string `yaml:"url"`
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
