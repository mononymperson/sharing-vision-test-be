package article

import "time"

/*
SQL Migration:
CREATE TABLE posts (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(200)  NOT NULL,
    content     TEXT          NOT NULL,
    category    VARCHAR(100)  NOT NULL,
    created_date DATETIME     NOT NULL,
    updated_date DATETIME     NOT NULL,
    status      VARCHAR(100)  NOT NULL DEFAULT 'draft'
);
*/

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,min=20"`
	Content     string    `json:"content" validate:"required,min=200"`
	Category    string    `json:"category" validate:"required,min=3"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Status      string    `json:"status" validate:"required,oneof=publish draft thrash"`
}

type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}

type UpdateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=publish draft thrash"`
}
