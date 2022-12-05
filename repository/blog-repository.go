package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"parthivrmenon/conkraft-blog-json-server/entity"
)

type BlogRepository interface {
	Retrieve() entity.Blogs
	GetBlogBySlug(slug string) entity.Blog
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

func (repo *fileBlogRepository) GetBlogBySlug(slug string) entity.Blog {
	for i := 0; i < len(repo.blogs.Blogs); i++ {
		if repo.blogs.Blogs[i].Slug == slug {
			fmt.Println("found blog with slug ", slug)
			dat, err := os.ReadFile(repo.blogs.Blogs[i].File)
			if err != nil {
				panic(err)
			}
			fmt.Print(string(dat))
			repo.blogs.Blogs[i].Body = string(dat)

			return repo.blogs.Blogs[i]
		}

	}
	return entity.Blog{}

}
