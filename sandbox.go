package main

import (
	"fmt"
	"time"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID          int
	Title       string
	Contents    string
	UserName    string
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

func main() {
	comment := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "this is the message",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "this is the message 2",
		CreatedAt: time.Now(),
	}

	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{comment, comment2},
		CreatedAt:   time.Now(),
	}

	// fmt.Printf("Comment Struct: %+v)\n", comment)
	fmt.Printf("%+v\n", article)

}
