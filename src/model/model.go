package model

type Post struct {
	ID          string `db:"id" form:"id"`
	Title       string `db:"title" form:"title"`
	Description string `db:"description" form:"description"`
	ImageSource string `db:"image_src" form:"image_src"`
}

type Response struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
