package repository

import (
	"myGram/entities"

	"gorm.io/gorm"
)

type CommentRepository interface {
	GetComments() ([]entities.Comment, error)
	GetCommentByID(id uint) (entities.Comment, error)
	CreateComment(data entities.Comment) (entities.Comment, error)
	UpdateComment(data entities.Comment) (entities.Comment, error)
	DeleteComment(commentID uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r commentRepository) GetComments() ([]entities.Comment, error) {
	var comments []entities.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return []entities.Comment{}, err
	}
	return comments, nil
}

func (r *commentRepository) GetCommentByID(id uint) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", id).First(&comment).Error
	if err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}

func (r commentRepository) CreateComment(data entities.Comment) (entities.Comment, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entities.Comment{}, err
	}
	return data, nil
}

func (r *commentRepository) UpdateComment(data entities.Comment) (entities.Comment, error) {
	err := r.db.Updates(&data).First(&data).Error
	if err != nil {
		return entities.Comment{}, err
	}
	return data, nil
}

func (r *commentRepository) DeleteComment(commentID uint) error {
	var comment entities.Comment
	comment.ID = commentID
	err := r.db.First(&comment).Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
