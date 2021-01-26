package model

import (
	"context"
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/client"
	em "github.com/Etpmls/Etpmls-Micro"
	"gorm.io/gorm"
	"sort"
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
	Thumbnail []Attachment `gorm:"-" json:"thumbnail"`
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

func (this *Post) WithAttachment(ctx *context.Context, c []Post, owner_type string) (error) {
	// 1.Get all ids
	var f func (v []Post, ids *[]uint32)
	f = func (v []Post, ids *[]uint32) {
		for _, sv := range v {
			*ids = append(*ids, uint32(sv.ID))
		}
	}
	var ids []uint32
	f(c, &ids)

	// 2.Get all thumbnail
	b, err := client.NewClient().Attachment_GetMany(ctx, ids, application.Relationship_Post_Thumbnail)
	if err != nil {
		return err
	}

	var tmp []Attachment
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		em.LogError.OutputSimplePath(err)
		return err
	}

	if len(tmp) == 0 {
		return nil
	}

	var f2 func(cat *[]Post)
	f2 = func(cat *[]Post) {
		for k, v := range *cat {
			// Find whether the current attachment contains a thumbnail
			for i := 0; i < len(tmp); i++ {
				if v.ID == tmp[i].OwnerID {
					(*cat)[k].Thumbnail = append((*cat)[k].Thumbnail, tmp[i])
					// Delete slice
					tmp = append(tmp[:i], tmp[i+1:]...)
					i--
				}
			}
		}
	}

	f2(&c)

	return nil
}

func (this *Post) AttachmentSortAsc(slice []Post) {
	for _, v := range slice {
		sort.Sort(attachment_SortAsc(v.Thumbnail))
	}
	return
}

