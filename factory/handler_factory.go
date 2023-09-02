package factory

import (
	"fmt"
	"github.com/nsdash/auth-lib/useCase/auth/query"
	"github.com/nsdash/config-lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HandlerFactory struct {
	config config.Manager
}

func NewHandlerFactory() HandlerFactory {
	return HandlerFactory{
		config: config.NewManager(),
	}
}

func (f HandlerFactory) createCheckTokenQueryHandler() (*query.CheckTokenQueryHandler, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", f.config.Get("AUTH_SERVICE_HOST"), f.config.Get("AUTH_SERVICE_PORT")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	handler := query.NewCheckTokenQueryHandler(conn)
	return &handler, nil
}
