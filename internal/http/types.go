package http

type ServerConfig struct {
	ListenAddress string `json:"addr"`
	Name          string `json:"name"`
	Config        struct {
		ServerHeader         string `json:"server_header"`
		StrictRouting        bool   `json:"strict_routing"`
		CaseSensitive        bool   `json:"case_sensitive"`
		UnescapePath         bool   `json:"unescape_path"`
		Etag                 bool   `json:"etag"`
		BodyLimit            int    `json:"body_limit"`
		Concurrency          int    `json:"concurrency"`
		ReadTimeout          int    `json:"read_timeout"`
		WriteTimeout         int    `json:"write_timeout"`
		IdleTimeout          int    `json:"idle_timeout"`
		ReadBufferSize       int    `json:"read_buffer_size"`
		WriteBufferSize      int    `json:"write_buffer_size"`
		CompressedFileSuffix string `json:"compressed_file_suffix"`
		GetOnly              bool   `json:"get_only"`
		DisableKeepalive     bool   `json:"disable_keepalive"`
		Network              string `json:"network"`
		EnablePrintRoutes    bool   `json:"enable_print_routes"`
		AttachErrorHandler   bool   `json:"attach_error_handler"`
	} `json:"conf"`
}
