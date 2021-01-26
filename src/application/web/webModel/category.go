package webModel

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/Etpmls/Etpmls-Micro/define"
	"github.com/go-redis/redis/v8"
	"strings"
)

type category struct {

}

func NewCategory() *category {
	return &category{}
}

// Get category according to url_path
// 根据url_path获取category
func (this *category) GetByUrlPath(url_path string) (model.Category, error) {
	s, _ := em.Kv.ReadKey(define.KvCacheEnable)
	if strings.ToLower(s) == "true" {
		return this.getByUrlPath_Cache(url_path, strings.ToLower(s) == "true")
	} else {
		return this.getByUrlPath_NoCache(url_path, strings.ToLower(s) == "true")
	}
}
func (this *category) getByUrlPath_Cache(url_path string, enableCache bool) (model.Category, error) {
	s, err := em.Cache.GetHash(application.Cache_CmsCategoryGetByUrlPath, url_path)
	if err != nil {
		if err == redis.Nil {
			return this.getByUrlPath_NoCache(url_path, enableCache)
		}
		em.LogError.OutputFullPath(err)
		return model.Category{}, err
	}

	var c model.Category
	err = json.Unmarshal([]byte(s), &c)
	if err != nil {
		em.LogError.OutputFullPath(err)
		return model.Category{}, err
	}

	return c, nil
}
func (this *category) getByUrlPath_NoCache(url_path string, enableCache bool) (model.Category, error) {
	var c model.Category
	result := em.DB.Where("url_path = ?", url_path).Order("sort desc").First(&c)
	if result.Error != nil {
		em.LogError.OutputFullPath(result.Error)
		return model.Category{}, result.Error
	}

	// If the cache function is turned on
	// 如果开启缓存功能
	if enableCache {
		b, err := json.Marshal(c)
		if err != nil  {
			em.LogError.OutputFullPath(err)
			return model.Category{}, err
		}
		var m = make(map[string]string)
		m[url_path] = string(b)
		em.Cache.SetHash(application.Cache_CmsCategoryGetByUrlPath, m)
	}

	return c, nil
}


// Get Children By Parent Id
// 根据父级ID获取子集
func (this *category) GetChildrenByParentId(parentId uint) ([]model.Category, error) {
	var c model.Category
	list, err := c.GetAll()
	if err != nil {
		return nil, err
	}
	var tmp []model.Category
	for _, v := range list {
		if v.ParentID == parentId {
			tmp = append(tmp, v)
		}
	}

	return tmp, nil
}

// Query category post-template (if no template is set for the current category, search for the parent category)
// 查询分类文章模板（若当前分类未设置模板，则寻找父级分类）
func (this *category) SearchPostTemplate(categoryId uint, list []model.Category) string {
	if categoryId == 0 {
		return ""
	}

	for _, v := range list {
		// 如果找到了分类，则获取模板
		if v.ID == categoryId {
			if len(v.PostTemplatePath) > 0 {
				return v.PostTemplatePath
			}

			// If the template path is empty and the parent ID is not 0, continue to find the parent template
			// 如果模板路径为空，且父级ID不为0，则继续查找父级模板
			if len(v.PostTemplatePath) == 0 && v.ParentID != 0 {
				return this.SearchPostTemplate(v.ParentID, list)
			} else {
				return ""
			}
		}
	}

	return ""
}



