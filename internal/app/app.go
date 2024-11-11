package app

import (
	grpcapp "github.com/farid21ola/sso/internal/app/grpc"
	"github.com/farid21ola/sso/internal/services/auth"
	"github.com/farid21ola/sso/internal/storage/postgres"
	"golang.org/x/exp/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration) *App {
	storage, err := postgres.New(connectionString)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
