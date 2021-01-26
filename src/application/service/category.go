package service

import (
	"context"
	"errors"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/client"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/Etpmls/Etpmls-Micro/define"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"github.com/google/uuid"
)

type ServiceCategory struct {
	protobuf.UnimplementedCategoryServer
}

type validate_CategoryCreate struct {
	Name string	`json:"name" validate:"required,min=1,max=255"`
	ParentID uint	`json:"parent_id" validate:"min=0,max=9999999999"`
	Kind string	`json:"kind" validate:"required,oneof=link webPage"`
	UrlPath string	`json:"url_path" validate:"omitempty,min=1,max=255"`
	Sort int	`json:"sort" validate:"min=0,max=9999999999"`
	Summary string	`json:"summary" validate:"max=1000"`
	TemplatePath string	`json:"template_path" validate:"max=255"`
	PostTemplatePath string	`json:"post_template_path" validate:"max=255"`
	Status int	`json:"status" validate:"numeric,max=99"`
	IsHidden int	`json:"is_hidden" validate:"numeric,max=99"`
	IsMain int	`json:"is_main" validate:"numeric,max=99"`
	Thumbnail []model.Attachment `json:"thumbnail" validate:"max=255"`
}
func (this *ServiceCategory) Create(ctx context.Context, request *protobuf.CategoryCreate) (response *em_protobuf.Response, err error) {
	// Validate
	{
		var vd validate_CategoryCreate
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	// Request -> Category
	var category model.Category
	form, err := category.InterfaceToCategory(request)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Create"), nil, err)
	}

	// If there is no url_path, use UUID make one new url_path
	// 如果没有URL路径为空，则用UUID生成一个新的URL路径
	if form.UrlPath == "" {
		form.UrlPath = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	err = em.DB.Transaction(func(tx *gorm.DB) error {
		// Automatically create and associate attachments
		// 自动创建及关联附件
		if err := tx.Omit(clause.Associations).Create(&form).Error; err != nil {
			return err
		}

		var paths []string
		for _, v := range form.Thumbnail {
			paths = append(paths, v.Path)
		}
		err := client.NewClient().Attachment_CreateMany(&ctx, paths, uint32(form.ID), application.Relationship_category_thumbnail)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}


	// If caching is enabled
	// 如果开启了缓存功能
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		em.Cache.DeleteHash(application.Cache_CmsCategoryGetByUrlPath, form.UrlPath)
		em.Cache.DeleteString(application.Cache_CmsCategoryGetTree)
		em.Cache.DeleteString(application.Cache_CmsCategoryGetAll)
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Create"), nil)
}

type validate_CategoryEdit struct {
	validate_CategoryCreate
	ID        uint `json:"id" validate:"min=0,max=9999999999"`
}
func (this *ServiceCategory) Edit(ctx context.Context, request *protobuf.CategoryEdit) (response *em_protobuf.Response, err error) {
	// Validate
	{
		var vd validate_CategoryEdit
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	// Get the URLpath before modification
	// 获取修改前的URLpath
	var tmpCategory model.Category
	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}
	if strings.ToLower(e) == "true" {
		tmpCategory, err = tmpCategory.GetById(uint(request.GetId()))
		if err != nil {
			return em.ErrorRpc(codes.Internal, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Edit"), nil, err)
		}
	}

	// Request -> Category
	var category model.Category
	form, err := category.InterfaceToCategory(request)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Edit"), nil, err)
	}

	// If there is no url_path, use UUID make one new url_path
	// 如果没有URL路径为空，则用UUID生成一个新的URL路径
	if form.UrlPath == "" {
		form.UrlPath = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	err = em.DB.Transaction(func(tx *gorm.DB) error {
		var paths []string
		for _, v := range form.Thumbnail {
			paths = append(paths, v.Path)
		}
		err := client.NewClient().Attachment_CreateMany(&ctx, paths, uint32(form.ID), application.Relationship_category_thumbnail)
		if err != nil {
			return err
		}

		if err := tx.Omit(clause.Associations).Save(&form).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}

	// If caching is enabled
	// 如果开启了缓存功能
	if strings.ToLower(e) == "true" {
		em.Cache.DeleteHash(application.Cache_CmsCategoryGetByUrlPath, tmpCategory.UrlPath)	// URLPATH before modification 修改前的URLPATH
		em.Cache.DeleteHash(application.Cache_CmsCategoryGetByUrlPath, form.UrlPath)	// Modified URLPATH 修改的URLPATH
		em.Cache.DeleteString(application.Cache_CmsCategoryGetTree)
		em.Cache.DeleteString(application.Cache_CmsCategoryGetAll)
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Edit"), nil)
}

type validate_CategoryDelete struct {
	Categories []model.Category `json:"categories" validate:"required,min=1"`
}
func (this *ServiceCategory) Delete(ctx context.Context, request *protobuf.CategoryDelete) (response *em_protobuf.Response, err error) {
	// Validate
	{
		var vd validate_CategoryDelete
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	e, err := em.Kv.ReadKey(define.KvCacheEnable)
	if err != nil {
		em.LogDebug.OutputSimplePath(err)
	}

	var ids []uint32
	for _, v := range request.Categories {
		ids = append(ids, v.Id)
	}

	// 删除文章
	result := em.DB.Where("category_id IN ?", ids).First(&model.Post{})
	if result.RowsAffected > 0 {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_MESSAGE_PostExist"), nil, err)
	}

	var tmpc []model.Category

	err = em.DB.Transaction(func(tx *gorm.DB) error {
		// 检查分类是否为父分类，如果为父分类，需要先删除子分类才可以继续删除父分类
		result := tx.Where("parent_id IN ?", ids).First(&model.Category{})
		if result.RowsAffected > 0 {
			return errors.New(em.I18n.TranslateFromRequest(ctx, "ERROR_MESSAGE_SubCategoriesExist"))
		}

		// Delete Attachments
		// 删除附件
		err := client.NewClient().Attachment_Delete(&ctx, ids, application.Relationship_category_thumbnail)
		if err != nil {
			return err
		}

		// 删除分类
		var c []model.Category
		tx.Where("id IN ?", ids).Find(&c)

		// If caching is enabled
		// 如果开启了缓存功能
		if strings.ToLower(e) == "true" {
			tmpc = make([]model.Category, len(c))
			copy(tmpc, c)
		}

		result = tx.Delete(&c)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, err.Error(), nil, err)
	}


	// If caching is enabled
	// 如果开启了缓存功能
	if strings.ToLower(e) == "true" {
		// Delete Url Path in the list to be deleted
		// 删除待删除列表中的Url Path
		var tmpurl []string
		for _, v := range tmpc {
			tmpurl = append(tmpurl, v.UrlPath)
		}
		em.Cache.DeleteHash(application.Cache_CmsCategoryGetByUrlPath, tmpurl...)
		em.Cache.DeleteString(application.Cache_CmsCategoryGetTree)
		em.Cache.DeleteString(application.Cache_CmsCategoryGetAll)
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Delete"), nil)
}

func (this *ServiceCategory) GetTree(ctx context.Context, request *em_protobuf.Empty) (response *em_protobuf.Response, err error) {
	var category model.Category
	tree, err := category.GetTree()
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Get"), nil, err)
	}

	err = category.WithAttachment(&ctx, &tree, application.Relationship_category_thumbnail)
	if err != nil {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Get"), nil, err)
	}

	category.AttachmentSortAsc(tree)

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Delete"), tree)
}