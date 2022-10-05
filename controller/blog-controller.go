package controller

import (
	"parthivrmenon/conkraft-blog-json-server/entity"
	"parthivrmenon/conkraft-blog-json-server/service"
)

type BlogController interface {
	FindAll() entity.Blogs
}

type controller struct {
	service service.BlogService
}

func New(service service.BlogService) BlogController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() entity.Blogs {
	return c.service.FindAll()
}
