package tree

import (
	"github.com/kuhufu/tree/binding"
	. "github.com/kuhufu/tree/ds"
	"log"
)

type Context struct {
	Request  Request
	Response Response

	handlers HandlersChain

	index int

	keys map[string]interface{}
}

func (c *Context) SetResp(code int, contentType ContentType, val interface{}, msg string) {
	c.Response.SetContentType(contentType)
	bodyBytes, err := Encode(contentType, val)
	if err != nil {
		log.Println("序列化错误", err)
		return
	}
	c.Response.SetBody(bodyBytes)
	c.Response.SetCode(code)
	c.Response.SetMsg(msg)
}
func (c *Context) String(code int, val string, msg string) {
	contentType := ContentType_Text
	c.SetResp(code, contentType, val, msg)
}

func (c *Context) JSON(code int, obj interface{}, msg string) {
	contentType := ContentType_JSON
	c.SetResp(code, contentType, obj, msg)
}

func (c *Context) Binary(code int, data []byte, msg string) {
	contentType := ContentType_Binary
	c.SetResp(code, contentType, data, msg)
}

func (c *Context) ProtoBuf(code int, val interface{}, msg string) {
	contentType := ContentType_ProtoBuf
	c.SetResp(code, contentType, val, msg)
}

func (c *Context) ShouldBindJSON(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.JSON)
}

func (c *Context) ShouldBindProtoBuf(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.ProtoBuf)
}

func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return b.Bind(c.Request, obj)
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

func (c *Context) AbortWithStatus(code int, msg ...string) {
	c.Abort()

	var m string
	if len(msg) != 0 {
		m = msg[0]
	}
	c.JSON(code, nil, m)
}

func (c *Context) IsAborted() bool {
	return c.index >= len(c.handlers)
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Get(key string) (val interface{}, exist bool) {
	val, exist = c.keys[key]

	return val, exist
}

func (c *Context) Set(key string, val interface{}) {
	if c.keys == nil {
		c.keys = make(map[string]interface{})
	}

	c.keys[key] = val
}

func (c *Context) reset() {
	c.Request = nil
	c.Response = nil
	c.handlers = nil
	c.index = 0
	c.keys = nil
}
