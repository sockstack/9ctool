package gin

import "github.com/gin-gonic/gin"

type GroupOption struct {
	Prefix string
	Middlewares []gin.HandlerFunc
}

type GroupHandleFunc func(opts *GroupOption)

//WithPrefix 添加路由前缀
func WithPrefix(prefix string) GroupHandleFunc {
	return func(opts *GroupOption) {
		opts.Prefix = prefix
	}
}

//WithMiddleware 添加中间件
func WithMiddleware(middleware gin.HandlerFunc) GroupHandleFunc {
	return func(opts *GroupOption) {
		opts.Middlewares = append(opts.Middlewares, middleware)
	}
}
