// +build mysql

package database

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
	ParentID uint	`gorm:"default:0;not null"`
	Kind string
	UrlPath string
	Sort int
	Summary string	`gorm:"type:varchar(1000)"`
	TemplatePath string
	PostTemplatePath string
	Status int                        `gorm:"type:tinyint(2);default:1;not null"`
	IsHidden int                      `gorm:"type:tinyint(2);default:0;not null"`
	IsMain int                        `gorm:"type:tinyint(2);default:0;not null"`
	// Attachment  []database.Attachment `gorm:"polymorphic:Owner;"`
}

type Post struct {
	gorm.Model
	Name string
	CategoryID uint  `gorm:"default:0;not null"`
	Content string `gorm:"type:mediumtext"`
	TemplatePath string
	UrlPath string
	Language string
	Summary string                    `gorm:"type:varchar(1000)"`
	Sort int                          `gorm:"default:0;not null"`
	Parameter string                  `gorm:"type:text"`
	Status int                        `gorm:"type:tinyint(2);default:1;not null"`
	// Attachment  []database.Attachment `gorm:"polymorphic:Owner;"`
}

type Variable struct {
	gorm.Model
	Name string
	Value string	`gorm:"type:varchar(5000)"`
	Remark string	`gorm:"type:varchar(1000)"`
}