package article

import "database/sql"

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *Article) error {
	result, err := r.db.Exec(
		`INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)`,
		article.Title, article.Content, article.Category, article.Status,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	article.ID = int(id)
	return nil
}

func (r *ArticleRepository) FindAll(limit, offset int) ([]Article, error) {
	rows, err := r.db.Query(
		`SELECT id, title, content, category, created_date, updated_date, status
		 FROM posts ORDER BY created_date DESC LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Category,
			&a.CreatedDate, &a.UpdatedDate, &a.Status)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, rows.Err()
}

func (r *ArticleRepository) FindByID(id int) (*Article, error) {
	var a Article
	err := r.db.QueryRow(
		`SELECT id, title, content, category, created_date, updated_date, status
		 FROM posts WHERE id = ?`, id,
	).Scan(&a.ID, &a.Title, &a.Content, &a.Category,
		&a.CreatedDate, &a.UpdatedDate, &a.Status)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *ArticleRepository) Update(article *Article) error {
	_, err := r.db.Exec(
		`UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?`,
		article.Title, article.Content, article.Category, article.Status, article.ID,
	)
	return err
}

func (r *ArticleRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM posts WHERE id = ?`, id)
	return err
}
