package main

import (
	"github.com/Etpmls/EM-CMS/src/application"
	"github.com/Etpmls/EM-CMS/src/application/database"
	"github.com/Etpmls/EM-CMS/src/register"
	"github.com/Etpmls/Etpmls-Micro"
)

func main()  {
	var reg = em.Register{
		Version: 		map[string]string{"EM-CMS Version": application.Version_Service},
		EnabledFeatureName:		[]string{em.EnableCircuitBreaker, em.EnableDatabase, em.EnableI18n, em.EnableServiceDiscovery, em.EnableValidator},
		GrpcServiceFunc:           register.RegisterRpcService,
		HttpServiceFunc:           register.RegisterHttpService,
		RouteFunc:                 register.RegisterRoute,
		DatabaseMigrate:		[]interface{}{
			&database.Category{},
			&database.Post{},
			&database.Variable{},
		},
	}
	reg.Init()
	reg.Run()
}