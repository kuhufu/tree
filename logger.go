package tree

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		begin := time.Now()
		c.Next()
		log.Println("[path]", c.Request.Path())
		end := time.Now()

		log.Printf("[tree] %vms %v %v", end.Sub(begin).Milliseconds(), c.Response.Code(), c.Request.Path())
	}
}
