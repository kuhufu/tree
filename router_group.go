package tree

import "path"

type IRouter interface {
	Handle(path string, handlerFunc ...HandlerFunc)
}

type RouterGroup struct {
	handlers HandlersChain
	basePath string
	engine   *Engine
}

func (g *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		handlers: g.combineHandlers(handlers),
		basePath: g.calculateAbsolutePath(relativePath),
		engine:   g.engine,
	}
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) *RouterGroup {
	g.handlers = g.combineHandlers(middleware)
	return g
}

func (g *RouterGroup) Handle(path string, handlerFunc ...HandlerFunc) {
	g.handle(path, handlerFunc)
}

func (g *RouterGroup) handle(relativePath string, handlers HandlersChain) {
	absolutePath := g.calculateAbsolutePath(relativePath)
	handlers = g.combineHandlers(handlers)
	g.engine.addRoute(absolutePath, handlers)
}

func (g *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	res := make(HandlersChain, 0, len(g.handlers)+len(handlers))
	res = append(res, g.handlers...)
	res = append(res, handlers...)
	return res
}

func (g *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return path.Join(g.basePath, relativePath)
}

func (g *RouterGroup) BasePath() string {
	return g.basePath
}
