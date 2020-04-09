package tree

import (
	"fmt"
	. "github.com/kuhufu/tree/ds"
	"sync"
)

type HandlerFunc func(c *Context)

type Path string
type Method string

type HandlersChain []HandlerFunc

var defaultNoRouter = func(c *Context) {
	c.String(404, "", fmt.Sprintf("%v not found", c.Request.Path()))
}

type Engine struct {
	RouterGroup
	router  map[Path]map[Method]HandlersChain
	noRoute HandlersChain
	pool    sync.Pool
}

func New() *Engine {
	engine := &Engine{
		router:  map[Path]map[Method]HandlersChain{},
		noRoute: HandlersChain{defaultNoRouter},
		pool: sync.Pool{New: func() interface{} {
			return &Context{}
		}},
	}

	engine.RouterGroup.engine = engine
	engine.RouterGroup.basePath = "/"

	return engine
}

func Default() *Engine {
	engine := New()

	engine.Use(Logger(), Recovery())

	return engine
}

func (e *Engine) addRoute(method string, path string, handlers []HandlerFunc) {
	fmt.Printf("[handler] %v %v\n", method, path)
	p := Path(path)
	m, ok := e.router[p]
	if !ok {
		m = make(map[Method]HandlersChain)
		e.router[p] = m
	}

	m[Method(method)] = handlers
}

func (e *Engine) Route(request Request, response Response) {
	c := e.pool.Get().(*Context)
	c.reset()

	c.Request = request
	c.Response = response

	e.route(c)

	e.pool.Put(c)
}

func (e *Engine) route(c *Context) {
	c.handlers = e.getHandlers(c)

	for _, handler := range c.handlers {
		if handler == nil {
			continue
		}

		handler(c)
		if c.IsAborted() {
			return
		}
	}
}

func (e *Engine) NoRoute(handlerFunc ...HandlerFunc) {
	e.noRoute = handlerFunc
}

func (e *Engine) getHandlers(c *Context) HandlersChain {
	handlers, ok := e.router[Path(c.Request.Path())][Method(c.Request.Method())]
	if !ok {
		return e.noRoute
	}

	return handlers
}
