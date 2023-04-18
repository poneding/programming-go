package middleware_demo

import "fmt"

type Context struct {
	handlers []func(c *Context)
	index    uint8
}

func (c *Context) Next() {
	c.index++
	c.handlers[c.index](c)
}

func (c *Context) Use(f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func (c *Context) Run() {
	c.handlers[0](c)
}

func (c *Context) Get(path string, f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func Do() {
	c := new(Context)
	c.Use(middleware1(c))
	c.Use(middleware2(c))

	c.Get("/", func(c *Context) {
		fmt.Println("hello world")
	})

	c.Run()
}

func middleware1(c *Context) func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware 1 in")
		c.Next()
		fmt.Println("middleware 1 out")
	}
}

func middleware2(c *Context) func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware 2 in")
		c.Next()
		fmt.Println("middleware 2 out")
	}
}
