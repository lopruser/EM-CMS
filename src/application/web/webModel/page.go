package webModel

import (
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/web"
	em "github.com/Etpmls/Etpmls-Micro/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
	"time"
)

type page struct {

}

func NewPage() *page {
	return &page{}
}



// Determine whether the template exists
// 判断模板是否存在
func (this *page) CheckTemplateExist(t string) error {
	_, err := em.GetServiceNameKvKey(web.KvServiceTemplate + "/" + t)
	if err != nil {
		return err
	}

	return nil
}

// Get Categroy page Data
// 获取单页面分类栏目数据
type ApiPageGetCategroyPageV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
}
// Get Category page Data
func (this *page) GetCategroyPage(j ApiPageGetCategroyPageV2) (h gin.H, err error) {
	// 查询分类
	url_path := strings.TrimPrefix(j.UrlPath, "/category/")

	category, err := NewCategory().GetByUrlPath(url_path)
	if err != nil {
		return h, err
	}

	// Get First Post
	p, err := NewPost().GetOnePostByCategoryIdWithHighestSort(category.ID)
	if err != nil {
		em.LogError.OutputFullPath(err)
		return h, err
	}

	list, err := category.GetAll()
	if err != nil {
		em.LogError.OutputFullPath(err)
		return h, err
	}

	// Breadcrumbs
	breadcrumbs := this.MakeBreadcrumbs(category.ID, list)

	// Get Category Tree
	tree := category.MakeTree(list)
	category.SortDesc(tree)

	// Get Sidebar
	parentCategory := NewCategory().GetTopLevelParentById(category.ID, list)
	sidebar := NewCategory().MakeTreeById(parentCategory.ID, list)
	category.SortDesc(sidebar.Children)

	// Active Ids
	allParentCategory := NewCategory().GetAllParentWithSelfById(category.ID, list)
	var active_ids []uint
	for _, v := range allParentCategory {
		active_ids = append(active_ids, v.ID)
	}

	return gin.H{
		"category": category,
		"post": p,
		"breadcrumbs": breadcrumbs,
		"category_tree": tree,
		"sidebar": sidebar,
		"top_category": parentCategory,
		"active_category_ids": active_ids,
	}, err
}


// Get Categroy List
// 获取列表分类栏目数据
type ApiPageGetCategroyListV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
	CurrentPage int `json:"current_page" validate:"omitempty,min=1"`
	PerPage	int `json:"per_page" validate:"omitempty,min=0"`
}
func (this *page) PageGetCategroyListV2(j ApiPageGetCategroyListV2) (h gin.H, err error) {
	// 查询分类
	url_path := strings.TrimPrefix(j.UrlPath, "/category/")
	category, err := NewCategory().GetByUrlPath(url_path)
	if err != nil {
		return h, err
	}

	// Get all category
	// 获取全部分类
	list, err := category.GetAll()
	if err != nil {
		em.LogError.Output(err)
		return h, err
	}

	// Get all children id with self
	// 获取包括自身在内的全部子集
	ids := NewCategory().MakeAllChildrenIdByParentId(category.ID, list)
	ids = append(ids, category.ID)
	var post []model.Post
	var ct int64
	em.DB.Model(&model.Post{}).Where("category_id IN (?)", ids).Count(&ct)
	count := int(ct)

	// If offset number > count number, then dont use DB
	if !( count < (j.PerPage * (j.CurrentPage - 1)) ) {
		em.DB.Where("category_id IN (?)", ids).Order("sort desc").Order("created_at desc").Offset(j.PerPage * (j.CurrentPage - 1)).Limit(j.PerPage).Find(&post)
	}

	var p model.Post
	var common common
	ctx, err := common.GetPublicToken()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	} else {
		err = p.WithAttachment(&ctx, &post, application.Relationship_Post_Thumbnail)
		if err != nil {
			em.LogWarn.OutputSimplePath(err)
		}
		p.AttachmentSortAsc(post)
	}


	// Breadcrumbs
	breadcrumbs := this.MakeBreadcrumbs(category.ID, list)

	// 3.Make Category Tree
	tree := category.MakeTree(list)
	category.SortDesc(tree)

	// 4.Get Sidebar
	parentCategory := NewCategory().GetTopLevelParentById(category.ID, list)
	sidebar := NewCategory().MakeTreeById(parentCategory.ID, list)
	category.SortDesc(sidebar.Children)

	// 5. Active Ids
	allParentCategory := NewCategory().GetAllParentWithSelfById(category.ID, list)
	var active_ids []uint
	for _, v := range allParentCategory {
		active_ids = append(active_ids, v.ID)
	}

	return gin.H{
		"category": category,
		"post": post,
		"breadcrumbs": breadcrumbs,
		"category_tree": tree,
		"sidebar": sidebar,
		"top_category": parentCategory,
		"active_category_ids": active_ids,
		"post_count": count,
	}, err
}



