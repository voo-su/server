package grpc

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	cliV2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"voo.su/api/grpc/pb"
	"voo.su/internal/config"
	"voo.su/internal/transport/grpc/middleware"
)

type AppProvider struct {
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	RoutesServices  *middleware.GrpcMethodService
	AuthHandler     pb.AuthServiceServer
	ChatHandler     pb.ChatServiceServer
	ContactHandler  pb.ContactServiceServer
}

func serve(app *AppProvider) error {
	listener, err := net.Listen(
		app.Conf.Server.Grpc.Protocol,
		app.Conf.Server.Grpc.GetGrpc(),
	)
	if err != nil {
		return err
	}

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingServerInterceptor,
		middleware.AuthorizationServerInterceptor,
	)))

	pb.RegisterAuthServiceServer(srv, app.AuthHandler)
	pb.RegisterChatServiceServer(srv, app.ChatHandler)
	pb.RegisterContactServiceServer(srv, app.ContactHandler)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}

func Run(ctx2 *cliV2.Context, app *AppProvider) error {
	ctx := context.Background()

	ctx = middleware.RegisterGlobalService(ctx, app.TokenMiddleware)
	ctx = middleware.RegisterGlobalService(ctx, app.RoutesServices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {
		log.Printf(
			"GRPC %s://%s:%d \n",
			app.Conf.Server.Grpc.Protocol,
			app.Conf.Server.Grpc.Host,
			app.Conf.Server.Grpc.Port,
		)
		return serve(app)
	})

	log.Fatal(group.Wait())

	return nil
}
