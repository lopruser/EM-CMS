package client

import (
	"context"
	"errors"
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	em "github.com/Etpmls/Etpmls-Micro"
	em_protobuf "github.com/Etpmls/Etpmls-Micro/protobuf"
)

func (this *client) Attachment_Delete(ctx *context.Context, owner_ids []uint32, owner_type string) error {
	cl := em.Micro.Client.NewClient()
	err := cl.ConnectServiceWithToken(application.Service_Attachment, ctx)
	if err != nil {
		em.LogWarn.OutputSimplePath(err)
		return err
	}

	c := protobuf.NewAttachmentClient(cl.Conn)
	return cl.Sync_Simple(func() (response *em_protobuf.Response, e error) {
		return c.Delete(*ctx, &protobuf.AttachmentDelete{
			Service:       em.Micro.Config.ServiceDiscovery.Service.Rpc.Name,
			OwnerIds:       owner_ids,
			OwnerType:     owner_type,
		})
	},nil)
}

func (this *client) Attachment_CreateMany(ctx *context.Context, paths []string, owner_id uint32, owner_type string) error {
	cl := em.Micro.Client.NewClient()
	err := cl.ConnectServiceWithToken(application.Service_Attachment, ctx)
	if err != nil {
		em.LogWarn.OutputSimplePath(err)
		return err
	}

	c := protobuf.NewAttachmentClient(cl.Conn)
	return cl.Sync_Simple(func() (response *em_protobuf.Response, e error) {
		return c.CreateMany(*ctx, &protobuf.AttachmentCreateMany{
			Service:       em.Micro.Config.ServiceDiscovery.Service.Rpc.Name,
			Paths:         paths,
			OwnerId:       owner_id,
			OwnerType:     owner_type,
		})
	},nil)
}

func (this *client) Attachment_GetMany(ctx *context.Context, owner_ids []uint32, owner_type string) ([]byte, error) {
	cl := em.Micro.Client.NewClient()
	err := cl.ConnectServiceWithToken(application.Service_Attachment, ctx)
	if err != nil {
		em.LogWarn.OutputSimplePath(err)
		return nil, err
	}

	c := protobuf.NewAttachmentClient(cl.Conn)
	return cl.Sync_SimpleV2(func() (response *em_protobuf.Response, e error) {
		return c.GetMany(*ctx, &protobuf.AttachmentGetMany{
			Service:       em.Micro.Config.ServiceDiscovery.Service.Rpc.Name,
			OwnerIds:       owner_ids,
			OwnerType:     owner_type,
		})
	},nil)
}

type Parameter_Post_CreateImage struct {
	Paths     []string
	OwnerId   uint32
	OwnerType string
}
func (this *client) Post_CreateImage(ctx context.Context, thumbnail Parameter_Post_CreateImage, content Parameter_Post_CreateImage) error {
	// 1.Connect Service
	cl := em.Micro.Client.NewClient()
	err := cl.ConnectService(em.Micro.Config.ServiceDiscovery.Service.Prefix + application.Service_Attachment)
	if err != nil {
		return err
	}
	c := protobuf.NewAttachmentClient(cl.Conn)

	// 2. Set Header
	// Get token By Request
	cl.Context = &ctx
	token, err := em.Micro.Auth.GetTokenFromCtx(ctx)
	if err != nil {
		return err
	}
	cl.Header = map[string]string{"token": token}

	// 3.Do
	err = cl.Sync(func() error {

		r, err := c.CreateMany(ctx, &protobuf.AttachmentCreateMany{
			Service:       em.Micro.Config.ServiceDiscovery.Service.Rpc.Name,
			Paths:		   thumbnail.Paths,
			OwnerId:       thumbnail.OwnerId,
			OwnerType:     thumbnail.OwnerType,
		})
		if err != nil {
			em.LogError.Output(em.MessageWithLineNum_OneRecord(err.Error()))
			return err
		}

		if r.GetStatus() == em.SUCCESS_Status {
			return nil
		} else {
			em.LogError.Output(em.MessageWithLineNum("Create failed!"))
			return errors.New("Create failed!")
		}

	}, nil)
	if err != nil {
		em.LogWarn.OutputSimplePath(err)
		return err
	}

	err = cl.Sync(func() error {

		r, err := c.Append(ctx, &protobuf.AttachmentAppend{
			Service:       em.Micro.Config.ServiceDiscovery.Service.Rpc.Name,
			Paths:		   content.Paths,
			OwnerId:       content.OwnerId,
			OwnerType:     content.OwnerType,
		})
		if err != nil {
			em.LogError.Output(em.MessageWithLineNum_OneRecord(err.Error()))
			return err
		}

		if r.GetStatus() == em.SUCCESS_Status {
			return nil
		} else {
			em.LogError.Output(em.MessageWithLineNum("Create failed!"))
			return errors.New("Create failed!")
		}

	}, nil)
	if err != nil {
		em.LogWarn.OutputSimplePath(err)
		return err
	}

	return nil
}

