package webModel

import (
	"context"
	em "github.com/Etpmls/Etpmls-Micro/v2"
)

type common struct {

}

func (this *common) GetPublicToken() (context.Context, error) {
	ctx := context.Background()
	s, err := em.Micro.Auth.CreateGeneralToken(0, "public")
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "token", s)
	return ctx, nil
}
