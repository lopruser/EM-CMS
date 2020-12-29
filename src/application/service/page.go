package service

import (
	"context"
	"github.com/Etpmls/EM-CMS/src/application/model"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
)

type ServicePage struct {
	protobuf.UnimplementedPageServer
}

func (this *ServicePage) GetTemplatePath(ctx context.Context, request *em_protobuf.Empty) (*em_protobuf.Response, error) {
	var p model.Page
	data := p.PageGetTemplatePath()
	return em.SuccessRpc(em.SUCCESS_Code, em.I18n.TranslateFromRequest(ctx, "SUCCESS_Get"), data)
}