package query

import (
	"context"
	proto2 "github.com/nsdash/auth-lib/proto"
	"github.com/nsdash/auth-lib/useCase/auth/shared"
	"google.golang.org/grpc"
)

type CheckTokenQueryHandler struct {
	clientConnection grpc.ClientConnInterface
}

func NewCheckTokenQueryHandler(clientConnection grpc.ClientConnInterface) CheckTokenQueryHandler {
	return CheckTokenQueryHandler{
		clientConnection: clientConnection,
	}
}

func (h CheckTokenQueryHandler) Handle(token string) (*shared.TokenInfoVm, error) {
	request := proto2.CheckTokenRequest{
		AccessToken: token,
	}

	client := proto2.NewGrpcMessageHandlerClient(h.clientConnection)

	response, err := client.CheckToken(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	vm := h.transformResponseToVm(*response)

	return &vm, nil
}

func (CheckTokenQueryHandler) transformResponseToVm(response proto2.CheckTokenResponse) shared.TokenInfoVm {
	return shared.TokenInfoVm{
		UserId: uint(response.UserId),
	}
}
