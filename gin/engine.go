package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Engine struct {
	engine *gin.Engine
	server *http.Server
}

func NewEngine(handles ...EngineOptionHandle) *Engine {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        gin.Default(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	option := &EngineOption{Server: s}

	for _, handle := range handles {
		handle(option)
	}

	return &Engine{
		engine: option.Server.Handler.(*gin.Engine),
		server: option.Server,
	}
}


func (engine *Engine) Register(handles ...ServerHandle) {
	for _, handle := range handles {
		handle.Route(engine)
	}
}

func (engine *Engine) Start() {
	engine.server.ListenAndServe()
}
