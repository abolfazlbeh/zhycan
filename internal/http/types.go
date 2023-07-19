package http

import "github.com/gofiber/fiber/v2"

type ServerConfig struct {
	ListenAddress string   `json:"addr"`
	Name          string   `json:"name"`
	Versions      []string `json:"versions"`
	SupportStatic bool     `json:"support_static"`
	Config        struct {
		ServerHeader         string   `json:"server_header"`
		StrictRouting        bool     `json:"strict_routing"`
		CaseSensitive        bool     `json:"case_sensitive"`
		UnescapePath         bool     `json:"unescape_path"`
		Etag                 bool     `json:"etag"`
		BodyLimit            int      `json:"body_limit"`
		Concurrency          int      `json:"concurrency"`
		ReadTimeout          int      `json:"read_timeout"`
		WriteTimeout         int      `json:"write_timeout"`
		IdleTimeout          int      `json:"idle_timeout"`
		ReadBufferSize       int      `json:"read_buffer_size"`
		WriteBufferSize      int      `json:"write_buffer_size"`
		CompressedFileSuffix string   `json:"compressed_file_suffix"`
		GetOnly              bool     `json:"get_only"`
		DisableKeepalive     bool     `json:"disable_keepalive"`
		Network              string   `json:"network"`
		EnablePrintRoutes    bool     `json:"enable_print_routes"`
		AttachErrorHandler   bool     `json:"attach_error_handler"`
		RequestMethods       []string `json:"request_methods"`
	} `json:"conf"`
	Middlewares struct {
		Order []string `json:"order"`
	} `json:"middlewares"`
	Static struct {
		Prefix string       `json:"prefix"`
		Root   string       `json:"root"`
		Config fiber.Static `json:"config"`
	} `json:"static"`
}

// LoggerMiddlewareConfig - defines the config for middleware.
type LoggerMiddlewareConfig struct {
	Format       string `json:"format"`
	TimeFormat   string `json:"time_format"`
	TimeZone     string `json:"time_zone"`
	TimeInterval int    `json:"time_interval"`
	Output       string `json:"output"`
}

type FaviconMiddlewareConfig struct {
	File         string `json:"file"`
	URL          string `json:"url"`
	CacheControl string `json:"cache_control"`
}
