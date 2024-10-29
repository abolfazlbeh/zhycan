package grpc

type ServerConfig struct {
	Host       string                 `json:"host"`
	Port       int                    `json:"port"`
	Protocol   string                 `json:"protocol"`
	Async      bool                   `json:"async"`
	Reflection bool                   `json:"reflection"`
	Configs    map[string]interface{} `json:"configs"`
}
