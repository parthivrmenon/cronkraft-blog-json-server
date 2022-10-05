package service

import (
	"parthivrmenon/conkraft-blog-json-server/entity"
	"parthivrmenon/conkraft-blog-json-server/repository"
)

type BlogService interface {
	FindAll() entity.Blogs
}

type blogService struct {
	repo repository.BlogRepository
}

func New() BlogService {
	var repo repository.BlogRepository = repository.New()
	return &blogService{
		repo: repo,
	}
}

func (service *blogService) FindAll() entity.Blogs {

	return service.repo.Retrieve()
}