// Get Post Data
// 获取文章数据
type ApiPageGetPostV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
}
func (this *page) PageGetPostV2(j ApiPageGetPostV2) (h gin.H, err error) {
	// Find Post By Url Path
	// 根据 URL PATH 查找文章
	url_path := strings.TrimPrefix(j.UrlPath, "/post/")

	post, err := NewPost().GetByUrlPath(url_path)
	if err != nil {
		return h, err
	}

	var p model.Post
	var common common
	ctx, err := common.GetPublicToken()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	} else {
		err = p.WithAttachmentByPost(&ctx, &post, application.Relationship_Post_Thumbnail)
		if err != nil {
			em.LogWarn.OutputSimplePath(err)
		}
		p.AttachmentSortAscByPost(&post)
	}

	// Find Category
	// 查找Category
	var category model.Category
	c, err := category.GetById(post.CategoryID)
	if err != nil {
		return h, err
	}

	// Get all category
	list, err := category.GetAll()
	if err != nil {
		em.LogError.Output(err)
		return h, err
	}

	// Breadcrumbs
	breadcrumbs := this.MakeBreadcrumbs(post.CategoryID, list)

	// Make Category Tree
	tree := category.MakeTree(list)
	category.SortDesc(tree)

	// Get Related products
	var related_post []model.Post
	em.DB.Where("category_id = ?", post.CategoryID).Not("id = ?", post.ID).Order("updated_at desc").Limit(20).Find(&related_post)

	ctx2, err := common.GetPublicToken()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	} else {
		err = p.WithAttachment(&ctx2, &related_post, application.Relationship_Post_Thumbnail)
		if err != nil {
			em.LogWarn.OutputSimplePath(err)
		}
		p.AttachmentSortAsc(related_post)
	}

	// Get Sidebar
	parentCategory := NewCategory().GetTopLevelParentById(category.ID, list)
	sidebar := NewCategory().MakeTreeById(parentCategory.ID, list)
	category.SortDesc(sidebar.Children)

	// Active Ids
	allParentCategory := NewCategory().GetAllParentWithSelfById(category.ID, list)
	var active_ids []uint
	for _, v := range allParentCategory {
		active_ids = append(active_ids, v.ID)
	}

	return gin.H{
		"category": c,
		"post": post,
		"breadcrumbs": breadcrumbs,
		"category_tree": tree,
		"related_post": related_post,
		"sidebar": sidebar,
		"top_category": parentCategory,
		"active_category_ids": active_ids,
	}, err
}


