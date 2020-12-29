package application

const (
	Service_Attachment = "AttachmentRpcService"
)

const (
	Relationship_category_thumbnail = "category-thumbnail"
	Relationship_Post_Thumbnail = "post-thumbnail"
	Relationship_Post_Content = "post-content"
)


// Redis variable
// Redis变量
var (

	Cache_CmsVariableGetAll = "CmsVariableGetAll"

	/*
		Category
	*/
	// String
	// Delete trigger: Create a category, edit a category, delete a category
	// 删除触发： 创建一个分类、编辑一个分类、删除一个分类
	Cache_CmsCategoryGetTree = "CmsCategoryGetTree"
	Cache_CmsCategoryGetAll = "CmsCategoryGetAll"

	// Hash[Url_Path]
	// Delete trigger: Create the URLPATH trigger (prevent the same URL), trigger when the category of the UrlPath is edited, trigger when the category of the UrlPath is deleted
	// 删除触发： 创建该URLPATH触发（防止同URL），编辑该UrlPath的分类时触发，删除该UrlPath的分类时触发
	Cache_CmsCategoryGetByUrlPath = "CmsCategoryGetByUrlPath"

	/*
		[Post]
	*/
	/*
		Hash[Category ID]

		Function: The user saves a single page page, such as about us, contact us
		作用：用户储存单页面page，如关于我们，联系我们

		Delete trigger: Only delete the cache under this category ID when creating an article, only delete the cache under this category ID when editing an article, and only delete the cache under this category ID when deleting an article
		删除触发： 创建文章时仅删除该分类ID下的缓存，编辑文章时仅删除该分类ID下的缓存，删除文章时仅删除该分类ID下的缓存
	*/
	Cache_CmsGetOnePostByCategoryIdWithHighestSort = "GetOnePostByCategoryIdWithHighestSort"

	// Hash[Url Path]
	// Delete trigger: Create the URLPATH trigger (prevent the same URL), trigger when editing the UrlPath article, trigger when deleting the UrlPath article
	// 删除触发：创建该URLPATH触发（防止同URL），编辑该UrlPath的文章时触发，删除该UrlPath的文章时触发
	Cache_CmsPostGetByUrlPath = "CmsPostGetByUrlPath"
)