// Make Category Tree
// 制作分类树
func (this *category) MakeTree(m []model.Category)  []model.Category {
	var data = make(map[uint]map[uint]model.Category)
	for _, v := range m {
		id := v.ID
		parent_id := v.ParentID
		if _, ok := data[parent_id]; !ok {
			data[parent_id] = make(map[uint]model.Category)
		}
		data[parent_id][id] = v
	}

	var f func(index uint, d map[uint]map[uint]model.Category) []model.Category

	f = func(index uint, d map[uint]map[uint]model.Category) []model.Category {
		tmp := make([]model.Category, 0)
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


func (this *category) MakeTreeById(categoryId uint, category []model.Category) model.Category {
	var thisCategory model.Category
	for _, v := range category {
		if v.ID == categoryId {
			thisCategory = v
			var func_SearchChildren func(parentCategory *model.Category, allCategory []model.Category)
			func_SearchChildren = func(parentCategory *model.Category, allCategory []model.Category){
				for _, c := range allCategory {
					if c.ParentID == parentCategory.ID {
						func_SearchChildren(&c, allCategory)
						parentCategory.Children = append(parentCategory.Children, c)
					}
				}
			}
			func_SearchChildren(&thisCategory, category)
			break
		}
	}
	return thisCategory
}


// Get all parent category with level sort with self
//得到带等级排序的所有父级分类（包括自身）
func (this *category) GetAllParentWithSelfById(u uint, all_category []model.Category) []model.Category {
	var tmp_category []model.Category
	for _, v := range all_category {
		if v.ID == u {
			tmp_category = append(tmp_category, v)
			if v.ParentID != 0 {
				var method func(currentId uint, allCategory []model.Category) []model.Category
				method = func(currentId uint, allCategory []model.Category) []model.Category {
					var t []model.Category
					for _, v := range allCategory {
						if currentId == v.ID {
							t = append(t, v)
							if v.ParentID != 0 {
								t = append(t, method(v.ParentID, allCategory)...)
							}
						}
					}
					return t
				}
				tmp_category = append(tmp_category, method(v.ParentID, all_category)...)
			}
			break
		}
	}
	return tmp_category
}


// Get top level parent
func (this *category) GetTopLevelParentById(u uint, all_category []model.Category) model.Category {
	var tmp model.Category
	for _, v := range all_category {
		if v.ID == u {
			if v.ParentID == 0 {
				return v
			} else {
				var func_SearchParent func(u uint, allCategory []model.Category) model.Category
				func_SearchParent = func(u uint, allCategory []model.Category) model.Category {
					for _, sv := range allCategory {
						if sv.ID == u {
							if sv.ParentID != 0 {
								tmp = func_SearchParent(sv.ParentID, allCategory)
							} else {
								tmp = sv
							}
							break
						}
					}
					return tmp
				}
				func_SearchParent(v.ParentID, all_category)
			}
			break
		}
	}
	return tmp
}


// Make all Category Children ID by parent category id
func (this *category) MakeAllChildrenIdByParentId(u uint, c []model.Category) []uint {
	for _, v := range c {
		if v.ID == u {

			// 匿名函数递归
			var method func(u uint, c []model.Category) []uint
			method = func(u uint, c []model.Category) []uint {
				var tmp []uint
				for _, v := range c {
					if u == v.ParentID {
						tmp = append(tmp, v.ID)
						tmp = append(tmp, method(v.ID, c)...)
					}
				}
				return tmp
			}

			child := method(v.ID, c)
			return child
		}
	}

	return nil
}


// Get All Category Tree
// 得到分类树
func (this *category) CategoryGetTreeV2() (interface{}, error) {
	var c model.Category
	data, err := c.GetAll()
	if err != nil {
		em.LogError.OutputFullPath(err)
		return nil, err
	}
	tree := this.MakeTree(data)
	c.SortDesc(tree)

	return tree, nil
}


/* Get all Category Children ID by parent category id*/
func (this *category) CategoryGetAllChildrenIdByParentIdV2(u uint) ([]uint) {
	var cat model.Category
	c, err := cat.GetAll()
	if err != nil {
		em.LogError.OutputFullPath(err)
		return nil
	}

	for _, v := range c {
		if v.ID == u {

			// 匿名函数递归
			var method func(u uint, c []model.Category) []uint
			method = func(u uint, c []model.Category) []uint {
				var tmp []uint
				for _, v := range c {
					if u == v.ParentID {
						tmp = append(tmp, v.ID)
						tmp = append(tmp, method(v.ID, c)...)
					}
				}
				return tmp
			}

			child := method(v.ID, c)
			return child
		}
	}

	return nil
}
