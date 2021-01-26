package register

import (
	"context"
	"github.com/Etpmls/EM-CMS/src/application/protobuf"
	"github.com/Etpmls/EM-CMS/src/application/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Register Rpc Service
func RegisterRpcService(s *grpc.Server)  {
	// protobuf.RegisterUserServer(s, &service.ServiceUser{})
	protobuf.RegisterCategoryServer(s, &service.ServiceCategory{})
	protobuf.RegisterPostServer(s, &service.ServicePost{})
	protobuf.RegisterVariableServer(s, &service.ServiceVariable{})
	protobuf.RegisterPageServer(s, &service.ServicePage{})
	protobuf.RegisterSettingServer(s, &service.ServiceSetting{})
	return
}

// Register Http Service
func RegisterHttpService(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint *string, opts []grpc.DialOption) error {
	/*err := protobuf.RegisterUserHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}*/
	err := protobuf.RegisterCategoryHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = protobuf.RegisterPostHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = protobuf.RegisterVariableHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = protobuf.RegisterPageHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	err = protobuf.RegisterSettingHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	return nil
}
