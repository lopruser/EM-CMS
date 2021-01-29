package main

import (
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/database"
	"github.com/Etpmls/EM-CMS/src/register"
	em "github.com/Etpmls/Etpmls-Micro/v2"
)

func main()  {
	var reg = em.Register{
		AppVersion: 		map[string]string{"EM-CMS Version": application.Version_Service},
		AppEnabledFeatureName:		[]string{em.EnableCircuitBreaker, em.EnableDatabase, em.EnableI18n, em.EnableServiceDiscovery, em.EnableValidator},
		RpcServiceFunc:           register.RegisterRpcService,
		HttpServiceFunc:           register.RegisterHttpService,
		HttpRouteFunc:                 register.RegisterRoute,
		CustomServerFunc: 			[]func(){register.RunWebServer},
		CustomServerServiceRegisterFunc: register.RegisterWebServer,
		CustomServerServiceExitFunc:	register.ExitWebServer,
		DatabaseMigrate:		[]interface{}{
			&database.Category{},
			&database.Post{},
			&database.Variable{},
		},
	}
	reg.Init()
	reg.Run()
}