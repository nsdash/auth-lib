package query

import (
	"context"
	proto2 "github.com/nsdash/auth-lib/proto"
	"github.com/nsdash/auth-lib/useCase/auth/shared"
	errorHandler "github.com/nsdash/go-error-handler"
	customError "github.com/nsdash/go-error-lib"
	"google.golang.org/grpc"
)

type CheckTokenQueryHandler struct {
	clientConnection grpc.ClientConnInterface
	errorHandler     errorHandler.GrpcErrorHandler
}

func NewCheckTokenQueryHandler(clientConnection grpc.ClientConnInterface) CheckTokenQueryHandler {
	return CheckTokenQueryHandler{
		clientConnection: clientConnection,
	}
}

func (h CheckTokenQueryHandler) Handle(token string) (*shared.TokenInfoVm, customError.Error) {
	request := proto2.CheckTokenRequest{
		AccessToken: token,
	}

	client := proto2.NewGrpcMessageHandlerClient(h.clientConnection)

	response, err := client.CheckToken(context.Background(), &request)

	if err != nil {
		return nil, h.errorHandler.Handle(err)
	}

	vm := h.transformResponseToVm(*response)

	return &vm, nil
}

func (CheckTokenQueryHandler) transformResponseToVm(response proto2.CheckTokenResponse) shared.TokenInfoVm {
	return shared.TokenInfoVm{
		UserId: uint(response.UserId),
	}
}
