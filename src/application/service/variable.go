package service

import (
	"context"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type ServiceVariable struct {
	protobuf.UnimplementedVariableServer
}

type validate_VariableCreate struct {
	Variable []struct{
		Name string	`json:"name" validate:"omitempty,max=255"`
		Value string	`json:"value" validate:"omitempty,max=5000"`
		Remark string `json:"remark" validate:"omitempty,max=1000"`
	}
}
func (this *ServiceVariable) Create(ctx context.Context, request *protobuf.VariableCreate) (*em_protobuf.Response, error) {
	// Validate
	{
		var vd validate_VariableCreate
		err := em.Validator.Validate(request, &vd)
		if err != nil {
			em.LogWarn.Output(em.MessageWithLineNum(err.Error()))
			return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_Validate"), nil, err)
		}
	}

	// If the user submits empty data, return directly
	// 如果用户提交空数据，直接返回
	if len(request.Variable) == 0 {
		return em.ErrorRpc(codes.InvalidArgument, em.ERROR_Code, em.I18n.TranslateFromRequest(ctx, "ERROR_MESSAGE_DataIsEmpty"), nil, nil)
	}

	err := em.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("1 = 1").Delete(model.Variable{}).Error; err != nil {
			return err
		}
		var vars []model.Variable
		for _, v := range request.Variable {
			var tmp model.Variable
			err := em.ChangeTypeV2(v, &tmp)
			if err != nil {
				return err
			}

			vars = append(vars, tmp)
		}
		result := tx.Create(&vars)
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
	if em.Micro.Config.App.EnableCache {
		// Delete cache
		// 删除缓存
		em.Cache.DeleteString(application.Cache_CmsVariableGetAll)
	}

	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Create"), nil)
}

func (this *ServiceVariable) GetAll(ctx context.Context, request *em_protobuf.Empty) (*em_protobuf.Response, error) {
	var v model.Variable
	data := v.GetAll()
	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Get"), data)
}