package gin

import "net/http"

type EngineOption struct {
	Server *http.Server
}

type EngineOptionHandle func(opts *EngineOption)

//WithHttpServer
func WithHttpServer(server *http.Server) EngineOptionHandle {
	return func(opts *EngineOption) {
		opts.Server = server
	}
}
