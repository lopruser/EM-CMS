package webService

import (
	"encoding/json"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/web"
	"github.com/Etpmls/EM-CMS/src/application/web/webModel"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

type ServicePage struct {
	
}


// Show Index
// 显示首页
func (this *ServicePage)PageShowIndex(c *gin.Context)  {
	var v model.Variable
	data, _ := v.GetMap()

	c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/index.html", data)
	return
}

// Show category By category template_path
// 根据分类template_path来显示分类
func (this *ServicePage)PageShowCategory(c *gin.Context)  {
	url_path := c.Param("url_path")
	page := c.Query("ServicePage")
	cpage, err := strconv.Atoi(page)
	if err != nil {
		cpage = 1
	}
	current_page := math.Abs(float64(cpage))

	// 查询分类，如果不存在返回404
	category, err := webModel.NewCategory().GetByUrlPath(url_path)
	if err != nil {
		c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/404.html", nil)
		return
	}

	// If the template does not exist, return 404
	// 如果模板不存在返回404
	err = webModel.NewPage().CheckTemplateExist(category.TemplatePath)
	if err != nil {
		c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/404.html", nil)
		return
	}

	// Get Variable
	var v model.Variable
	data, _ := v.GetMap()

	data["current_page"] = int(current_page)

	c.HTML(http.StatusOK, category.TemplatePath, data)
	return
}

// Show Post
// 显示文章
func (this *ServicePage)PageShowPost(c *gin.Context)  {
	url_path := c.Param("url_path")

	// Query the article, if it does not exist, return 404
	// 查询文章，如果不存在返回404
	post, err := webModel.NewPost().GetByUrlPath(url_path)
	if err != nil {
		c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/404.html", nil)
		return
	}

	// If the article template does not exist, then query the category template
	// 如果文章模板不存在 再查询分类模板
	var tpl string
	err = webModel.NewPage().CheckTemplateExist(post.TemplatePath)
	if err != nil {
		var category model.Category
		// Get All CmsCategory
		// 获取全部CmsCategory
		list, err := category.GetAll()
		if err != nil {
			em.LogError.OutputFullPath(err)
			c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/404.html", nil)
			return
		}

		// Search Template by Post
		// 搜索文章模板
		searchTpl := webModel.NewCategory().SearchPostTemplate(post.CategoryID, list)

		// Determine whether the template exists
		// 判断模板是否存在
		err = webModel.NewPage().CheckTemplateExist(searchTpl)
		if err != nil {
			c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/404.html", nil)
			return
		} else {
			tpl = searchTpl
		}

	} else {
		tpl = post.TemplatePath
	}

	// Get Variable
	var v model.Variable
	data, _ := v.GetMap()

	c.HTML(http.StatusOK, tpl, data)
	return
}

// Show Search
// 显示搜索
func (this *ServicePage)PageShowSearch(c *gin.Context)  {
	keyword := c.Query("keyword")
	page := c.Query("ServicePage")
	cpage, err := strconv.Atoi(page)
	if err != nil {
		cpage = 1
	}
	current_page := math.Abs(float64(cpage))

	// Get Variable
	var v model.Variable
	data, _ := v.GetMap()

	data["current_page"] = int(current_page)
	data["keyword"] = keyword

	c.HTML(http.StatusOK, em.MustGetServiceNameKvKey(web.KvServiceDefaultTemplate)+ "/search.html", data)
	return
}

// Get Index ServicePage Data
// 获取首页数据
func (this *ServicePage)PageGetIndex(c *gin.Context)  {
	// 1.Get all category
	var category model.Category
	list, err := category.GetAll()
	if err != nil {
		em.LogError.OutputSimplePath(err)
	}

	// Get Category Tree
	tree := category.MakeTree(list)
	category.SortDesc(tree)

	// []Category -> []ApiCategoryGetAllV2
	jsonBytes, _ := json.Marshal(tree)

	var newTree []model.Category
	_ = json.Unmarshal([]byte(string(jsonBytes)), &newTree)

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", gin.H{
		"category_tree": newTree,
	})
	return
}


