package entity

type Blog struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type Blogs struct {
	Blogs []Blog `json:"blogs"`
}
