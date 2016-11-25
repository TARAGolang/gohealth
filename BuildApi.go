package gohealth

import "github.com/fulldump/golax"

func BuildApi(parent *golax.Node, m *MonitorWatch) {

	h := NewHandler(m)

	parent.Node("health").Method("GET", func(c *golax.Context) {
		h(c.Response, c.Request)
	})
}
