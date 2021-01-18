package model

import (
	"gorm.io/gorm"
	"time"
)

type Attachment struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	StorageMethod string	`json:"storage_method"`
	Path string	`json:"path"`
	OwnerID uint	`json:"owner_id"`
	OwnerType string	`json:"owner_type"`
}

// Sort CmsCategory	[Go Sort package]
// 分类排序		[Go Sort包]
type attachment_SortAsc []Attachment
func (c attachment_SortAsc) Len() int           { return len(c) }
func (c attachment_SortAsc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c attachment_SortAsc) Less(i, j int) bool { return c[i].CreatedAt.Before(c[j].CreatedAt) }