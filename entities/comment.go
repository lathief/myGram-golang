package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an comment
type Comment struct {
	GormModel
	UserID  uint
	User    *User
	PhotoID uint `json:"photo_id" form:"photo_id"`
	Photo   *Photo
	Message string `json:"message" form:"message" valid:"required~Comment is required"`
}
type RequestComment struct {
	Message string `json:"message" example:"A Photo"`
	PhotoID uint   `json:"photo_id,omitempty" example:"1"`
	UserID  uint   `json:"user_id,omitempty" swaggerignore:"true"`
}
type ResponseComment struct {
	ID        uint      `json:"id" example:"1"`
	Message   string    `json:"message" example:"Nice Photo"`
	PhotoID   uint      `json:"photo_id" example:"1"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-08-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-08-01T00:00:00Z"`
	User      struct {
		Email    string `json:"email" example:"test@example.com"`
		Username string `json:"username" example:"test"`
	} `json:"user"`
	Photo struct {
		Title    string `json:"title" example:"A Photo"`
		Caption  string `json:"caption" example:"My Photo"`
		PhotoURL string `json:"photo_url" example:"https://example.com/photo.jpg"`
	} `json:"photo"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
