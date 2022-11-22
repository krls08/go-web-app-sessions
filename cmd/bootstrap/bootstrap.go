package bootstrap

import (
	"context"

	"github.com/krls08/go-web-app-sessions/internal/server"
)

func Run() error {

	ctx := context.TODO()
	s := server.New(ctx, "localhost", 60002)

	return s.Run(ctx)
}
