package service

import (
	"fmt"
	"myGram/entities"
	"myGram/repository"
)

type CommentService interface {
	GetAllComments() ([]entities.ResponseComment, error)
	GetCommentByID(commentID uint) (entities.ResponseComment, error)
	CreateComment(data entities.Comment) (entities.Comment, error)
	UpdateComment(data entities.Comment, commentID uint) (entities.Comment, error)
	DeleteComment(commentID uint) error
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{commentRepo: commentRepo}
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func (service *commentService) GetAllComments() ([]entities.ResponseComment, error) {
	resComment, err := service.commentRepo.GetComments()
	fmt.Println(resComment)
	if err != nil {
		return []entities.ResponseComment{}, err
	}
	var response []entities.ResponseComment
	for _, comment := range resComment {
		tempResp := entities.ResponseComment{}
		tempResp.ID = comment.ID
		tempResp.PhotoID = comment.PhotoID
		tempResp.Message = comment.Message
		tempResp.User.Username = comment.User.Username
		tempResp.User.Email = comment.User.Email
		tempResp.CreatedAt = *comment.CreatedAt
		tempResp.UpdatedAt = *comment.UpdatedAt
		tempResp.Photo.Title = comment.Photo.Title
		tempResp.Photo.Caption = comment.Photo.Caption
		tempResp.Photo.PhotoURL = comment.Photo.PhotoURL
		response = append(response, tempResp)
	}
	fmt.Println("p")
	return response, nil
}

func (service *commentService) GetCommentByID(CommentID uint) (entities.ResponseComment, error) {
	resComment, err := service.commentRepo.GetCommentByID(CommentID)
	fmt.Println(resComment)
	if err != nil {
		return entities.ResponseComment{}, err
	}
	var comment entities.ResponseComment
	comment.ID = resComment.ID
	comment.PhotoID = resComment.PhotoID
	comment.Message = resComment.Message
	comment.User.Username = resComment.User.Username
	comment.User.Email = resComment.User.Email
	comment.CreatedAt = *resComment.CreatedAt
	comment.UpdatedAt = *resComment.UpdatedAt
	comment.Photo.Title = resComment.Photo.Title
	comment.Photo.Caption = resComment.Photo.Caption
	comment.Photo.PhotoURL = resComment.Photo.PhotoURL
	return comment, nil
}

func (service *commentService) CreateComment(data entities.Comment) (entities.Comment, error) {
	create, err := service.commentRepo.CreateComment(data)
	if err != nil {
		return entities.Comment{}, err
	}
	return create, nil
}

func (service *commentService) UpdateComment(data entities.Comment, CommentID uint) (entities.Comment, error) {
	entityComment := entities.Comment{}
	entityComment.ID = uint(CommentID)
	entityComment.Message = data.Message
	update, err := service.commentRepo.UpdateComment(entityComment)
	if err != nil {
		return entities.Comment{}, err
	}
	return update, nil
}

func (service *commentService) DeleteComment(commentID uint) error {
	err := service.commentRepo.DeleteComment(commentID)
	if err != nil {
		return err
	}
	return nil
}
