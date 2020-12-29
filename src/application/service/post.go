package service

import (
	"context"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/client"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_library "github.com/Etpmls/Etpmls-Micro/library"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
)

type ServicePost struct {
	protobuf.UnimplementedPostServer
}

type validate_PostCreate struct {
	Name string	`json:"name" validate:"required,min=1,max=255"`
	CategoryID uint	`json:"category_id" validate:"required,min=1,max=9999999999"`
	Content string	`json:"content" validate:"max=16777215"`
	TemplatePath string	`json:"template_path" validate:"max=255"`
	UrlPath string	`json:"url_path" validate:"max=255"`
	Language string	`json:"language" validate:"max=255"`
	Summary string	`json:"summary" validate:"max=1000"`
	Sort int	`json:"sort" validate:"min=0,max=9999999999"`
	Parameter string `json:"parameter"`
	Status int	`json:"status" validate:"numeric,max=99"`
	// tmp field
	Thumbnail  []model.Attachment    `gorm:"polymorphic:Owner;polymorphicValue:post-thumbnail" json:"thumbnail"`
	ContentImage  []model.Attachment `gorm:"polymorphic:Owner;polymorphicValue:post-content" json:"content_image"`
}
func (this *ServicePost) Create(ctx context.Context, request *protobuf.PostCreate) (*em_protobuf.Response, error) {
	// Validate
	{
		var vd validate_PostCreate
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	// Request -> Post
	var post model.Post
	form, err := post.InterfaceToPost(request)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Create"), nil, err)
	}

	err = em.DB.Transaction(func(tx *gorm.DB) error {

		// 1.Create Post 创建文章

		// - If there is no url_path, use UUID make one new url_path
		if form.UrlPath == "" {
			form.UrlPath = strings.ReplaceAll(uuid.New().String(), "-", "")
		}

		form.Status = 1

		if err := tx.Omit(clause.Associations).Create(&form).Error; err != nil {
			return err
		}

		// 2. Create Attachment 创建附件
		// -	[1] Create Thumbnail 创建缩略图
		var thumbnail_path []string
		for _, v := range request.Thumbnail {
			thumbnail_path = append(thumbnail_path, v.Path)
		}
		var t = client.Parameter_Post_CreateImage{
			Paths:     thumbnail_path,
			OwnerId:   uint32(form.ID),
			OwnerType: application.Relationship_Post_Thumbnail,
		}

		var content_path []string
		for _, v := range request.ContentImage {
			content_path = append(content_path, v.Path)
		}
		var c = client.Parameter_Post_CreateImage{
			Paths:     content_path,
			OwnerId:   uint32(form.ID),
			OwnerType: application.Relationship_Post_Content,
		}

		err := client.NewClient().Post_CreateImage(ctx, t, c)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}

	// If caching is enabled
	// 如果开启了缓存的功能
	if em.Micro.Config.App.EnableCache {
		em.Cache.DeleteHash(application.Cache_CmsPostGetByUrlPath, form.UrlPath)
		em.Cache.DeleteHash(application.Cache_CmsGetOnePostByCategoryIdWithHighestSort, strconv.Itoa(int(form.ID)))
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Create"), nil)
}

type validate_PostEdit struct {
	validate_PostCreate
	ID        uint `json:"id" validate:"required,min=1"`
}
func (this *ServicePost) Edit(ctx context.Context, request *protobuf.PostEdit) (*em_protobuf.Response, error) {
	// Validate
	{
		var vd validate_PostEdit
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	// Request -> Post
	var post model.Post
	form, err := post.InterfaceToPost(request)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Create"), nil, err)
	}

	// If caching is enabled
	// 如果开启了缓存的功能
	var tmpPost model.Post
	if em.Micro.Config.App.EnableCache {
		em.DB.First(&tmpPost, request.GetId())
	}


	// If there is no url_path, use UUID make one new url_path
	if form.UrlPath == "" {
		form.UrlPath = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	form.Status = 1

	err = em.DB.Transaction(func(tx *gorm.DB) error {

		// Create Attachment 创建附件
		// -	If there is a thumbnail 如果有传缩略图
		var thumbnail_path []string
		for _, v := range request.Thumbnail {
			thumbnail_path = append(thumbnail_path, v.Path)
		}
		var t = client.Parameter_Post_CreateImage{
			Paths:     thumbnail_path,
			OwnerId:   uint32(form.ID),
			OwnerType: application.Relationship_Post_Thumbnail,
		}

		var content_path []string
		for _, v := range request.ContentImage {
			content_path = append(content_path, v.Path)
		}
		var c = client.Parameter_Post_CreateImage{
			Paths:     content_path,
			OwnerId:   uint32(form.ID),
			OwnerType: application.Relationship_Post_Content,
		}

		err := client.NewClient().Post_CreateImage(ctx, t, c)
		if err != nil {
			return err
		}

		//  Edit Post 修改文章
		if err := tx.Save(&form).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}

	// If caching is enabled
	// 如果开启了缓存的功能
	if em.Micro.Config.App.EnableCache {
		em.Cache.DeleteHash(application.Cache_CmsPostGetByUrlPath, tmpPost.UrlPath)
		em.Cache.DeleteHash(application.Cache_CmsPostGetByUrlPath, form.UrlPath)
		em.Cache.DeleteHash(application.Cache_CmsGetOnePostByCategoryIdWithHighestSort, strconv.Itoa(int(request.GetId())))
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Edit"), nil)
}

type validate_PostDelete struct {
	Posts []model.Post `json:"posts" validate:"required,min=1"`
}
func (this *ServicePost) Delete(ctx context.Context, request *protobuf.PostDelete) (*em_protobuf.Response, error){
	// Validate
	{
		var vd validate_PostDelete
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	var ids []uint32
	for _, v := range request.Posts {
		ids = append(ids, v.Id)
	}

	// If caching is enabled
	// 如果开启了缓存的功能
	if em.Micro.Config.App.EnableCache {
		var ctmp []model.Post
		em.DB.Where(ids).Find(&ctmp)
		var tmpurl []string
		for _, v := range ctmp {
			tmpurl = append(tmpurl, v.UrlPath)
		}
		// []uint to []string
		var str []string
		for _, v := range ids {
			str = append(str, strconv.Itoa(int(v)))
		}
		// Delete cache
		// 删除缓存
		em.Cache.DeleteHash(application.Cache_CmsPostGetByUrlPath, tmpurl...)
		em.Cache.DeleteHash(application.Cache_CmsGetOnePostByCategoryIdWithHighestSort, str...)
	}

	err := em.DB.Transaction(func(tx *gorm.DB) error {

		err := client.NewClient().Attachment_Delete(&ctx, ids, application.Relationship_Post_Thumbnail)
		if err != nil {
			return err
		}
		err = client.NewClient().Attachment_Delete(&ctx, ids, application.Relationship_Post_Content)
		if err != nil {
			return err
		}

		// 删除分类
		result := tx.Where("id IN ?", ids).Delete(&model.Post{})
		if result.Error != nil {
			em.LogError.OutputSimplePath(result.Error.Error())
			return result.Error
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Delete"), nil)
}

func (this *ServicePost) GetAll(ctx context.Context, request *em_protobuf.Pagination) (*em_protobuf.Response, error){

	var data []model.Post

	var orm em_library.Gorm
	limit, offset := orm.GeneratePaginationLimit(int(request.Number), int(request.Size))

	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := request.Search

	em.DB.Model(&model.Post{}).Preload("Category").Where("name " + em.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Order("sort desc").Order("updated_at desc").Find(&data)

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Get"), map[string]interface{}{"data": data, "count":count})
}