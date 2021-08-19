package middleware

import (
	"github.com/kuhufu/tree"
	"log"
	"time"
)

func Logger() tree.HandlerFunc {
	return func(c *tree.Context) {
		begin := time.Now()
		c.Next()
		end := time.Now()

		log.Printf("[tree] %vms %v %v", end.Sub(begin).Milliseconds(), c.Response.Code(), c.Request.Path())
	}
}
