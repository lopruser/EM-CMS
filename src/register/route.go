package register

import (
	em "github.com/Etpmls/Etpmls-Micro"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

// Register Route
func RegisterRoute(mux *runtime.ServeMux)  {
	mux.HandlePath("GET", em.Micro.Config.ServiceDiscovery.Service.CheckUrl, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("hello"))
	})
}