// Get Search Data
// 获取搜索页数据
type ApiPageGetSearchV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
	CurrentPage int `json:"current_page" validate:"omitempty,min=1"`
	PerPage	int `json:"per_page" validate:"omitempty,min=0"`
	Keyword string `json:"keyword"`
	SidebarCategoryId uint `json:"sidebar_category_id" validate:"omitempty,min=1"`
}
func (this *page) PageGetSearchV2(j ApiPageGetSearchV2) (h gin.H, err error) {
	var post []model.Post
	var count int64

	em.DB.Model(&model.Post{}).Where("name " + em.FUZZY_SEARCH + " ?", "%" + j.Keyword + "%").Count(&count)
	// If offset number > count number, then dont use DB
	if !( int(count) < (j.PerPage * (j.CurrentPage - 1)) ) {
		em.DB.Where("name " + em.FUZZY_SEARCH + " ?", "%" + j.Keyword + "%").Offset(j.PerPage * (j.CurrentPage - 1)).Limit(j.PerPage).Find(&post)
	}

	var p model.Post
	var common common
	ctx, err := common.GetPublicToken()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	} else {
		err = p.WithAttachment(&ctx, &post, application.Relationship_Post_Thumbnail)
		if err != nil {
			em.LogWarn.OutputSimplePath(err)
		}
		p.AttachmentSortAsc(post)
	}


	// Get all category
	var category model.Category
	list, err := category.GetAll()
	if err != nil {
		em.LogError.OutputFullPath(err)
		return h, err
	}


	// 1.Make Category Tree
	tree := category.MakeTree(list)
	category.SortDesc(tree)

	// 4.Get Sidebar
	sidebar := NewCategory().MakeTreeById(j.SidebarCategoryId, list)
	category.SortDesc(sidebar.Children)

	return gin.H{
		"post": post,
		"category_tree": tree,
		"sidebar": sidebar,
		"post_count": count,
		"keyword": j.Keyword,
	}, err
}


// Obtain the Category instance according to the Category ID (without the Children sub Category)
// 根据Category ID获取Category实例（不带Children子Category）
type ApiPageGetCategoryByCategoryIdV2 struct {
	ID        uint `uri:"id" binding:"required" validate:"min=1,max=9999999999"`
	CreatedAt time.Time `uri:"-"`
	UpdatedAt time.Time `uri:"-"`
	DeletedAt *time.Time `uri:"-"`
}
func (this *page) PageGetCategoryByCategoryIdV2(u ApiPageGetCategoryByCategoryIdV2) (c model.Category, err error) {
	var category model.Category
	c, err = category.GetById(u.ID)
	if err != nil {
		return model.Category{}, err
	}
	return c, err
}


// Obtain multiple Category instances based on Category IDs (without Children)
// 根据Category IDs获取多个Category实例（不带Children）
type ApiPageGetCategoryManyByCategoryIdsV2 struct {
	Ids []uint `json:"ids" binding:"required"`
}
func (this *page) PageGetCategoryManyByCategoryIdsV2(j ApiPageGetCategoryManyByCategoryIdsV2) (c []model.Category, err error) {
	var categories []model.Category
	if err := em.DB.Where("id IN (?)", j.Ids).Find(&categories).Error; err != nil {
		return c, err
	}
	return categories, err
}


// Get template path
// 获取模板路径
func (this *page)PageGetTemplatePathV2() []string {
	var list []string
	files,_ := ioutil.ReadDir("storage/view/template")

	// Search Files In Dir
	// 在目录中搜索文件
	var f func(path string, list *[]string)
	f = func(path string, list *[]string)  {
		files,_ := ioutil.ReadDir(path)
		for _, v := range files {
			if v.IsDir() {
				if v.Name() != "layout" {
					f(path + "/" + v.Name(), list)
				}
			} else {
				*list = append(*list, path + "/" + v.Name())
			}
		}
	}

	for _,v := range files {
		if v.IsDir() {
			if v.Name() != "layout" {
				f("storage/view/template/" + v.Name(), &list)
			}
		} else {
			if v.Name() != ".gitignore" {
				list = append(list, v.Name())
			}
		}
	}

	// Trim Prefix
	for k, v := range list {
		list[k] = strings.TrimPrefix(v, "storage/view/template/")
	}
	return list
}


