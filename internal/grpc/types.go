package grpc

type ServerConfig struct {
	Host     string                 `json:"host"`
	Port     int                    `json:"port"`
	Protocol string                 `json:"protocol"`
	Async    bool                   `json:"async"`
	Configs  map[string]interface{} `json:"configs"`
}
