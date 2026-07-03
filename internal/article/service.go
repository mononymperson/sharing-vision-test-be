package article

import (
	"database/sql"
	"errors"
)

type ArticleService struct {
	repo *ArticleRepository
}

func NewArticleService(repo *ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(req *CreateArticleRequest) (*Article, error) {
	article := &Article{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}
	if err := s.repo.Create(article); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) FindAll(limit, offset int) ([]Article, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *ArticleService) FindByID(id int) (*Article, error) {
	article, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) Update(id int, req *UpdateArticleRequest) (*Article, error) {
	article, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}

	article.Title = req.Title
	article.Content = req.Content
	article.Category = req.Category
	article.Status = req.Status

	if err := s.repo.Update(article); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) Delete(id int) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("article not found")
		}
		return err
	}
	return s.repo.Delete(id)
}