// Get Categroy ServicePage Data
// 获取单页面分类栏目数据
func (this *ServicePage)PageGetCategroyPage(c *gin.Context)  {
	var j webModel.ApiPageGetCategroyPageV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "success", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data, err := webModel.NewPage().GetCategroyPage(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get Categroy List
// 获取列表分类栏目数据
func (this *ServicePage)PageGetCategroyList(c *gin.Context)  {
	var j webModel.ApiPageGetCategroyListV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data, err := webModel.NewPage().PageGetCategroyListV2(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get Post Data
// 获取文章数据
func (this *ServicePage)PageGetPost(c *gin.Context)  {
	var j webModel.ApiPageGetPostV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}


	data, err := webModel.NewPage().PageGetPostV2(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get Search Data
// 获取搜索页数据
func (this *ServicePage)PageGetSearch(c *gin.Context)  {
	var j webModel.ApiPageGetSearchV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}


	data, err := webModel.NewPage().PageGetSearchV2(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Obtain the Category instance according to the Category ID (without the Children sub Category)
// 根据Category ID获取Category实例（不带Children子Category）
func (this *ServicePage)PageGetCategoryByCategoryId(c *gin.Context)  {
	var u webModel.ApiPageGetCategoryByCategoryIdV2
	if err := c.ShouldBindUri(&u); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(u)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data, err := webModel.NewPage().PageGetCategoryByCategoryIdV2(u)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Obtain multiple Category instances based on Category IDs (without Children)
// 根据Category IDs获取多个Category实例（不带Children）
func (this *ServicePage)PageGetCategoryManyByCategoryIds(c *gin.Context)  {
	var j webModel.ApiPageGetCategoryManyByCategoryIdsV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data, err := webModel.NewPage().PageGetCategoryManyByCategoryIdsV2(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get All Category Tree
// 得到分类树
func (this *ServicePage)PageGetCategoryTree(c *gin.Context)  {
	var category model.Category
	tree, _ := category.GetTree()
	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", tree)
	return
}


// Get Category Tree by current ServicePage url
// 根据当前页面的Url获取分类树
func (this *ServicePage)PageGetCategoryTreeByUrlPath(c *gin.Context)  {
	var j webModel.ApiPageGetCategoryTreeByUrlPathV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	tree, err := webModel.NewPage().PageGetCategoryTreeByUrlPathV2(j)
	if err != nil {
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", tree)
	return
}


// 根据父级category_ID得到子集
func (this *ServicePage)PageGetCategoryChildren(c *gin.Context)  {
	var u webModel.ApiCategoryGetChildrenV2
	if err := c.ShouldBindUri(&u); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(u)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data := webModel.NewPage().CategoryGetChildrenV2(u)
	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// 根据多个文章ID获取多个文章
func (this *ServicePage)PageGetPostManyByIds(c *gin.Context)  {
	var j webModel.ApiPostGetManyByIdsV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data := webModel.NewPage().PostGetManyByIdsV2(j)
	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get Post by Categroy UrlPath and post of sub category
// 根据分类UrlPath得到对应文章，以及子栏目文章
func (this *ServicePage)PageGetPostByCategoryUrlPath(c *gin.Context)  {
	var j webModel.ApiPageGetPostByCategoryUrlPathV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data, _ := webModel.NewPage().PageGetPostByCategoryUrlPathV2(j)

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}


// Get Post by Categroy ID
// 根据分类ID得到对应文章
func (this *ServicePage)PageGetPostByCategoryId(c *gin.Context)  {
	var j webModel.ApiPageGetPostByCategoryIdV2
	if err := c.ShouldBindJSON(&j); err != nil{
		web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
		return
	}

	// Validate
	{
		err := em.Validator.ValidateStruct(j)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			web.JsonError(c, http.StatusBadRequest, web.WebErrorCode, "error", nil, err)
			return
		}
	}

	data := webModel.NewPage().PageGetPostByCategoryIdV2(j)

	web.JsonSuccess(c, http.StatusOK, web.WebSuccessCode, "success", data)
	return
}
