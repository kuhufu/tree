package tree

import (
	"log"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		c.Next()
		log.Println("[path]", c.Request.Path())
	}
}
