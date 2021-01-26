package service

import (
	"context"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
	"google.golang.org/grpc/codes"
)

type ServiceSetting struct {
	protobuf.UnimplementedSettingServer
}



func (this *ServiceSetting) ReloadTemplate(ctx context.Context, request *em_protobuf.Empty) (response *em_protobuf.Response, err error) {
	var s model.Setting
	err = s.ReloadTemplate(application.HTMLRender)
	if err != nil {
		return em.ErrorRpc(codes.Internal, em.ERROR_Code, err.Error(), nil, err)
	}
	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Operation"), nil)
}