package grpc

import (
	"context"
	cliV2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc/middleware"
)

type AppProvider struct {
	Conf             *config.Config
	Middleware       middleware.Middleware
	AuthHandler      pb.AuthServiceServer
	AccountHandler   pb.AccountServiceServer
	ChatHandler      pb.ChatServiceServer
	GroupChatHandler pb.GroupChatServiceServer
	ContactHandler   pb.ContactServiceServer
}

func Run(ctx2 *cliV2.Context, app *AppProvider) error {
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

		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

		srv := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				middleware.UnaryLoggingInterceptor,
				app.Middleware.Auth.UnaryAuthInterceptor,
			),
			grpc.ChainStreamInterceptor(
				middleware.StreamLoggingInterceptor,
				app.Middleware.Auth.StreamAuthInterceptor,
			),
		)

		pb.RegisterAuthServiceServer(srv, app.AuthHandler)
		pb.RegisterAccountServiceServer(srv, app.AccountHandler)
		pb.RegisterChatServiceServer(srv, app.ChatHandler)
		pb.RegisterGroupChatServiceServer(srv, app.GroupChatHandler)
		pb.RegisterContactServiceServer(srv, app.ContactHandler)

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