// Make BreadCrumbs
// 制作面包屑
func (this *page) MakeBreadcrumbs(cId uint, list []model.Category) []model.Category {
	// 1.Get breadcrumbs
	breadcrumbs := NewCategory().GetAllParentWithSelfById(cId, list)
	// 2.Sort Breadcrumbs 反转切片，顺序从父级到子级
	var func_reverse func(s []model.Category) []model.Category
	func_reverse = func(s []model.Category) []model.Category {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	func_reverse(breadcrumbs)
	return breadcrumbs
}


// Get Category Tree by current page url
// 根据当前页面的Url获取分类树
type ApiPageGetCategoryTreeByUrlPathV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
}
func (this *page) PageGetCategoryTreeByUrlPathV2(j ApiPageGetCategoryTreeByUrlPathV2) (tree model.Category, err error) {
	// 查询分类
	url_path := strings.TrimPrefix(j.UrlPath, "/category/")
	var category model.Category
	c, err := NewCategory().GetByUrlPath(url_path)
	if err != nil {
		return tree, err
	}

	// Get all category
	list, err := category.GetAll()
	if err != nil {
		em.LogError.OutputFullPath(err)
		return tree, err
	}

	tree = NewCategory().MakeTreeById(c.ID, list)

	return tree, err
}


// 根据父级category_ID得到子集
type ApiCategoryGetChildrenV2 struct {
	ID        uint `uri:"id" binding:"required" validate:"min=1,max=9999999999"`
	CreatedAt time.Time `uri:"-"`
	UpdatedAt time.Time `uri:"-"`
	DeletedAt *time.Time `uri:"-"`
}
// Get Sub 1 Level Category children
func (this *page) CategoryGetChildrenV2(u ApiCategoryGetChildrenV2) interface{} {
	var category model.Category
	data, _ := NewCategory().GetChildrenByParentId(u.ID)

	category.SortDesc(data)
	return data
}


// Get many posts by many ids
type ApiPostGetManyByIdsV2 struct {
	Ids []uint `json:"ids" binding:"required"`
}
func (this *page) PostGetManyByIdsV2(j ApiPostGetManyByIdsV2) interface{} {
	var data []model.Post
	em.DB.Where(j.Ids).Find(&data)

	var p model.Post
	var common common
	ctx, err := common.GetPublicToken()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	} else {
		err = p.WithAttachment(&ctx, &data, application.Relationship_Post_Thumbnail)
		if err != nil {
			em.LogWarn.OutputSimplePath(err)
		}
		p.AttachmentSortAsc(data)
	}

	return data
}


type ApiPageGetPostByCategoryUrlPathV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	UrlPath string `json:"url_path" binding:"required" validate:"min=1,max=255"`
}
func (this *page) PageGetPostByCategoryUrlPathV2(j ApiPageGetPostByCategoryUrlPathV2) (post []model.Post, err error) {
	url_path := strings.TrimPrefix(j.UrlPath, "/category/")

	category, err := NewCategory().GetByUrlPath(url_path)
	if err != nil {
		em.LogError.OutputFullPath(err)
		return post, err
	}

	// Get all children id with self
	ids := NewCategory().CategoryGetAllChildrenIdByParentIdV2(category.ID)
	ids = append(ids, category.ID)

	em.DB.Where("category_id IN (?)", ids).Find(&post)

	return post, nil
}


type ApiPageGetPostByCategoryIdV2 struct {
	ID        uint `json:"id" binding:"required" validate:"min=1"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
func (this *page) PageGetPostByCategoryIdV2(j ApiPageGetPostByCategoryIdV2) (post []model.Post) {
	// Get all children id with self
	ids := NewCategory().CategoryGetAllChildrenIdByParentIdV2(j.ID)
	ids = append(ids, j.ID)

	em.DB.Where("category_id IN (?)", ids).Find(&post)

	return post
}



