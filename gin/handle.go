package gin

type ServerHandle interface {
	Route(engine *Engine) error
}
