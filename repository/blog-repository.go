package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"parthivrmenon/conkraft-blog-json-server/entity"
)

type BlogRepository interface {
	Retrieve() entity.Blogs
}

type fileBlogRepository struct {
	blogs entity.Blogs
}

func New() BlogRepository {
	fmt.Println("loading blogs from db.json...")

	file, err := ioutil.ReadFile("db.json")
	if err != nil {
		log.Fatal(err)
	}
	blogs := entity.Blogs{}
	_ = json.Unmarshal([]byte(file), &blogs)

	for i := 0; i < len(blogs.Blogs); i++ {

		fmt.Println("Title: ", blogs.Blogs[i].Title)
	}

	return &fileBlogRepository{
		blogs: blogs,
	}
}

func (repo *fileBlogRepository) Retrieve() entity.Blogs {
	return repo.blogs
}
