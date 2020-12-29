package model

import (
	"encoding/json"
	em "github.com/Etpmls/Etpmls-Micro"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name string	`json:"name"`
	CategoryID uint	`json:"category_id"`
	Content string	`json:"content"`
	TemplatePath string	`json:"template_path"`
	UrlPath string	`json:"url_path"`
	Language string	`json:"language"`
	Summary string	`json:"summary"`
	Parameter string `json:"parameter"`
	Sort int	`json:"sort"`
	Status int	`json:"status"`
	// tmp field
	Thumbnail []Attachment `gorm:"polymorphic:Owner;polymorphicValue:post-thumbnail" json:"thumbnail"`
	Category  Category           `gorm:"foreignKey:CategoryID" json:"category"`
}

// interface conversion User
// interface转换User
func (this *Post) InterfaceToPost(i interface{}) (Post, error) {
	var p Post
	us, err := json.Marshal(i)
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum("Object to JSON failed! err:" + err.Error()))
		return Post{}, err
	}
	err = json.Unmarshal(us, &p)
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum("JSON conversion object failed! err:" + err.Error()))
		return Post{}, err
	}
	return p, nil
}