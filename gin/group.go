package gin

import "github.com/gin-gonic/gin"

//type RouterGroup gin.RouterGroup

type Group struct {
	Prefix string
	Handles []ServerHandle
	Middlewares []gin.HandlerFunc
}

func NewGroup(handles ...GroupHandleFunc) *Group {
	opts := &GroupOption{
		Prefix:      "",
		Middlewares: nil,
	}
	for _, handle := range handles {
		handle(opts)
	}
	return &Group{
		Prefix:  opts.Prefix,
		Middlewares: opts.Middlewares,
	}
}

func (group *Group) RegisterHandle(handles ...ServerHandle) {
	group.Handles = handles
}


