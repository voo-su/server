package grpc

import (
	"context"
	cliV2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc/middleware"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
	"voo.su/internal/provider"
)

type AppProvider struct {
	Conf             *config.Config
	Middleware       middleware.Middleware
	AuthHandler      pb.AuthServiceServer
	AccountHandler   pb.AccountServiceServer
	ChatHandler      pb.ChatServiceServer
	GroupChatHandler pb.GroupChatServiceServer
	ContactHandler   pb.ContactServiceServer
	UploadHandler    pb.UploadServiceServer
	LoggerRepository *clickhouseRepo.LoggerRepository
}

func Run(ctx2 *cliV2.Context, app *AppProvider) error {
	log.SetOutput(provider.NewLoggerWriter(app.Conf, os.Stdout, app.LoggerRepository))

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {

		listener, err := net.Listen(
			app.Conf.Server.Grpc.Protocol,
			app.Conf.Server.Grpc.GetGrpc(),
		)
		if err != nil {
			return err
		}

		srv := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				app.Middleware.Logging.UnaryLoggingInterceptor,
				app.Middleware.Auth.UnaryAuthInterceptor,
			),
			grpc.ChainStreamInterceptor(
				app.Middleware.Logging.StreamLoggingInterceptor,
				app.Middleware.Auth.StreamAuthInterceptor,
			),
		)

		pb.RegisterAuthServiceServer(srv, app.AuthHandler)
		pb.RegisterAccountServiceServer(srv, app.AccountHandler)
		pb.RegisterChatServiceServer(srv, app.ChatHandler)
		pb.RegisterGroupChatServiceServer(srv, app.GroupChatHandler)
		pb.RegisterContactServiceServer(srv, app.ContactHandler)
		pb.RegisterUploadServiceServer(srv, app.UploadHandler)

		reflection.Register(srv)

		log.Printf(
			"GRPC %s://%s:%d \n",
			app.Conf.Server.Grpc.Protocol,
			app.Conf.Server.Grpc.Host,
			app.Conf.Server.Grpc.Port,
		)

		if err := srv.Serve(listener); err != nil {
			return err
		}

		return nil
	})

	log.Fatal(group.Wait())

	return nil
}
