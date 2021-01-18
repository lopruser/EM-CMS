package model

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/client"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/Etpmls/Etpmls-Micro/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sort"
	"strings"
	"time"
)

type Category struct {
	ID        uint                `json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt gorm.DeletedAt      `json:"deleted_at"`
	Name string                   `json:"name"`
	ParentID uint                 `json:"parent_id"`
	Kind string                   `json:"kind"`
	UrlPath string                `json:"url_path"`
	Sort int                      `json:"sort"`
	Summary string                `json:"summary"`
	TemplatePath string           `json:"template_path"`
	PostTemplatePath string       `json:"post_template_path"`
	Status int                    `json:"status"`
	IsHidden int                  `json:"is_hidden"`
	IsMain int                    `json:"is_main"`
	Children []Category           `gorm:"-" json:"children,omitempty"`
	Thumbnail  []Attachment `gorm:"-" json:"thumbnail"`
	Posts []Post				  `gorm:"foreignKey:CategoryID"`
}

// interface conversion User
// interface转换User
func (this *Category) InterfaceToCategory(i interface{}) (Category, error) {
	var c Category
	us, err := json.Marshal(i)
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum("Object to JSON failed! err:" + err.Error()))
		return Category{}, err
	}
	err = json.Unmarshal(us, &c)
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum("JSON conversion object failed! err:" + err.Error()))
		return Category{}, err
	}
	return c, nil
}

// Get All Category
// 获取所有Category
func (this *Category) GetAll() ([]Category, error) {
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		return this.getAll_Cache()
	} else {
		return this.getAll_NoCache()
	}
}
func (this *Category) getAll_Cache() ([]Category, error) {
	s, err := em.Cache.GetString(application.Cache_CmsCategoryGetAll)
	if err != nil {
		if err == redis.Nil {
			return this.getAll_NoCache()
		}
		em.LogError.Output(em.MessageWithLineNum(err.Error()))
		return nil, err
	}

	var list []Category
	err = json.Unmarshal([]byte(s), &list)
	if err != nil {
		return nil, err
	}

	em.LogDebug.Output(em.MessageWithLineNum("Get CmsCategory from the cache according to UrlPath."))	// debug

	return list, err
}
func (this *Category) getAll_NoCache() ([]Category, error) {
	em.LogDebug.Output(em.MessageWithLineNum("CmsCategory cache not found"))	// debug

	var list []Category
	/*em.DB.Preload("Thumbnail", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at ASC")
	}).Find(&list)*/

	em.DB.Find(&list)

	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		b, err := json.Marshal(list)
		if err != nil {
			em.LogError.Output(em.MessageWithLineNum(err.Error()))
			return nil, err
		}
		em.Cache.SetString(application.Cache_CmsCategoryGetAll, string(b), 0)
	}

	return list, nil
}

// Get all Category trees
// 获取所有Category树
func (this *Category) GetTree() ([]Category, error) {
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		return this.getTree_Cache()
	} else {
		return this.getTree_NoCache()
	}
}
func (this *Category) getTree_NoCache() ([]Category, error) {
	em.LogDebug.Output(em.MessageWithLineNum("Category-Tree cache not found"))	// debug
	data, err := this.GetAll()
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum(err.Error()))
		return nil, err
	}

	var category Category
	tree := category.MakeTree(data)
	category.SortDesc(tree)

	// []Category -> []ApiCategoryGetAllV2
	jsonBytes, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}

	var newTree []Category
	err = json.Unmarshal([]byte(string(jsonBytes)), &newTree)
	if err != nil {
		return nil, err
	}

	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		b, err := json.Marshal(newTree)
		if err != nil {
			em.LogError.Output(err)
		} else {
			em.Cache.SetString(application.Cache_CmsCategoryGetTree, string(b), 0)
		}
	}

	return newTree, nil
}
func (this *Category) getTree_Cache() ([]Category, error) {
	j, err := em.Cache.GetString(application.Cache_CmsCategoryGetTree)
	if err != nil {
		return this.getTree_NoCache()
	}

	var newTree []Category
	err = json.Unmarshal([]byte(j), &newTree)
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum(err.Error()))
		return this.getTree_NoCache()
	}

	em.LogDebug.OutputSimplePath("Get CmsCategory-Tree from the cache according to UrlPath.")	// debug

	return newTree, nil
}

// Get Category By ID
// 根据ID获取Category
func (this *Category) GetById(categoryId uint) (Category, error) {
	list, err := this.GetAll()
	if err != nil {
		em.LogError.Output(em.MessageWithLineNum(err.Error()))
		return Category{}, err
	}
	for _, v := range list {
		if v.ID == categoryId {
			return v, nil
		}
	}
	return Category{}, errors.New("CmsCategory GetById Not Found!")
}

// Sort CmsCategory	[Go Sort package]
// 分类排序		[Go Sort包]
type category_SortDesc []Category
func (c category_SortDesc) Len() int           { return len(c) }
func (c category_SortDesc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c category_SortDesc) Less(i, j int) bool { return c[i].Sort > c[j].Sort }
func (this *Category) SortDesc(slice []Category) {
	sort.Sort(category_SortDesc(slice))
	for _, v := range slice {
		if len(v.Children) > 0 {
			this.SortDesc(v.Children)
		}
	}
	return
}
type category_SortAsc []Category
func (c category_SortAsc) Len() int           { return len(c) }
func (c category_SortAsc) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c category_SortAsc) Less(i, j int) bool { return c[i].Sort < c[j].Sort }
func (this *Category) SortAsc(slice []Category) {
	sort.Sort(category_SortAsc(slice))
	for _, v := range slice {
		if len(v.Children) > 0 {
			this.SortAsc(v.Children)
		}
	}
	return
}

// Make Category Tree
// 制作分类树
func (this *Category) MakeTree(m []Category)  []Category {
	var data = make(map[uint]map[uint]Category)
	for _, v := range m {
		id := v.ID
		parent_id := v.ParentID
		if _, ok := data[parent_id]; !ok {
			data[parent_id] = make(map[uint]Category)
		}
		data[parent_id][id] = v
	}

	var f func(index uint, d map[uint]map[uint]Category) []Category

	f = func(index uint, d map[uint]map[uint]Category) []Category {
		tmp := make([]Category, 0)
		for id, item := range d[index] {
			if d[id] != nil {
				item.Children = f(id, d)
			}
			tmp = append(tmp, item)
		}
		return tmp
	}

	tree := f(0, data)

	return tree
}

func (this *Category) WithAttachment(ctx *context.Context, c *[]Category, owner_type string) (error) {
	// 1.Get all ids
	var f func (v []Category, ids *[]uint32)
	f = func (v []Category, ids *[]uint32) {
		for _, sv := range v {
			*ids = append(*ids, uint32(sv.ID))
			if sv.Children != nil {
				f(sv.Children, ids)
			}
		}
	}
	var ids []uint32
	f(*c, &ids)

	// 2.Get all thumbnail
	b, err := client.NewClient().Attachment_GetMany(ctx, ids, application.Relationship_category_thumbnail)
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

	var f2 func(cat *[]Category)
	f2 = func(cat *[]Category) {
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
			// If have child
			if v.Children != nil {
				f2(&v.Children)
			}
		}
	}

	f2(c)



	return nil
}


func (this *Category) AttachmentSortAsc(slice []Category) {
	for _, v := range slice {
		sort.Sort(attachment_SortAsc(v.Thumbnail))
		if len(v.Children) > 0 {
			this.AttachmentSortAsc(v.Children)
		}
	}
	return
}